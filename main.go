package main

import (
	"context"
	"errors"
	"fmt"
	uuid2 "github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection
var ctx = context.TODO()
var dbName = "minesweeper"
var collectionName = "games"

const gameStartedStatus = "Started"
const defaultDurationTime = 10 * time.Minute

func init() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)

	_ = fmt.Sprint(uuid2.NewUUID())
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	database := client.Database(dbName)
	collection = database.Collection(collectionName)
}

type Games struct {
	ID        primitive.ObjectID
	CreatedAt time.Time   `bson:"created_at"`
	UpdatedAt time.Time   `bson:"updated_at"`
	UserID    string      `bson:"user_id"`
	Finished  bool        `bson:"finished"`
	Timer     *time.Timer `bson:"timer"`
	Details   Details     `bson:"game_details"`
}

type Details struct {
	Board      *Board `bson:"board"`
	Status     string `bson:"status"`
	CountMoves int    `bson:"moves"`
}

type Board struct {
	Positions []Positions `bson:"positions"`
	Actions   []*Actions  `bson:"actions"`
}

type Positions struct {
	status string `bson:"positions"`
}

type Actions struct {
	PositionX int `bson:"position_x"`
	PositionY int `bson:"position_y"`
}

func main() {
	app := &cli.App{
		Name:  "Gamer Starter",
		Usage: "A simple CLI program to manage games",
		Action: func(c *cli.Context) error {
			_, err := getPending()
			if err != nil {
				if err == mongo.ErrNoDocuments {
					fmt.Print("Nothing to see here.\nRun `add 'game'` to add a game")
					return nil
				}

				return err
			}

			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "add a game to the list",
				Action: func(c *cli.Context) error {
					str := c.Args().First()
					if str == "" {
						return errors.New("Cannot add an empty game")
					}

					game := &Games{
						ID:        primitive.NewObjectID(),
						CreatedAt: time.Now(),
						UpdatedAt: time.Now(),
						Details: Details{
							Board: &Board{
								Positions: nil,
								Actions:   nil,
							},
							Status: gameStartedStatus,
						},
						Timer: time.NewTimer(defaultDurationTime),
					}

					return createGame(game)
				},
			},
			{
				Name:    "all",
				Aliases: []string{"l"},
				Usage:   "list all games",
				Action: func(c *cli.Context) error {
					_, err := getAll()
					if err != nil {
						if err == mongo.ErrNoDocuments {
							fmt.Print("Nothing to see here.\nRun `add 'games'` to add a game")
							return nil
						}

						return err
					}

					return nil
				},
			},
			{
				Name:    "done",
				Aliases: []string{"d"},
				Usage:   "complete a game on the list",
				Action: func(c *cli.Context) error {
					text := c.Args().First()
					return completeGames(text)
				},
			},
			{
				Name:    "finished",
				Aliases: []string{"f"},
				Usage:   "list completed game",
				Action: func(c *cli.Context) error {
					_, err := getFinished()
					if err != nil {
						if err == mongo.ErrNoDocuments {
							fmt.Print("Nothing to see here.\nRun `done 'game'` to complete a game")
							return nil
						}

						return err
					}

					return nil
				},
			},
			{
				Name:  "rm",
				Usage: "deletes a game on the list",
				Action: func(c *cli.Context) error {
					text := c.Args().First()
					err := deleteGame(text)
					if err != nil {
						return err
					}

					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func createGame(game *Games) error {
	_, err := collection.InsertOne(ctx, game)
	return err
}

func getAll() ([]*Games, error) {
	// passing bson.D{{}} matches all documents in the collection
	filter := bson.D{{}}
	return filterGames(filter)
}

func filterGames(filter interface{}) ([]*Games, error) {
	// A slice of games for storing the decoded documents
	var games []*Games

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return games, err
	}

	// Iterate through the cursor and decode each document one at a time
	for cur.Next(ctx) {
		var t Games
		err := cur.Decode(&t)
		if err != nil {
			return games, err
		}

		games = append(games, &t)
	}

	if err := cur.Err(); err != nil {
		return games, err
	}

	// once exhausted, close the cursor
	cur.Close(ctx)

	if len(games) == 0 {
		return games, mongo.ErrNoDocuments
	}

	return games, nil
}

func completeGames(text string) error {
	filter := bson.D{primitive.E{Key: "text", Value: text}}

	update := bson.D{primitive.E{Key: "$set", Value: bson.D{
		primitive.E{Key: "completed", Value: true},
	}}}

	t := &Games{}
	return collection.FindOneAndUpdate(ctx, filter, update).Decode(t)
}

func getPending() ([]*Games, error) {
	filter := bson.D{
		primitive.E{Key: "completed", Value: false},
	}

	return filterGames(filter)
}

func getFinished() ([]*Games, error) {
	filter := bson.D{
		primitive.E{Key: "completed", Value: true},
	}

	return filterGames(filter)
}

func deleteGame(text string) error {
	filter := bson.D{primitive.E{Key: "text", Value: text}}

	res, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	if res.DeletedCount == 0 {
		return errors.New("No games were deleted")
	}

	return nil
}
