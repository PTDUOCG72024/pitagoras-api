package users

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Request
type UserCreateRequest struct {
	Name     string `bson:"name,omitempty" json:"name" validate:"required"`
	Email    string `bson:"email,omitempty" json:"email" validate:"required,email"`
	Password string `bson:"password,omitempty" json:"password,omitempty" validate:"required"`
}

type UserUpdateRequest struct {
	Name     string `bson:"name,omitempty" json:"name"`
	Email    string `bson:"email,omitempty" json:"email"`
	Password string `bson:"password,omitempty" json:"password,omitempty"`
}

type UserRequest struct {
	Email     string    `bson:"email,omitempty" json:"email" validate:"required,email"`
	Password  string    `bson:"password,omitempty" json:"password,omitempty" validate:"required"`
	LastLogin time.Time `bson:"last_login,omitempty" json:"last_login"`
	CreatedAt time.Time `bson:"created_at,omitempty" json:"created_at" default:"now()"`
	UpdatedAt time.Time `bson:"updated_at,omitempty" json:"updated_at" default:"now()"`
	DeleteAt  time.Time `bson:"delete_at,omitempty" json:"delete_at"`
}

// Response
type UserResponse struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name,omitempty" json:"name"`
	Email     string             `bson:"email,omitempty" json:"email"`
	Password  string             `bson:"password,omitempty" json:"password,omitempty"`
	IsDeleted bool               `bson:"is_deleted,omitempty" json:"is_deleted"`
	LastLogin time.Time          `bson:"last_login,omitempty" json:"last_login"`
	CreatedAt time.Time          `bson:"created_at,omitempty" json:"created_at" default:"now()"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty" json:"updated_at" default:"now()"`
	DeleteAt  time.Time          `bson:"delete_at,omitempty" json:"delete_at"`
}

type UserLoginResponse struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name,omitempty" json:"name"`
	Email     string             `bson:"email,omitempty" json:"email"`
	Password  string             `bson:"password,omitempty" json:"password,omitempty"`
	LastLogin time.Time          `bson:"last_login,omitempty" json:"last_login"`
	CreatedAt time.Time          `bson:"created_at,omitempty" json:"created_at" default:"now()"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty" json:"updated_at" default:"now()"`
}

type UserCreateResponse struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name  string             `bson:"name,omitempty" json:"name"`
	Email string             `bson:"email,omitempty" json:"email"`
}

type ResultResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}
