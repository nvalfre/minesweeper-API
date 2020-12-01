package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID
	UserUUID string `bson:"user_uuid"`
	Nickname string `bson:"nickname"`
}
