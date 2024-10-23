package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FirstName    string             `json:"first_name" validate:"required,min=2,max=50"`
	LastName     string             `json:"last_name" validate:"required,min=2,max=50"`
	Password     string             `json:"-" validate:"required,min=8"`
	Email        string             `json:"email" validate:"required,email"`
	Phone        string             `json:"phone" validate:"omitempty,e164"`
	UserType     string             `json:"user_type" validate:"required,oneof=admin user"`
	Token        string             `bson:"-" json:"token,omitempty"`
	RefreshToken string             `bson:"-" json:"refresh_token,omitempty"`
	CreatedAt    time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at" json:"updated_at"`
	UserId       string             `bson:"user_id" json:"user_id"`
}
