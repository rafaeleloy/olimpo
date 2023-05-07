package route

import (
	"olimpo/app/domain"
	"olimpo/app/repository"
	"olimpo/app/usecase"
	"olimpo/bootstrap"
	"olimpo/infra/database"
	"time"

	"github.com/gin-gonic/gin"
)

func NewProfileRouter(env *bootstrap.Env, timeout time.Duration, db database.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)

	pu := usecase.NewProfileUsecase(env, ur, timeout)
	group.POST("/set-user-profile", pu.SetUserProfile)
}
