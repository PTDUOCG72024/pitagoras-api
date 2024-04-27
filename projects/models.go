package projects

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Project struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name,omitempty" json:"name"`
	Description string             `bson:"description,omitempty" json:"description"`
	StartDate   time.Time          `bson:"start_date,omitempty" json:"start_date"`
	EndDate     time.Time          `bson:"end_date,omitempty" json:"end_date"`
	IsActive    bool               `bson:"is_active,omitempty" json:"is_active"`
	IsDeleted   bool               `bson:"is_deleted,omitempty" json:"is_deleted"`
	CreatedAt   time.Time          `bson:"created_at,omitempty" json:"created_at" default:"now()"`
	UpdatedAt   time.Time          `bson:"updated_at,omitempty" json:"updated_at" default:"now()"`
	DeleteAt    time.Time          `bson:"delete_at,omitempty" json:"delete_at"`
}
