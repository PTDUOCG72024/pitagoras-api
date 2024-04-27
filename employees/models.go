package employees

import "go.mongodb.org/mongo-driver/bson/primitive"

type Employee struct {
	ID                   primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name                 string             `bson:"name,omitempty" json:"name"`
	IdentificationType   string             `bson:"identification_type,omitempty" json:"identification_type"`
	IdentificationNumber string             `bson:"identification_number,omitempty" json:"identification_number"`
	InLaw                bool               `bson:"in_law,omitempty" json:"in_law"`
	ContractType         string             `bson:"contract_type,omitempty" json:"contract_type"`
	StartDate            primitive.DateTime `bson:"start_date,omitempty" json:"start_date"`
	EndDate              primitive.DateTime `bson:"end_date,omitempty" json:"end_date"`
	Supervisor           Supervisor         `bson:"supervisor,omitempty" json:"supervisor"`
	Position             Position           `bson:"position,omitempty" json:"position"`
	Nationality          Nationality        `bson:"nationality,omitempty" json:"nationality"`
}

type Supervisor struct {
	ID       string `bson:"_id,omitempty" json:"id"`
	Name     string `bson:"name,omitempty" json:"name"`
	IsActive bool   `bson:"is_active,omitempty" json:"is_active"`
}

type Position struct {
	ID   string `bson:"_id,omitempty" json:"id"`
	Name string `bson:"name,omitempty" json:"name"`
}

type Nationality struct {
	ID   string `bson:"_id,omitempty" json:"id"`
	Name string `bson:"name,omitempty" json:"name"`
}
