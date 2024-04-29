package employees

type CreateEmployeeRequest struct {
	Name                 string `json:"name"`
	IdentificationType   string `json:"identification_type"`
	IdentificationNumber string `json:"identification_number"`
	Email                string `json:"email"`
	InLaw                bool   `json:"in_law"`
	ContractType         string `json:"contract_type"`
	StartDate            string `json:"start_date"`
	EndDate              string `json:"end_date"`
	SupervisorID         string `json:"supervisor_id"`
	PositionID           string `json:"position_id"`
	NationalityID        string `json:"nationality_id"`
}

type CreateEmployeeResponse struct {
	ID                   string      `json:"id"`
	Name                 string      `json:"name"`
	IdentificationType   string      `json:"identification_type"`
	IdentificationNumber string      `json:"identification_number"`
	Email                string      `json:"email"`
	InLaw                bool        `json:"in_law"`
	ContractType         string      `json:"contract_type"`
	StartDate            string      `json:"start_date"`
	EndDate              string      `json:"end_date"`
	IsActive             bool        `json:"is_active"`
	IsDeleted            bool        `json:"is_deleted"`
	Supervisor           Supervisor  `json:"supervisor"`
	Position             Position    `json:"position"`
	Nationality          Nationality `json:"nationality"`
	CreatedAt            string      `json:"created_at"`
}

type UpdateEmployeeRequest struct {
	Name                 string `json:"name"`
	IdentificationType   string `json:"identification_type"`
	IdentificationNumber string `json:"identification_number"`
	Email                string `json:"email"`
	InLaw                bool   `json:"in_law"`
	ContractType         string `json:"contract_type"`
	StartDate            string `json:"start_date"`
	EndDate              string `json:"end_date"`
	SupervisorID         string `json:"supervisor_id"`
	PositionID           string `json:"position_id"`
	NationalityID        string `json:"nationality_id"`
}

type UpdateEmployeeResponse struct {
	ID                   string      `json:"id"`
	Name                 string      `json:"name"`
	IdentificationType   string      `json:"identification_type"`
	IdentificationNumber string      `json:"identification_number"`
	Email                string      `json:"email"`
	InLaw                bool        `json:"in_law"`
	ContractType         string      `json:"contract_type"`
	StartDate            string      `json:"start_date"`
	EndDate              string      `json:"end_date"`
	IsActive             bool        `json:"is_active"`
	IsDeleted            bool        `json:"is_deleted"`
	Supervisor           Supervisor  `json:"supervisor"`
	Position             Position    `json:"position"`
	Nationality          Nationality `json:"nationality"`
	CreatedAt            string      `json:"created_at"`
	UpdatedAt            string      `json:"updated_at"`
}

type CreateSupervisorRequest struct {
	Name string `json:"name"`
}

type CreateSupervisorResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	IsActive  bool   `json:"is_active"`
	IsDeleted bool   `json:"is_deleted"`
	CreatedAt string `json:"created_at"`
}

type UpdateSupervisorRequest struct {
	Name string `json:"name"`
}

type UpdateSupervisorResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	IsActive  bool   `json:"is_active"`
	IsDeleted bool   `json:"is_deleted"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CreatePositionRequest struct {
	Name string `json:"name"`
}

type CreatePositionResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	IsActive  bool   `json:"is_active"`
	IsDeleted bool   `json:"is_deleted"`
	CreatedAt string `json:"created_at"`
}

type UpdatePositionRequest struct {
	Name string `json:"name"`
}

type UpdatePositionResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	IsActive  bool   `json:"is_active"`
	IsDeleted bool   `json:"is_deleted"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CreateNationalityRequest struct {
	Name string `json:"name"`
}

type CreateNationalityResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	IsActive  bool   `json:"is_active"`
	IsDeleted bool   `json:"is_deleted"`
	CreatedAt string `json:"created_at"`
}

type UpdateNationalityRequest struct {
	Name string `json:"name"`
}

type UpdateNationalityResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	IsActive  bool   `json:"is_active"`
	IsDeleted bool   `json:"is_deleted"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type GetEmployeeResponse struct {
	ID                   string      `json:"id"`
	Name                 string      `json:"name"`
	IdentificationType   string      `json:"identification_type"`
	IdentificationNumber string      `json:"identification_number"`
	Email                string      `json:"email"`
	InLaw                bool        `json:"in_law"`
	ContractType         string      `json:"contract_type"`
	StartDate            string      `json:"start_date"`
	EndDate              string      `json:"end_date"`
	IsActive             bool        `json:"is_active"`
	IsDeleted            bool        `json:"is_deleted"`
	Supervisor           Supervisor  `json:"supervisor"`
	Position             Position    `json:"position"`
	Nationality          Nationality `json:"nationality"`
	CreatedAt            string      `json:"created_at"`
	UpdatedAt            string      `json:"updated_at"`
	DeleteAt             string      `json:"delete_at"`
}

type GetSupervisorResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	IsActive  bool   `json:"is_active"`
	IsDeleted bool   `json:"is_deleted"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeleteAt  string `json:"delete_at"`
}

type GetPositionResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	IsActive  bool   `json:"is_active"`
	IsDeleted bool   `json:"is_deleted"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeleteAt  string `json:"delete_at"`
}

type GetNationalityResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	IsActive  bool   `json:"is_active"`
	IsDeleted bool   `json:"is_deleted"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeleteAt  string `json:"delete_at"`
}
