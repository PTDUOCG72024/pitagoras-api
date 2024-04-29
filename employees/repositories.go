package employees

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type repository struct {
	logger     *zap.Logger
	collection *mongo.Collection
}

func NewRepository(logger *zap.Logger, collection *mongo.Collection) Repository {
	return &repository{
		logger:     logger,
		collection: collection,
	}
}

// CreateEmployee implements Repository.
func (r *repository) CreateEmployee(ctx context.Context, employee *Employee) (*Employee, error) {
	result, err := r.collection.InsertOne(ctx, employee)
	if err != nil {
		r.logger.Sugar().Errorw(err.Error(), "func", "CreateEmployee", "request_id", ctx.Value("request_id"))
		return nil, err
	}

	employee.ID = result.InsertedID.(primitive.ObjectID)

	return employee, nil
}

// CreateNationality implements Repository.
func (r *repository) CreateNationality(ctx context.Context, nationality *Nationality) (*Nationality, error) {
	result, err := r.collection.InsertOne(ctx, nationality)
	if err != nil {
		r.logger.Sugar().Errorw(err.Error(), "func", "CreateNationality", "request_id", ctx.Value("request_id"))
		return nil, err
	}

	nationality.ID = result.InsertedID.(primitive.ObjectID)

	return nationality, nil
}

// CreatePosition implements Repository.
func (r *repository) CreatePosition(ctx context.Context, position *Position) (*Position, error) {
	result, err := r.collection.InsertOne(ctx, position)
	if err != nil {
		r.logger.Sugar().Errorw(err.Error(), "func", "CreatePosition", "request_id", ctx.Value("request_id"))
		return nil, err
	}

	position.ID = result.InsertedID.(primitive.ObjectID)

	return position, nil
}

// CreateSupervisor implements Repository.
func (r *repository) CreateSupervisor(ctx context.Context, supervisor *Supervisor) (*Supervisor, error) {
	result, err := r.collection.InsertOne(ctx, supervisor)
	if err != nil {
		r.logger.Sugar().Errorw(err.Error(), "func", "CreateSupervisor", "request_id", ctx.Value("request_id"))
		return nil, err
	}

	supervisor.ID = result.InsertedID.(primitive.ObjectID)

	return supervisor, nil
}

// GetEmployeeById implements Repository.
func (r *repository) GetEmployeeById(ctx context.Context, id string) (*Employee, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		r.logger.Sugar().Errorw(err.Error(), "func", "GetEmployeeById")
		return nil, err
	}

	employee := new(Employee)
	if err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(employee); err != nil {
		if err == mongo.ErrNoDocuments {
			r.logger.Sugar().Errorw("employee not found", "func", "GetEmployeeById")
			return nil, err
		}

		r.logger.Sugar().Errorw(err.Error(), "func", "GetEmployeeById")
		return nil, err
	}

	return employee, nil
}

// GetEmployees implements Repository.
func (r *repository) GetEmployees(ctx context.Context) ([]Employee, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		r.logger.Sugar().Errorw(err.Error(), "func", "GetEmployees")
		return nil, err
	}

	var employees []Employee
	if err = cursor.All(ctx, &employees); err != nil {
		r.logger.Sugar().Errorw(err.Error(), "func", "GetEmployees")
		return nil, err
	}

	return employees, nil
}

// GetNationalities implements Repository.
func (r *repository) GetNationalities(ctx context.Context) ([]Nationality, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		r.logger.Sugar().Errorw(err.Error(), "func", "GetNationalities")
		return nil, err
	}

	var nationalities []Nationality
	if err = cursor.All(ctx, &nationalities); err != nil {
		r.logger.Sugar().Errorw(err.Error(), "func", "GetNationalities")
		return nil, err
	}

	return nationalities, nil
}

// GetNationalityById implements Repository.
func (r *repository) GetNationalityById(ctx context.Context, id string) (*Nationality, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		r.logger.Sugar().Errorw(err.Error(), "func", "GetNationalityById")
		return nil, err
	}

	nationality := new(Nationality)
	if err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(nationality); err != nil {
		if err == mongo.ErrNoDocuments {
			r.logger.Sugar().Errorw("employee not found", "func", "GetNationalityById")
			return nil, err
		}

		r.logger.Sugar().Errorw(err.Error(), "func", "GetNationalityById")
		return nil, err
	}

	return nationality, nil
}

// GetPositionById implements Repository.
func (r *repository) GetPositionById(ctx context.Context, id string) (*Position, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		r.logger.Sugar().Errorw(err.Error(), "func", "GetPositionById")
		return nil, err
	}

	position := new(Position)
	if err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(position); err != nil {
		if err == mongo.ErrNoDocuments {
			r.logger.Sugar().Errorw("employee not found", "func", "GetPositionById")
			return nil, err
		}

		r.logger.Sugar().Errorw(err.Error(), "func", "GetPositionById")
		return nil, err
	}

	return position, nil
}

// GetPositions implements Repository.
func (r *repository) GetPositions(ctx context.Context) ([]Position, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		r.logger.Sugar().Errorw(err.Error(), "func", "GetPositions")
		return nil, err
	}

	var positions []Position
	if err = cursor.All(ctx, &positions); err != nil {
		r.logger.Sugar().Errorw(err.Error(), "func", "GetPositions")
		return nil, err
	}

	return positions, nil
}

// GetSupervisorById implements Repository.
func (r *repository) GetSupervisorById(ctx context.Context, id string) (*Supervisor, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		r.logger.Sugar().Errorw(err.Error(), "func", "GetSupervisorById")
		return nil, err
	}

	supervisor := new(Supervisor)
	if err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(supervisor); err != nil {
		if err == mongo.ErrNoDocuments {
			r.logger.Sugar().Errorw("employee not found", "func", "GetSupervisorById")
			return nil, err
		}

		r.logger.Sugar().Errorw(err.Error(), "func", "GetSupervisorById")
		return nil, err
	}

	return supervisor, nil
}

// GetSupervisors implements Repository.
func (r *repository) GetSupervisors(ctx context.Context) ([]Supervisor, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		r.logger.Sugar().Errorw(err.Error(), "func", "GetSupervisors")
		return nil, err
	}

	var supervisors []Supervisor
	if err = cursor.All(ctx, &supervisors); err != nil {
		r.logger.Sugar().Errorw(err.Error(), "func", "GetSupervisors")
		return nil, err
	}

	return supervisors, nil
}

// UpdateEmployee implements Repository.
func (r *repository) UpdateEmployee(ctx context.Context, employee *Employee) error {
	_, err := r.collection.ReplaceOne(ctx, bson.M{"_id": employee.ID}, employee)
	if err != nil {
		r.logger.Sugar().Errorw(err.Error(), "func", "UpdateEmployee", "request_id", ctx.Value("request_id"))
		return err
	}

	return nil
}

// UpdateNationality implements Repository.
func (r *repository) UpdateNationality(ctx context.Context, nationality *Nationality) error {
	_, err := r.collection.ReplaceOne(ctx, bson.M{"_id": nationality.ID}, nationality)
	if err != nil {
		r.logger.Sugar().Errorw(err.Error(), "func", "UpdateNationality", "request_id", ctx.Value("request_id"))
		return err
	}

	return nil
}

// UpdatePosition implements Repository.
func (r *repository) UpdatePosition(ctx context.Context, position *Position) error {
	_, err := r.collection.ReplaceOne(ctx, bson.M{"_id": position.ID}, position)
	if err != nil {
		r.logger.Sugar().Errorw(err.Error(), "func", "UpdatePosition", "request_id", ctx.Value("request_id"))
		return err
	}

	return nil
}

// UpdateSupervisor implements Repository.
func (r *repository) UpdateSupervisor(ctx context.Context, supervisor *Supervisor) error {
	_, err := r.collection.ReplaceOne(ctx, bson.M{"_id": supervisor.ID}, supervisor)
	if err != nil {
		r.logger.Sugar().Errorw(err.Error(), "func", "UpdateSupervisor", "request_id", ctx.Value("request_id"))
		return err
	}

	return nil
}

// GetEmployeeByEmail implements Repository.
func (r *repository) GetEmployeeByEmail(ctx context.Context, email string) (*Employee, error) {
	employee := new(Employee)
	err := r.collection.FindOne(ctx, bson.M{"email": email}).Decode(employee)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}

		r.logger.Sugar().Errorw(err.Error(), "func", "GetEmployeeByEmail")
		return nil, err
	}

	return employee, nil
}

// GetNationalityByName implements Repository.
func (r *repository) GetNationalityByName(ctx context.Context, name string) (*Nationality, error) {
	nationality := new(Nationality)
	err := r.collection.FindOne(ctx, bson.M{"name": name}).Decode(nationality)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}

		r.logger.Sugar().Errorw(err.Error(), "func", "GetNationalityByName")
		return nil, err
	}

	return nationality, nil
}

// GetPositionByName implements Repository.
func (r *repository) GetPositionByName(ctx context.Context, name string) (*Position, error) {
	position := new(Position)
	err := r.collection.FindOne(ctx, bson.M{"name": name}).Decode(position)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}

		r.logger.Sugar().Errorw(err.Error(), "func", "GetPositionByName")
		return nil, err
	}

	return position, nil
}

// GetSupervisorByName implements Repository.
func (r *repository) GetSupervisorByName(ctx context.Context, name string) (*Supervisor, error) {
	supervisor := new(Supervisor)
	err := r.collection.FindOne(ctx, bson.M{"name": name}).Decode(supervisor)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}

		r.logger.Sugar().Errorw(err.Error(), "func", "GetSupervisorByName")
		return nil, err
	}

	return supervisor, nil
}

// GetEmployeeByIdentificationNumber implements Repository.
func (r *repository) GetEmployeeByIdentificationNumber(ctx context.Context, identificationNumber string) (*Employee, error) {
	employee := new(Employee)
	err := r.collection.FindOne(ctx, bson.M{"identification_number": identificationNumber}).Decode(employee)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}

		r.logger.Sugar().Errorw(err.Error(), "func", "GetEmployeeByIdentificationNumber")
		return nil, err
	}

	return employee, nil
}
