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
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type SignupUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
	env            *bootstrap.Env
}

func NewSignupUsecase(env *bootstrap.Env, userRepository domain.UserRepository, timeout time.Duration) *SignupUsecase {
	return &SignupUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
		env:            env,
	}
}

func (su *SignupUsecase) Signup(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()

	var request domain.SignupRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: err.Error()})
		return
	}

	_, err = su.userRepository.GetByEmail(ctx, request.Email)
	if err == nil {
		c.JSON(http.StatusConflict, response.ErrorResponse{Message: "User already exists with the given email"})
		return
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	request.Password = string(encryptedPassword)

	user := domain.User{
		ID:       primitive.NewObjectID(),
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	err = su.userRepository.Create(ctx, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	accessToken, err := tokenutil.CreateAccessToken(&user, su.env.AccessTokenSecret, su.env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err := tokenutil.CreateRefreshToken(&user, su.env.RefreshTokenSecret, su.env.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: err.Error()})
		return
	}

	signupResponse := domain.SignupResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, signupResponse)
}
