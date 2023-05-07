package repository

import (
	"context"

	"olimpo/app/domain"
	"olimpo/infra/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type orgRepository struct {
	database   database.Database
	collection string
}

func NewOrgRepository(db database.Database, collection string) domain.OrgRepository {
	return &orgRepository{
		database:   db,
		collection: collection,
	}
}

func (or *orgRepository) Create(c context.Context, org *domain.Org) error {
	collection := or.database.Collection(or.collection)

	_, err := collection.InsertOne(c, org)

	return err
}

func (or *orgRepository) Fetch(c context.Context) ([]domain.Org, error) {
	collection := or.database.Collection(or.collection)

	cursor, err := collection.Find(c, bson.D{})

	if err != nil {
		return nil, err
	}

	var orgs []domain.Org

	err = cursor.All(c, &orgs)
	if orgs == nil {
		return []domain.Org{}, err
	}

	return orgs, err
}

func (or *orgRepository) GetByID(c context.Context, id string) (domain.Org, error) {
	collection := or.database.Collection(or.collection)

	var org domain.Org

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return org, err
	}

	err = collection.FindOne(c, bson.M{"_id": idHex}).Decode(&org)
	return org, err
}

func (or *orgRepository) GetByName(c context.Context, name string) (domain.Org, error) {
	collection := or.database.Collection(or.collection)

	var org domain.Org

	err := collection.FindOne(c, bson.M{"name": name}).Decode(&org)
	return org, err
}
