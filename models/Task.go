package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"
	Message     string             `json:"message"`
	Description string             `json:"descriptin"`
	IsDone      bool               `json:"is_done" bson"is_done"`
	CreatedAt   primitive.DateTime `json:"created_at" bson:"create_at"`
	User        User               `json:"user"`
}
