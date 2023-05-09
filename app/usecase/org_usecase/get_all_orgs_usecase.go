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

type GetAllOrgsUseCase struct {
	orgRepository  domain.OrgRepository
	contextTimeout time.Duration
	env            *bootstrap.Env
}

func NewGetAllOrgsUseCase(env *bootstrap.Env, orgRepository domain.OrgRepository, timeout time.Duration) *GetAllOrgsUseCase {
	return &GetAllOrgsUseCase{
		orgRepository:  orgRepository,
		contextTimeout: timeout,
		env:            env,
	}
}

func (cou *GetAllOrgsUseCase) GetAllOrgs(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, cou.contextTimeout)
	defer cancel()

	orgs, err := cou.orgRepository.Fetch(ctx)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	orgsResponse := domain.GetAllOrgsResponse{
		Orgs: orgs,
	}

	c.JSON(http.StatusCreated, orgsResponse)
}
