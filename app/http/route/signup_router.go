package route

import (
	"time"

	"olimpo/app/domain"
	"olimpo/app/repository"
	usecase "olimpo/app/usecase/user_usecase"
	"olimpo/bootstrap"
	"olimpo/infra/database"

	"github.com/gin-gonic/gin"
)

func NewSignupRouter(env *bootstrap.Env, timeout time.Duration, db database.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)

	sc := usecase.NewSignupUsecase(env, ur, timeout)
	group.POST("/signup", sc.Signup)
}
