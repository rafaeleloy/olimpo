package org_usecase

import (
	"context"
	"net/http"
	"reflect"
	"time"

	"olimpo/app/domain"
	"olimpo/app/http/response"
	"olimpo/bootstrap"
	"olimpo/internal/utils"

	"github.com/gin-gonic/gin"
)

type GetMyOrgUseCase struct {
	orgRepository  domain.OrgRepository
	userRepository domain.UserRepository
	contextTimeout time.Duration
	env            *bootstrap.Env
}

func NewGetMyOrgUseCase(env *bootstrap.Env, orgRepository domain.OrgRepository,
	userRepository domain.UserRepository, timeout time.Duration) *GetMyOrgUseCase {
	return &GetMyOrgUseCase{
		orgRepository:  orgRepository,
		userRepository: userRepository,
		contextTimeout: timeout,
		env:            env,
	}
}

func (cou *GetMyOrgUseCase) GetMyOrg(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, cou.contextTimeout)
	defer cancel()

	loggedUserID := utils.GetLoggedUser(c).ID
	loggedUser, err := cou.userRepository.GetByID(ctx, loggedUserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	} else if reflect.ValueOf(loggedUser).IsZero() && err == nil {
		c.JSON(http.StatusUnauthorized, response.ErrorResponse{Message: "Unauthorized"})
	}

	org, err := cou.orgRepository.GetByID(ctx, loggedUser.OrgID.String())
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	} else if reflect.ValueOf(loggedUser).IsZero() && err == nil {
		c.JSON(http.StatusNotFound, response.ErrorResponse{Message: "Org not found"})
	}

	myOrgResponse := domain.GetOrgResponse{
		Org: org,
	}

	c.JSON(http.StatusCreated, myOrgResponse)
}
