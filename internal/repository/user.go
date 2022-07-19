package repository

import (
	"context"
	"errors"

	"github.com/amirmahdiamini/ebook-go/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	UserRepository interface {
		Insert(user *model.User) error
		UpdateOne(filter any, update any) error
		FindByUsername(username string) (*model.User, error)
		FindByPhone(phone string) (*model.User, error)
		VerifyAccount(phone string) error
		FindByUsernameOrPhone(username_or_phone string) (*model.User, error)
		FindBySID(username_or_phone string) (*model.User, error)
		Notification(phone string, data string) error
	}
	userRepository struct {
		database *mongo.Client
	}
)

func NewUserRepository(db *mongo.Client) UserRepository {
	return &userRepository{
		database: db,
	}
}
func (ctx *userRepository) Insert(user *model.User) error {
	collection := ctx.database.Database("vbook").Collection("user")
	_, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return errors.New("internal error #11")
	}
	return nil
}
func (ctx *userRepository) UpdateOne(filter any, update any) error {
	collection := ctx.database.Database("vbook").Collection("user")
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return errors.New("internal error #12")
	}
	return nil
}
func (ctx *userRepository) FindByUsername(username string) (*model.User, error) {
	var user *model.User
	collection := ctx.database.Database("vbook").Collection("user")
	res := collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if res != nil {
		return nil, errors.New("not found")
	}
	return user, nil
}
func (ctx *userRepository) FindByPhone(phone string) (*model.User, error) {
	var user *model.User
	collection := ctx.database.Database("vbook").Collection("user")
	res := collection.FindOne(context.Background(), bson.M{"phone": phone}).Decode(&user)
	if res != nil {
		return nil, errors.New("not found")
	}
	return user, nil
}

func (Ctx *userRepository) VerifyAccount(phone string) error {
	collection := Ctx.database.Database("vbook").Collection("user")
	_, err := collection.UpdateOne(context.Background(), bson.M{"phone": phone}, bson.M{"$set": bson.M{"is_verified": true}})
	if err != nil {
		return errors.New("we have an internal error #13")
	}
	return nil
}

func (ctx *userRepository) FindByUsernameOrPhone(username_or_phone string) (*model.User, error) {
	var user *model.User
	collection := ctx.database.Database("vbook").Collection("user")
	res := collection.FindOne(context.Background(), bson.M{"$or": []bson.M{bson.M{"username": username_or_phone}, bson.M{"phone": username_or_phone}}}).Decode(&user)
	if res != nil {
		return nil, errors.New("not found")
	}
	return user, nil
}
func (ctx *userRepository) FindBySID(sid string) (*model.User, error) {
	var user *model.User
	collection := ctx.database.Database("vbook").Collection("user")
	res := collection.FindOne(context.Background(), bson.M{"sid": sid}).Decode(&user)
	if res != nil {
		return nil, errors.New("not found")
	}
	return user, nil
}

func (ctx *userRepository) Notification(phone string, data string) error {
	collection := ctx.database.Database("vbook").Collection("notification")
	_, err := collection.UpdateOne(context.Background(), bson.M{"phone": phone}, bson.M{"$push": bson.M{"notifications": data}})
	if err != nil {
		return errors.New("we have an internal error #14")
	}
	return nil
}
