package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionUser = "users"
)

type User struct {
	ID          primitive.ObjectID   `bson:"_id"`
	Name        string               `bson:"name"`
	Email       string               `bson:"email"`
	Password    string               `bson:"password"`
	OrgID       primitive.ObjectID   `bson:"org_id"`
	CampaignsID []primitive.ObjectID `bson:"campaigns_id"`
	ProfileRole Role                 `bson:"profile_role"`
}

type UserRepository interface {
	Create(c context.Context, user *User) error
	Fetch(c context.Context) ([]User, error)
	GetByEmail(c context.Context, email string) (User, error)
	SetUserProfile(c context.Context, userID primitive.ObjectID, profile Role) error
	GetByID(c context.Context, id string) (User, error)
}
