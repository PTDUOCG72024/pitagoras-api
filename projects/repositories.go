package projects

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

func (r *repository) CreateProject(ctx context.Context, project *Project) error {
	_, err := r.collection.InsertOne(ctx, project)
	if err != nil {
		r.logger.Sugar().Errorw(err.Error(), "func", "CreateProject", "request_id", ctx.Value("request_id"))
		return err
	}

	return nil
}

func (r *repository) GetProjectById(ctx context.Context, id string) (Project, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		r.logger.Sugar().Errorw(err.Error(), "func", "GetProjectById")
		return Project{}, err
	}

	project := new(Project)
	if err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(project); err != nil {
		if err == mongo.ErrNoDocuments {
			r.logger.Sugar().Errorw("project not found", "func", "GetProjectById")
			return Project{}, err
		}

		r.logger.Sugar().Errorw(err.Error(), "func", "GetProjectById")
		return Project{}, err
	}

	return *project, nil
}

func (r *repository) GetProjects(ctx context.Context) ([]Project, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		r.logger.Sugar().Errorw(err.Error(), "func", "GetProjects")
		return nil, err
	}

	var projects []Project
	if err = cursor.All(ctx, &projects); err != nil {
		r.logger.Sugar().Errorw(err.Error(), "func", "GetProjects")
		return nil, err
	}

	return projects, nil
}

func (r *repository) UpdateProject(ctx context.Context, project *Project) error {
	_, err := r.collection.ReplaceOne(ctx, bson.M{"_id": project.ID}, project)
	if err != nil {
		r.logger.Sugar().Errorw(err.Error(), "func", "UpdateProject", "request_id", ctx.Value("request_id"))
		return err
	}

	return nil
}
