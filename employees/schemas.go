package employees

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateEmployeeRequest struct {
	Name                 string             `json:"name"`
	IdentificationType   string             `json:"identification_type"`
	IdentificationNumber string             `json:"identification_number"`
	Email                string             `json:"email"`
	InLaw                bool               `json:"in_law"`
	ContractType         string             `json:"contract_type"`
	StartDate            time.Time          `json:"start_date"`
	EndDate              time.Time          `json:"end_date"`
	SupervisorID         primitive.ObjectID `json:"supervisor_id"`
	PositionID           primitive.ObjectID `json:"position_id"`
	NationalityID        primitive.ObjectID `json:"nationality_id"`
}

type CreateEmployeeResponse struct {
	ID                   primitive.ObjectID `json:"id"`
	Name                 string             `json:"name"`
	IdentificationType   string             `json:"identification_type"`
	IdentificationNumber string             `json:"identification_number"`
	Email                string             `json:"email"`
	InLaw                bool               `json:"in_law"`
	ContractType         string             `json:"contract_type"`
	StartDate            time.Time          `json:"start_date"`
	EndDate              time.Time          `json:"end_date"`
	IsActive             bool               `json:"is_active"`
	IsDeleted            bool               `json:"is_deleted"`
	Supervisor           Supervisor         `json:"supervisor"`
	Position             Position           `json:"position"`
	Nationality          Nationality        `json:"nationality"`
	CreatedAt            time.Time          `json:"created_at"`
}

type UpdateEmployeeRequest struct {
	Name                 string             `json:"name"`
	IdentificationType   string             `json:"identification_type"`
	IdentificationNumber string             `json:"identification_number"`
	Email                string             `json:"email"`
	InLaw                bool               `json:"in_law"`
	ContractType         string             `json:"contract_type"`
	StartDate            time.Time          `json:"start_date"`
	EndDate              time.Time          `json:"end_date"`
	SupervisorID         primitive.ObjectID `json:"supervisor_id"`
	PositionID           primitive.ObjectID `json:"position_id"`
	NationalityID        primitive.ObjectID `json:"nationality_id"`
}

type UpdateEmployeeResponse struct {
	ID                   primitive.ObjectID `json:"id"`
	Name                 string             `json:"name"`
	IdentificationType   string             `json:"identification_type"`
	IdentificationNumber string             `json:"identification_number"`
	Email                string             `json:"email"`
	InLaw                bool               `json:"in_law"`
	ContractType         string             `json:"contract_type"`
	StartDate            time.Time          `json:"start_date"`
	EndDate              time.Time          `json:"end_date"`
	IsActive             bool               `json:"is_active"`
	IsDeleted            bool               `json:"is_deleted"`
	Supervisor           Supervisor         `json:"supervisor"`
	Position             Position           `json:"position"`
	Nationality          Nationality        `json:"nationality"`
	CreatedAt            time.Time          `json:"created_at"`
	UpdatedAt            time.Time          `json:"updated_at"`
}

type CreateSupervisorRequest struct {
	Name string `json:"name"`
}

type CreateSupervisorResponse struct {
	ID        primitive.ObjectID `json:"id"`
	Name      string             `json:"name"`
	IsActive  bool               `json:"is_active"`
	IsDeleted bool               `json:"is_deleted"`
	CreatedAt time.Time          `json:"created_at"`
}

type UpdateSupervisorRequest struct {
	Name string `json:"name"`
}

type UpdateSupervisorResponse struct {
	ID        primitive.ObjectID `json:"id"`
	Name      string             `json:"name"`
	IsActive  bool               `json:"is_active"`
	IsDeleted bool               `json:"is_deleted"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

type CreatePositionRequest struct {
	Name string `json:"name"`
}

type CreatePositionResponse struct {
	ID        primitive.ObjectID `json:"id"`
	Name      string             `json:"name"`
	IsActive  bool               `json:"is_active"`
	IsDeleted bool               `json:"is_deleted"`
	CreatedAt time.Time          `json:"created_at"`
}

type UpdatePositionRequest struct {
	Name string `json:"name"`
}

type UpdatePositionResponse struct {
	ID        primitive.ObjectID `json:"id"`
	Name      string             `json:"name"`
	IsActive  bool               `json:"is_active"`
	IsDeleted bool               `json:"is_deleted"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

type CreateNationalityRequest struct {
	Name string `json:"name"`
}

type CreateNationalityResponse struct {
	ID        primitive.ObjectID `json:"id"`
	Name      string             `json:"name"`
	IsActive  bool               `json:"is_active"`
	IsDeleted bool               `json:"is_deleted"`
	CreatedAt time.Time          `json:"created_at"`
}

type UpdateNationalityRequest struct {
	Name string `json:"name"`
}

type UpdateNationalityResponse struct {
	ID        primitive.ObjectID `json:"id"`
	Name      string             `json:"name"`
	IsActive  bool               `json:"is_active"`
	IsDeleted bool               `json:"is_deleted"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}

type GetEmployeeResponse struct {
	ID                   primitive.ObjectID `json:"id"`
	Name                 string             `json:"name"`
	IdentificationType   string             `json:"identification_type"`
	IdentificationNumber string             `json:"identification_number"`
	Email                string             `json:"email"`
	InLaw                bool               `json:"in_law"`
	ContractType         string             `json:"contract_type"`
	StartDate            time.Time          `json:"start_date"`
	EndDate              time.Time          `json:"end_date"`
	IsActive             bool               `json:"is_active"`
	IsDeleted            bool               `json:"is_deleted"`
	Supervisor           Supervisor         `json:"supervisor"`
	Position             Position           `json:"position"`
	Nationality          Nationality        `json:"nationality"`
	CreatedAt            string             `json:"created_at"`
	UpdatedAt            time.Time          `json:"updated_at"`
	DeleteAt             time.Time          `json:"delete_at"`
}

type GetSupervisorResponse struct {
	ID        primitive.ObjectID `json:"id"`
	Name      string             `json:"name"`
	IsActive  bool               `json:"is_active"`
	IsDeleted bool               `json:"is_deleted"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
	DeleteAt  time.Time          `json:"delete_at"`
}

type GetPositionResponse struct {
	ID        primitive.ObjectID `json:"id"`
	Name      string             `json:"name"`
	IsActive  bool               `json:"is_active"`
	IsDeleted bool               `json:"is_deleted"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
	DeleteAt  time.Time          `json:"delete_at"`
}

type GetNationalityResponse struct {
	ID        primitive.ObjectID `json:"id"`
	Name      string             `json:"name"`
	IsActive  bool               `json:"is_active"`
	IsDeleted bool               `json:"is_deleted"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
	DeleteAt  time.Time          `json:"delete_at"`
}
