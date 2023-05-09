package org_usecase

import (
	"context"
	"net/http"
	"time"

	"olimpo/app/domain"
	"olimpo/app/http/response"
	"olimpo/bootstrap"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
)

type CreateOrgUseCase struct {
	orgRepository  domain.OrgRepository
	contextTimeout time.Duration
	env            *bootstrap.Env
}

func NewCreateOrgUseCase(env *bootstrap.Env, orgRepository domain.OrgRepository, timeout time.Duration) *CreateOrgUseCase {
	return &CreateOrgUseCase{
		orgRepository:  orgRepository,
		contextTimeout: timeout,
		env:            env,
	}
}

func (cou *CreateOrgUseCase) CreateOrg(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, cou.contextTimeout)
	defer cancel()

	var request domain.CreateOrgRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	_, err = cou.orgRepository.GetByName(ctx, request.Name)
	if err == nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: "Org already exists with the given name"})
	}

	org := domain.Org{
		ID:          primitive.NewObjectID(),
		Name:        request.Name,
		UsersID:     request.UsersID,
		CampaignsID: request.CampaignsID,
	}

	err = cou.orgRepository.Create(ctx, &org)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}
