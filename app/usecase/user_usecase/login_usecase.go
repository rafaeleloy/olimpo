package user_usecase

import (
	"context"
	"net/http"
	"time"

	"olimpo/app/domain"
	"olimpo/app/http/response"
	"olimpo/bootstrap"
	"olimpo/internal/tokenutil"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

type LoginUseCase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
	env            *bootstrap.Env
}

func NewLoginUsecase(env *bootstrap.Env, userRepository domain.UserRepository, timeout time.Duration) *LoginUseCase {
	return &LoginUseCase{
		userRepository: userRepository,
		contextTimeout: timeout,
		env:            env,
	}
}

func (lu *LoginUseCase) Login(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()

	var request domain.LoginRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	user, err := lu.userRepository.GetByEmail(ctx, request.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, response.ErrorResponse{Message: "User not found with the given email"})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		c.JSON(http.StatusUnauthorized, response.ErrorResponse{Message: "Invalid credentials"})
		return
	}

	accessToken, err := tokenutil.CreateAccessToken(&user, lu.env.AccessTokenSecret, lu.env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err := tokenutil.CreateRefreshToken(&user, lu.env.RefreshTokenSecret, lu.env.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	loginResponse := domain.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, loginResponse)
}
