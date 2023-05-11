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

type UpdateOrgRequest struct {
	Name        string               `form:"name" binding:"required"`
	UsersID     []primitive.ObjectID `form:"users_id"`
	CampaignsID []primitive.ObjectID `form:"campaigns_id"`
}

type GetAllOrgsResponse struct {
	Orgs []Org `json:"orgs"`
}

type GetOrgResponse struct {
	Org `json:"orgs"`
}

type OrgRepository interface {
	Create(c context.Context, org *Org) error
	Fetch(c context.Context) ([]Org, error)
	GetByID(c context.Context, id string) (Org, error)
	GetByName(c context.Context, name string) (Org, error)
	Update(c context.Context, orgID string, newOrg UpdateOrgRequest) error
}
