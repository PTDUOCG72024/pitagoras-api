package accidents

import (
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
	AccidentDate     primitive.DateTime `bson:"accident_date,omitempty" json:"accident_date"`
}

type Gravity struct {
	ID   string `bson:"_id,omitempty" json:"id"`
	Name string `bson:"name,omitempty" json:"name"`
}

type InjuredPart struct {
	ID   string `bson:"_id,omitempty" json:"id"`
	Name string `bson:"name,omitempty" json:"name"`
}

type Classification struct {
	ID   string `bson:"_id,omitempty" json:"id"`
	Name string `bson:"name,omitempty" json:"name"`
}
