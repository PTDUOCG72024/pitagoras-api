package accidents

import (
	"time"

	"github.com/xbizzybone/pitagoras-api/employees"
	"github.com/xbizzybone/pitagoras-api/projects"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Accident struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Description      string             `bson:"description,omitempty" json:"description"`
	ConstructionArea string             `bson:"construction_area,omitempty" json:"construction_area"`
	Classification   Classification     `bson:"classification,omitempty" json:"classification"`
	Gravity          Gravity            `bson:"gravity,omitempty" json:"gravity"`
	Project          projects.Project   `bson:"project,omitempty" json:"project"`
	Employee         employees.Employee `bson:"employee,omitempty" json:"employee"`
	InjuredPart      InjuredPart        `bson:"injured_part,omitempty" json:"injured_part"`
	AccidentDate     time.Time          `bson:"accident_date,omitempty" json:"accident_date"`
	IsDeleted        bool               `bson:"is_deleted,omitempty" json:"is_deleted"`
	CreatedAt        time.Time          `bson:"created_at,omitempty" json:"created_at"`
	UpdatedAt        time.Time          `bson:"updated_at,omitempty" json:"updated_at"`
	DeletedAt        time.Time          `bson:"deleted_at,omitempty" json:"deleted_at"`
}

type Gravity struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name,omitempty" json:"name"`
	IsActive  bool               `bson:"is_active,omitempty" json:"is_active"`
	IsDeleted bool               `bson:"is_deleted,omitempty" json:"is_deleted"`
	CreatedAt time.Time          `bson:"created_at,omitempty" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty" json:"updated_at"`
	DeletedAt time.Time          `bson:"deleted_at,omitempty" json:"deleted_at"`
}

type InjuredPart struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name,omitempty" json:"name"`
	IsActive  bool               `bson:"is_active,omitempty" json:"is_active"`
	IsDeleted bool               `bson:"is_deleted,omitempty" json:"is_deleted"`
	CreatedAt time.Time          `bson:"created_at,omitempty" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty" json:"updated_at"`
	DeletedAt time.Time          `bson:"deleted_at,omitempty" json:"deleted_at"`
}

type Classification struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name,omitempty" json:"name"`
	IsActive  bool               `bson:"is_active,omitempty" json:"is_active"`
	IsDeleted bool               `bson:"is_deleted,omitempty" json:"is_deleted"`
	CreatedAt time.Time          `bson:"created_at,omitempty" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty" json:"updated_at"`
	DeletedAt time.Time          `bson:"deleted_at,omitempty" json:"deleted_at"`
}
