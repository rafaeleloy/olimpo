package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionOrg = "orgs"
)

type Org struct {
	ID          primitive.ObjectID   `bson:"_id"`
	Name        string               `bson:"name"`
	UsersID     []primitive.ObjectID `bson:"users_id"`
	CampaignsID []primitive.ObjectID `bson:"campaigns_id"`
}

type CreateOrgRequest struct {
	Name        string               `form:"name" binding:"required"`
	UsersID     []primitive.ObjectID `form:"users_id"`
	CampaignsID []primitive.ObjectID `form:"campaigns_id"`
}

type UpdateOrgNameRequest struct {
	Name  string             `form:"name" binding:"required"`
	OrgID primitive.ObjectID `form:"org_id" binding:"required"`
}

type GetAllOrgsResponse struct {
	Orgs []Org `json:"orgs"`
}

type GetOrgByIDResponse struct {
	Org `json:"orgs"`
}

type OrgRepository interface {
	Create(c context.Context, org *Org) error
	Fetch(c context.Context) ([]Org, error)
	GetByID(c context.Context, id string) (Org, error)
	GetByName(c context.Context, name string) (Org, error)
	UpdateOrgName(c context.Context, orgID primitive.ObjectID, name string) error
}
