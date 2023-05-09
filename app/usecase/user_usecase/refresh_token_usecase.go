package user_usecase

import (
	"context"
	"net/http"
	"time"

	"olimpo/app/domain"
	"olimpo/app/http/response"
	"olimpo/bootstrap"
	"olimpo/internal/tokenutil"

	"github.com/gin-gonic/gin"
)

type RefreshTokenUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
	env            *bootstrap.Env
}

func NewRefreshTokenUsecase(env *bootstrap.Env, userRepository domain.UserRepository, timeout time.Duration) *RefreshTokenUsecase {
	return &RefreshTokenUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
		env:            env,
	}
}

func (rtu *RefreshTokenUsecase) RefreshToken(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, rtu.contextTimeout)
	defer cancel()

	var request domain.RefreshTokenRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	userInformation, err := tokenutil.ExtractUserInformationFromToken(request.RefreshToken, rtu.env.RefreshTokenSecret)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.ErrorResponse{Message: "User not found"})
		return
	}

	user, err := rtu.userRepository.GetByID(ctx, userInformation.ID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, response.ErrorResponse{Message: "User not found"})
		return
	}

	accessToken, err := tokenutil.CreateAccessToken(&user, rtu.env.AccessTokenSecret, rtu.env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err := tokenutil.CreateRefreshToken(&user, rtu.env.RefreshTokenSecret, rtu.env.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	refreshTokenResponse := domain.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, refreshTokenResponse)
}
