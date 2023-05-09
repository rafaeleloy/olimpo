package org_usecase

import (
	"context"
	"net/http"
	"reflect"
	"time"

	"olimpo/app/domain"
	"olimpo/app/http/response"
	"olimpo/bootstrap"

	"github.com/gin-gonic/gin"
)

type GetOrgByIDUseCase struct {
	orgRepository  domain.OrgRepository
	contextTimeout time.Duration
	env            *bootstrap.Env
}

func NewGetOrgByIDUseCase(env *bootstrap.Env, orgRepository domain.OrgRepository, timeout time.Duration) *GetOrgByIDUseCase {
	return &GetOrgByIDUseCase{
		orgRepository:  orgRepository,
		contextTimeout: timeout,
		env:            env,
	}
}

func (cou *GetOrgByIDUseCase) GetOrgByID(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, cou.contextTimeout)
	defer cancel()

	orgID := c.Param("id")

	org, err := cou.orgRepository.GetByID(ctx, orgID)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	} else if reflect.ValueOf(org).IsZero() && err == nil {
		c.JSON(http.StatusNotFound, response.ErrorResponse{Message: "Org not found"})
	}

	orgResponse := domain.GetOrgByIDResponse{
		Org: org,
	}

	c.JSON(http.StatusCreated, orgResponse)
}
