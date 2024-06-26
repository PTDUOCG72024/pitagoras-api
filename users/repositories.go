package users

import (
	"context"
	"errors"
	"time"

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

func (r *repository) CreateUser(ctx context.Context, user *User) error {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		r.logger.Sugar().Errorw(err.Error(), "func", "CreateUser", "request_id", ctx.Value("request_id"))
		return err
	}
	return nil
}

func (r *repository) GetUserById(ctx context.Context, id string) (User, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		r.logger.Sugar().Errorw(err.Error(), "func", "GetUserById")
		return User{}, errors.New("invalid id")
	}

	user := new(User)
	if err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(user); err != nil {
		if err == mongo.ErrNoDocuments {
			r.logger.Sugar().Errorw("user not found", "func", "GetUserById")
			return User{}, errors.New("user not found")
		}

		r.logger.Sugar().Errorw(err.Error(), "func", "GetUserById")
		return User{}, err
	}
	return *user, nil
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (User, error) {
	user := new(User)
	err := r.collection.FindOne(ctx, bson.M{"email": email}).Decode(user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return User{}, errors.New("user not found")
		}

		r.logger.Sugar().Errorw(err.Error(), "func", "GetUserByEmail")
		return User{}, err
	}
	return *user, nil
}

func (r *repository) UpdateUser(ctx context.Context, user *User) error {
	filter := bson.M{"_id": user.ID}
	update := bson.M{"$set": bson.M{
		"name":       user.Name,
		"email":      user.Email,
		"password":   user.Password,
		"last_login": user.LastLogin,
		"updated_at": primitive.NewDateTimeFromTime(time.Now()),
		"delete_at":  user.DeleteAt,
	}}
	_, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			r.logger.Sugar().Errorw("user not found", "func", "UpdateUser")
			return errors.New("user not found")
		}

		r.logger.Sugar().Errorw(err.Error(), "func", "UpdateUser")
		return errors.New("error while updating user")
	}

	return nil
}

func (r *repository) Deactivate(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		r.logger.Sugar().Errorw(err.Error(), "func", "Deactivate")
		return errors.New("invalid id")
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": bson.M{
		"is_deleted": true,
		"delete_at":  primitive.NewDateTimeFromTime(time.Now()),
	}}

	_, err = r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			r.logger.Sugar().Errorw("user not found", "func", "Deactivate")
			return errors.New("user not found")
		}

		r.logger.Sugar().Errorw(err.Error(), "func", "Deactivate")
		return errors.New("error while deactivating user")
	}

	return nil
}

func (r *repository) Activate(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		r.logger.Sugar().Errorw(err.Error(), "func", "Activate")
		return errors.New("invalid id")
	}

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": bson.M{
		"is_deleted": false,
		"delete_at":  primitive.NewDateTimeFromTime(time.Time{}),
		"updated_at": primitive.NewDateTimeFromTime(time.Now()),
	}}

	_, err = r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			r.logger.Sugar().Errorw("user not found", "func", "Activate")
			return errors.New("user not found")
		}

		r.logger.Sugar().Errorw(err.Error(), "func", "Activate")
		return errors.New("error while activating user")
	}

	return nil
}
