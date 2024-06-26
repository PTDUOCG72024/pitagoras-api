package employees

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Employee struct {
	ID                   primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name                 string             `bson:"name,omitempty" json:"name"`
	IdentificationType   string             `bson:"identification_type,omitempty" json:"identification_type"`
	IdentificationNumber string             `bson:"identification_number,omitempty" json:"identification_number"`
	Email                string             `bson:"email,omitempty" json:"email"`
	InLaw                bool               `bson:"in_law,omitempty" json:"in_law"`
	ContractType         string             `bson:"contract_type,omitempty" json:"contract_type"`
	BirthDate            time.Time          `bson:"birth_date,omitempty" json:"birth_date"`
	StartDate            time.Time          `bson:"start_date,omitempty" json:"start_date"`
	EndDate              time.Time          `bson:"end_date,omitempty" json:"end_date"`
	Supervisor           Supervisor         `bson:"supervisor,omitempty" json:"supervisor"`
	Position             Position           `bson:"position,omitempty" json:"position"`
	Nationality          Nationality        `bson:"nationality,omitempty" json:"nationality"`
	IsActive             bool               `bson:"is_active,omitempty" json:"is_active"`
	IsDeleted            bool               `bson:"is_deleted,omitempty" json:"is_deleted"`
	CreatedAt            time.Time          `bson:"created_at,omitempty" json:"created_at" default:"now()"`
	UpdatedAt            time.Time          `bson:"updated_at,omitempty" json:"updated_at" default:"now()"`
	DeleteAt             time.Time          `bson:"delete_at,omitempty" json:"delete_at"`
}

type Supervisor struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name,omitempty" json:"name"`
	IsActive  bool               `bson:"is_active,omitempty" json:"is_active"`
	IsDeleted bool               `bson:"is_deleted,omitempty" json:"is_deleted"`
	CreatedAt time.Time          `bson:"created_at,omitempty" json:"created_at" default:"now()"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty" json:"updated_at" default:"now()"`
	DeleteAt  time.Time          `bson:"delete_at,omitempty" json:"delete_at"`
}

type Position struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name,omitempty" json:"name"`
	IsActive  bool               `bson:"is_active,omitempty" json:"is_active"`
	IsDeleted bool               `bson:"is_deleted,omitempty" json:"is_deleted"`
	CreatedAt time.Time          `bson:"created_at,omitempty" json:"created_at" default:"now()"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty" json:"updated_at" default:"now()"`
	DeleteAt  time.Time          `bson:"delete_at,omitempty" json:"delete_at"`
}

type Nationality struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name,omitempty" json:"name"`
	IsActive  bool               `bson:"is_active,omitempty" json:"is_active"`
	IsDeleted bool               `bson:"is_deleted,omitempty" json:"is_deleted"`
	CreatedAt time.Time          `bson:"created_at,omitempty" json:"created_at" default:"now()"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty" json:"updated_at" default:"now()"`
	DeleteAt  time.Time          `bson:"delete_at,omitempty" json:"delete_at"`
}
