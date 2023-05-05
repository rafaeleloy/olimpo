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

func NewProfileRouter(env *bootstrap.Env, timeout time.Duration, db database.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	pc := &controller.ProfileController{
		ProfileUsecase: usecase.NewProfileUsecase(ur, timeout),
	}
	group.GET("/profile", pc.Fetch)
}
