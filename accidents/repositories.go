package accidents

import (
	"context"

	"github.com/xbizzybone/pitagoras-api/employees"
	"github.com/xbizzybone/pitagoras-api/projects"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

type repository struct {
	logger                    *zap.Logger
	accidentsCollection       *mongo.Collection
	classificationsCollection *mongo.Collection
	gravitiesCollection       *mongo.Collection
	injuredPartsCollection    *mongo.Collection
	employeesCollection       *mongo.Collection
	proyectsCollection        *mongo.Collection
}

func NewRepository(logger *zap.Logger, accidentsCollection, classificationsCollection, gravitiesCollection, injuredPartsCollection, employeesCollection, proyectsCollection *mongo.Collection) Repository {
	return &repository{
		logger:                    logger,
		accidentsCollection:       accidentsCollection,
		classificationsCollection: classificationsCollection,
		gravitiesCollection:       gravitiesCollection,
		injuredPartsCollection:    injuredPartsCollection,
		employeesCollection:       employeesCollection,
		proyectsCollection:        proyectsCollection,
	}
}

// CreateAccident implements Repository.
func (r *repository) CreateAccident(ctx context.Context, accident *Accident) (*Accident, error) {
	result, err := r.accidentsCollection.InsertOne(ctx, accident)
	if err != nil {
		r.logger.Sugar().Errorw(err.Error(), "func", "CreateAccident", "request_id", ctx.Value("request_id"))
		return nil, err
	}

	accident.ID = result.InsertedID.(primitive.ObjectID)

	return accident, nil
}

// CreateClassification implements Repository.
func (r *repository) CreateClassification(ctx context.Context, classification *Classification) (*Classification, error) {
	result, err := r.classificationsCollection.InsertOne(ctx, classification)
	if err != nil {
		r.logger.Sugar().Errorw(err.Error(), "func", "CreateClassification", "request_id", ctx.Value("request_id"))
		return nil, err
	}

	classification.ID = result.InsertedID.(primitive.ObjectID)

	return classification, nil
}

// CreateGravity implements Repository.
func (r *repository) CreateGravity(ctx context.Context, gravity *Gravity) (*Gravity, error) {
	result, err := r.gravitiesCollection.InsertOne(ctx, gravity)
	if err != nil {
		r.logger.Sugar().Errorw(err.Error(), "func", "CreateGravity", "request_id", ctx.Value("request_id"))
		return nil, err
	}

	gravity.ID = result.InsertedID.(primitive.ObjectID)

	return gravity, nil
}

// CreateInjuredPart implements Repository.
func (r *repository) CreateInjuredPart(ctx context.Context, injuredPart *InjuredPart) (*InjuredPart, error) {
	result, err := r.injuredPartsCollection.InsertOne(ctx, injuredPart)
	if err != nil {
		r.logger.Sugar().Errorw(err.Error(), "func", "CreateInjuredPart", "request_id", ctx.Value("request_id"))
		return nil, err
	}

	injuredPart.ID = result.InsertedID.(primitive.ObjectID)

	return injuredPart, nil
}

// GetAccidentById implements Repository.
func (r *repository) GetAccidentById(ctx context.Context, id primitive.ObjectID) (*Accident, error) {
	accident := new(Accident)
	if err := r.accidentsCollection.FindOne(ctx, primitive.M{"_id": id}).Decode(accident); err != nil {
		if err == mongo.ErrNoDocuments {
			r.logger.Sugar().Errorw("accident not found", "func", "GetAccidentById", "request_id", ctx.Value("request_id"))
			return nil, err
		}

		r.logger.Sugar().Errorw(err.Error(), "func", "GetAccidentById", "request_id", ctx.Value("request_id"))
		return nil, err
	}

	return accident, nil
}

// GetAccidents implements Repository.
func (r *repository) GetAccidents(ctx context.Context) ([]Accident, error) {
	cursor, err := r.accidentsCollection.Find(ctx, primitive.M{})
	if err != nil {
		r.logger.Sugar().Errorw(err.Error(), "func", "GetAccidents", "request_id", ctx.Value("request_id"))
		return nil, err
	}
	defer cursor.Close(ctx)

	accidents := make([]Accident, 0)
	if err := cursor.All(ctx, &accidents); err != nil {
		r.logger.Sugar().Errorw(err.Error(), "func", "GetAccidents", "request_id", ctx.Value("request_id"))
		return nil, err
	}

	return accidents, nil
}

// GetClassificationById implements Repository.
func (r *repository) GetClassificationById(ctx context.Context, id primitive.ObjectID) (*Classification, error) {
	classification := new(Classification)
	if err := r.classificationsCollection.FindOne(ctx, primitive.M{"_id": id}).Decode(classification); err != nil {
		if err == mongo.ErrNoDocuments {
			r.logger.Sugar().Errorw("classification not found", "func", "GetClassificationById", "request_id", ctx.Value("request_id"))
			return nil, err
		}

		r.logger.Sugar().Errorw(err.Error(), "func", "GetClassificationById", "request_id", ctx.Value("request_id"))
		return nil, err
	}

	return classification, nil
}

func (r *repository) GetClassificationByName(ctx context.Context, name string) (*Classification, error) {
	classification := new(Classification)
	if err := r.classificationsCollection.FindOne(ctx, primitive.M{"name": name}).Decode(classification); err != nil {
		if err == mongo.ErrNoDocuments {
			r.logger.Sugar().Errorw("classification not found", "func", "GetClassificationByName", "request_id", ctx.Value("request_id"))
			return nil, err
		}

		r.logger.Sugar().Errorw(err.Error(), "func", "GetClassificationByName", "request_id", ctx.Value("request_id"))
		return nil, err
	}

	return classification, nil
}

// GetClassifications implements Repository.
func (r *repository) GetClassifications(ctx context.Context) ([]Classification, error) {
	cursor, err := r.classificationsCollection.Find(ctx, primitive.M{})
	if err != nil {
		r.logger.Sugar().Errorw(err.Error(), "func", "GetClassifications", "request_id", ctx.Value("request_id"))
		return nil, err
	}
	defer cursor.Close(ctx)

	classifications := make([]Classification, 0)
	if err := cursor.All(ctx, &classifications); err != nil {
		r.logger.Sugar().Errorw(err.Error(), "func", "GetClassifications", "request_id", ctx.Value("request_id"))
		return nil, err
	}

	return classifications, nil
}

// GetGravities implements Repository.
func (r *repository) GetGravities(ctx context.Context) ([]Gravity, error) {
	cursor, err := r.gravitiesCollection.Find(ctx, primitive.M{})
	if err != nil {
		r.logger.Sugar().Errorw(err.Error(), "func", "GetGravities", "request_id", ctx.Value("request_id"))
		return nil, err
	}
	defer cursor.Close(ctx)

	gravities := make([]Gravity, 0)
	if err := cursor.All(ctx, &gravities); err != nil {
		r.logger.Sugar().Errorw(err.Error(), "func", "GetGravities", "request_id", ctx.Value("request_id"))
		return nil, err
	}

	return gravities, nil
}

// GetGravityById implements Repository.
func (r *repository) GetGravityById(ctx context.Context, id primitive.ObjectID) (*Gravity, error) {
	gravity := new(Gravity)
	if err := r.gravitiesCollection.FindOne(ctx, primitive.M{"_id": id}).Decode(gravity); err != nil {
		if err == mongo.ErrNoDocuments {
			r.logger.Sugar().Errorw("gravity not found", "func", "GetGravityById", "request_id", ctx.Value("request_id"))
			return nil, err
		}

		r.logger.Sugar().Errorw(err.Error(), "func", "GetGravityById", "request_id", ctx.Value("request_id"))
		return nil, err
	}

	return gravity, nil
}

func (r *repository) GetGravityByName(ctx context.Context, name string) (*Gravity, error) {
	gravity := new(Gravity)
	if err := r.gravitiesCollection.FindOne(ctx, primitive.M{"name": name}).Decode(gravity); err != nil {
		if err == mongo.ErrNoDocuments {
			r.logger.Sugar().Errorw("gravity not found", "func", "GetGravityByName", "request_id", ctx.Value("request_id"))
			return nil, err
		}

		r.logger.Sugar().Errorw(err.Error(), "func", "GetGravityByName", "request_id", ctx.Value("request_id"))
		return nil, err
	}

	return gravity, nil
}

// GetInjuredPartById implements Repository.
func (r *repository) GetInjuredPartById(ctx context.Context, id primitive.ObjectID) (*InjuredPart, error) {
	injuredPart := new(InjuredPart)
	if err := r.injuredPartsCollection.FindOne(ctx, primitive.M{"_id": id}).Decode(injuredPart); err != nil {
		if err == mongo.ErrNoDocuments {
			r.logger.Sugar().Errorw("injured part not found", "func", "GetInjuredPartById", "request_id", ctx.Value("request_id"))
			return nil, err
		}

		r.logger.Sugar().Errorw(err.Error(), "func", "GetInjuredPartById", "request_id", ctx.Value("request_id"))
		return nil, err
	}

	return injuredPart, nil
}

func (r *repository) GetInjuredPartByName(ctx context.Context, name string) (*InjuredPart, error) {
	injuredPart := new(InjuredPart)
	if err := r.injuredPartsCollection.FindOne(ctx, primitive.M{"name": name}).Decode(injuredPart); err != nil {
		if err == mongo.ErrNoDocuments {
			r.logger.Sugar().Errorw("injured part not found", "func", "GetInjuredPartByName", "request_id", ctx.Value("request_id"))
			return nil, err
		}

		r.logger.Sugar().Errorw(err.Error(), "func", "GetInjuredPartByName", "request_id", ctx.Value("request_id"))
		return nil, err
	}

	return injuredPart, nil
}

// GetInjuredParts implements Repository.
func (r *repository) GetInjuredParts(ctx context.Context) ([]InjuredPart, error) {
	cursor, err := r.injuredPartsCollection.Find(ctx, primitive.M{})
	if err != nil {
		r.logger.Sugar().Errorw(err.Error(), "func", "GetInjuredParts", "request_id", ctx.Value("request_id"))
		return nil, err
	}
	defer cursor.Close(ctx)

	injuredParts := make([]InjuredPart, 0)
	if err := cursor.All(ctx, &injuredParts); err != nil {
		r.logger.Sugar().Errorw(err.Error(), "func", "GetInjuredParts", "request_id", ctx.Value("request_id"))
		return nil, err
	}

	return injuredParts, nil
}

// UpdateAccident implements Repository.
func (r *repository) UpdateAccident(ctx context.Context, accident *Accident) error {
	_, err := r.accidentsCollection.ReplaceOne(ctx, primitive.M{"_id": accident.ID}, accident)
	if err != nil {
		r.logger.Sugar().Errorw(err.Error(), "func", "UpdateAccident", "request_id", ctx.Value("request_id"))
		return err
	}

	return nil
}

// UpdateClassification implements Repository.
func (r *repository) UpdateClassification(ctx context.Context, classification *Classification) error {
	_, err := r.classificationsCollection.ReplaceOne(ctx, primitive.M{"_id": classification.ID}, classification)
	if err != nil {
		r.logger.Sugar().Errorw(err.Error(), "func", "UpdateClassification", "request_id", ctx.Value("request_id"))
		return err
	}

	return nil
}

// UpdateGravity implements Repository.
func (r *repository) UpdateGravity(ctx context.Context, gravity *Gravity) error {
	_, err := r.gravitiesCollection.ReplaceOne(ctx, primitive.M{"_id": gravity.ID}, gravity)
	if err != nil {
		r.logger.Sugar().Errorw(err.Error(), "func", "UpdateGravity", "request_id", ctx.Value("request_id"))
		return err
	}

	return nil
}

// UpdateInjuredPart implements Repository.
func (r *repository) UpdateInjuredPart(ctx context.Context, injuredPart *InjuredPart) error {
	_, err := r.injuredPartsCollection.ReplaceOne(ctx, primitive.M{"_id": injuredPart.ID}, injuredPart)
	if err != nil {
		r.logger.Sugar().Errorw(err.Error(), "func", "UpdateInjuredPart", "request_id", ctx.Value("request_id"))
		return err
	}

	return nil
}

func (r *repository) GetEmployeeById(ctx context.Context, id primitive.ObjectID) (*employees.Employee, error) {
	employee := new(employees.Employee)
	if err := r.employeesCollection.FindOne(ctx, bson.M{"_id": id}).Decode(employee); err != nil {
		if err == mongo.ErrNoDocuments {
			r.logger.Sugar().Errorw("employee not found", "func", "GetEmployeeById")
			return nil, err
		}

		r.logger.Sugar().Errorw(err.Error(), "func", "GetEmployeeById")
		return nil, err
	}

	return employee, nil
}

func (r *repository) GetProjectById(ctx context.Context, id primitive.ObjectID) (*projects.Project, error) {
	project := new(projects.Project)
	if err := r.proyectsCollection.FindOne(ctx, bson.M{"_id": id}).Decode(project); err != nil {
		if err == mongo.ErrNoDocuments {
			r.logger.Sugar().Errorw("project not found", "func", "GetProjectById")
			return nil, err
		}

		r.logger.Sugar().Errorw(err.Error(), "func", "GetProjectById")
		return nil, err
	}

	return project, nil
}
