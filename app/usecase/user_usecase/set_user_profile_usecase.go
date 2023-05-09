package user_usecase

import (
	"context"
	"net/http"
	"olimpo/app/domain"
	"olimpo/app/http/response"
	"olimpo/bootstrap"
	"time"

	"github.com/gin-gonic/gin"
)

type ProfileUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
	env            *bootstrap.Env
}

func NewProfileUsecase(env *bootstrap.Env, userRepository domain.UserRepository, timeout time.Duration) *ProfileUsecase {
	return &ProfileUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
		env:            env,
	}
}

func (pu *ProfileUsecase) SetUserProfile(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()

	var request domain.ProfileRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	err = pu.userRepository.SetUserProfile(ctx, request.UserID, request.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
