package user_route

import (
	"olimpo/app/domain"
	"olimpo/app/repository"
	usecase "olimpo/app/usecase/user_usecase"
	"olimpo/bootstrap"
	"olimpo/infra/database"
	"time"

	"github.com/gin-gonic/gin"
)

func NewProfileRouter(env *bootstrap.Env, timeout time.Duration, db database.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)

	pu := usecase.NewProfileUsecase(env, ur, timeout)
	group.POST("/user/set-profile", pu.SetUserProfile)
}
