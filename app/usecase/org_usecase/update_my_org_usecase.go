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

type UpdateMyOrgUseCase struct {
	orgRepository  domain.OrgRepository
	userRepository domain.UserRepository
	contextTimeout time.Duration
	env            *bootstrap.Env
}

func NewUpdateMyOrgUseCase(env *bootstrap.Env, orgRepository domain.OrgRepository, userRepository domain.UserRepository, timeout time.Duration) *UpdateMyOrgUseCase {
	return &UpdateMyOrgUseCase{
		orgRepository:  orgRepository,
		userRepository: userRepository,
		contextTimeout: timeout,
		env:            env,
	}
}

func (cou *UpdateMyOrgUseCase) UpdateOrgName(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, cou.contextTimeout)
	defer cancel()

	var request domain.UpdateOrgRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	loggedUserID := utils.GetLoggedUser(c).ID
	loggedUser, err := cou.userRepository.GetByID(ctx, loggedUserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	} else if reflect.ValueOf(loggedUser).IsZero() && err == nil {
		c.JSON(http.StatusUnauthorized, response.ErrorResponse{Message: "Unauthorized"})
	}

	err = cou.orgRepository.Update(ctx, loggedUser.OrgID.String(), request)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
	}

	c.Status(http.StatusOK)
}
