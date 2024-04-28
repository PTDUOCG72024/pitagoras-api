package users

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
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
