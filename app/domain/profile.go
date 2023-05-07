package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Role uint32

const (
	OrgAdmin Role = iota
	OrgCollaborator
)

type ProfileRequest struct {
	UserID primitive.ObjectID `json:"userID"`
	Role   `json:"role"`
}
