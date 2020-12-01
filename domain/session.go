package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Session struct {
	ID          primitive.ObjectID
	SessionUUID string `bson:"session_uuid"`
	UserID      string `bson:"user_id"`
	Password    string `bson:"password"`
	Status      string `bson:"status"`
}
