package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	ID          primitive.ObjectID `json:"_id,omitempty"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	UserID      primitive.ObjectID `json:"user_id"`
}