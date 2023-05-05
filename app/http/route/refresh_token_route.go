package route

import (
	"time"

	"olimpo/app/domain"
	"olimpo/app/http/controller"
	"olimpo/app/repository"
	"olimpo/app/usecase"
	"olimpo/bootstrap"
	"olimpo/infra/database"

	"github.com/gin-gonic/gin"
)

func NewRefreshTokenRouter(env *bootstrap.Env, timeout time.Duration, db database.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	rtc := &controller.RefreshTokenController{
		RefreshTokenUsecase: usecase.NewRefreshTokenUsecase(ur, timeout),
		Env:                 env,
	}
	group.POST("/refresh", rtc.RefreshToken)
}
