package org_usecase

import (
	"context"
	"net/http"
	"time"

	"olimpo/app/domain"
	"olimpo/app/http/response"
	"olimpo/bootstrap"

	"github.com/gin-gonic/gin"
)

type UpdateOrgNameUseCase struct {
	orgRepository  domain.OrgRepository
	contextTimeout time.Duration
	env            *bootstrap.Env
}

func NewUpdateOrgNameUseCase(env *bootstrap.Env, orgRepository domain.OrgRepository, timeout time.Duration) *UpdateOrgNameUseCase {
	return &UpdateOrgNameUseCase{
		orgRepository:  orgRepository,
		contextTimeout: timeout,
		env:            env,
	}
}

func (cou *UpdateOrgNameUseCase) UpdateOrgName(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, cou.contextTimeout)
	defer cancel()

	var request domain.UpdateOrgNameRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	_, err = cou.orgRepository.GetByName(ctx, request.Name)
	if err == nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: "Org already exists with the given name"})
	}

	err = cou.orgRepository.UpdateOrgName(ctx, request.OrgID, request.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}
