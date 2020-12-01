package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Games struct {
	ID        primitive.ObjectID
	UserUUID  string      `bson:"game_uuid"`
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
