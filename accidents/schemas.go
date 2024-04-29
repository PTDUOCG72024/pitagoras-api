package accidents

import (
	"time"

	"github.com/xbizzybone/pitagoras-api/employees"
	"github.com/xbizzybone/pitagoras-api/projects"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateAccidentRequest struct {
	Description      string             `json:"description"`
	ConstructionArea string             `json:"construction_area"`
	ClassificationID primitive.ObjectID `json:"classification_id"`
	GravityID        primitive.ObjectID `json:"gravity_id"`
	ProjectID        primitive.ObjectID `json:"project_id"`
	EmployeeID       primitive.ObjectID `json:"employee_id"`
	InjuredPartID    primitive.ObjectID `json:"injured_part_id"`
	AccidentDate     time.Time          `json:"accident_date"`
}

type CreateAccidentResponse struct {
	ID               primitive.ObjectID `json:"id"`
	Description      string             `json:"description"`
	ConstructionArea string             `json:"construction_area"`
	Classification   Classification     `json:"classification"`
	Gravity          Gravity            `json:"gravity"`
	Project          projects.Project   `json:"project"`
	Employee         employees.Employee `json:"employee"`
	InjuredPart      InjuredPart        `json:"injured_part"`
	AccidentDate     time.Time          `json:"accident_date"`
	IsDeleted        bool               `json:"is_deleted"`
	CreatedAt        time.Time          `json:"created_at"`
}

type UpdateAccidentRequest struct {
	Description      string             `json:"description"`
	ConstructionArea string             `json:"construction_area"`
	ClassificationID primitive.ObjectID `json:"classification_id"`
	GravityID        primitive.ObjectID `json:"gravity_id"`
	ProjectID        primitive.ObjectID `json:"project_id"`
	EmployeeID       primitive.ObjectID `json:"employee_id"`
	InjuredPartID    primitive.ObjectID `json:"injured_part_id"`
	AccidentDate     time.Time          `json:"accident_date"`
}

type UpdateAccidentResponse struct {
	ID               primitive.ObjectID `json:"id"`
	Description      string             `json:"description"`
	ConstructionArea string             `json:"construction_area"`
	Classification   Classification     `json:"classification"`
	Gravity          Gravity            `json:"gravity"`
	Project          projects.Project   `json:"project"`
	Employee         employees.Employee `json:"employee"`
	InjuredPart      InjuredPart        `json:"injured_part"`
	AccidentDate     time.Time          `json:"accident_date"`
	IsDeleted        bool               `json:"is_deleted"`
	CreatedAt        time.Time          `json:"created_at"`
	UpdateAt         time.Time          `json:"updated_at"`
}

type CreateClassificationRequest struct {
	Name string `json:"name"`
}

type CreateClassificationResponse struct {
	ID        primitive.ObjectID `json:"id"`
	Name      string             `json:"name"`
	IsActive  bool               `json:"is_active"`
	IsDeleted bool               `json:"is_deleted"`
	CreatedAt time.Time          `json:"created_at"`
}

type UpdateClassificationRequest struct {
	Name string `json:"name"`
}

type UpdateClassificationResponse struct {
	ID        primitive.ObjectID `json:"id"`
	Name      string             `json:"name"`
	IsActive  bool               `json:"is_active"`
	IsDeleted bool               `json:"is_deleted"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

type CreateGravityRequest struct {
	Name string `json:"name"`
}

type CreateGravityResponse struct {
	ID        primitive.ObjectID `json:"id"`
	Name      string             `json:"name"`
	IsActive  bool               `json:"is_active"`
	IsDeleted bool               `json:"is_deleted"`
	CreatedAt time.Time          `json:"created_at"`
}

type UpdateGravityRequest struct {
	Name string `json:"name"`
}

type UpdateGravityResponse struct {
	ID        primitive.ObjectID `json:"id"`
	Name      string             `json:"name"`
	IsActive  bool               `json:"is_active"`
	IsDeleted bool               `json:"is_deleted"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

type CreateInjuredPartRequest struct {
	Name string `json:"name"`
}

type CreateInjuredPartResponse struct {
	ID        primitive.ObjectID `json:"id"`
	Name      string             `json:"name"`
	IsActive  bool               `json:"is_active"`
	IsDeleted bool               `json:"is_deleted"`
	CreatedAt time.Time          `json:"created_at"`
}

type UpdateInjuredPartRequest struct {
	Name string `json:"name"`
}

type UpdateInjuredPartResponse struct {
	ID        primitive.ObjectID `json:"id"`
	Name      string             `json:"name"`
	IsActive  bool               `json:"is_active"`
	IsDeleted bool               `json:"is_deleted"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

type GetAccidentByIdResponse struct {
	ID               primitive.ObjectID `json:"id"`
	Description      string             `json:"description"`
	ConstructionArea string             `json:"construction_area"`
	Classification   Classification     `json:"classification"`
	Gravity          Gravity            `json:"gravity"`
	Project          projects.Project   `json:"project"`
	Employee         employees.Employee `json:"employee"`
	InjuredPart      InjuredPart        `json:"injured_part"`
	AccidentDate     time.Time          `json:"accident_date"`
	IsDeleted        bool               `json:"is_deleted"`
	CreatedAt        time.Time          `json:"created_at"`
	UpdatedAt        time.Time          `json:"updated_at"`
	DeletedAt        time.Time          `json:"deleted_at"`
}

type GetClassificationByIdResponse struct {
	ID        primitive.ObjectID `json:"id"`
	Name      string             `json:"name"`
	IsActive  bool               `json:"is_active"`
	IsDeleted bool               `json:"is_deleted"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
	DeletedAt time.Time          `json:"deleted_at"`
}

type GetGravityByIdResponse struct {
	ID        primitive.ObjectID `json:"id"`
	Name      string             `json:"name"`
	IsActive  bool               `json:"is_active"`
	IsDeleted bool               `json:"is_deleted"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
	DeletedAt time.Time          `json:"deleted_at"`
}

type GetInjuredPartByIdResponse struct {
	ID        primitive.ObjectID `json:"id"`
	Name      string             `json:"name"`
	IsActive  bool               `json:"is_active"`
	IsDeleted bool               `json:"is_deleted"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
	DeletedAt time.Time          `json:"deleted_at"`
}
