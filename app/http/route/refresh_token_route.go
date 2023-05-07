package route

import (
	"time"

	"olimpo/app/domain"
	"olimpo/app/repository"
	"olimpo/app/usecase"
	"olimpo/bootstrap"
	"olimpo/infra/database"

	"github.com/gin-gonic/gin"
)

func NewRefreshTokenRouter(env *bootstrap.Env, timeout time.Duration, db database.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)

	rtc := usecase.NewRefreshTokenUsecase(env, ur, timeout)
	group.POST("/refresh", rtc.RefreshToken)
}
