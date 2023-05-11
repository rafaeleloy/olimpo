package org_route

import (
	"time"

	"olimpo/bootstrap"
	"olimpo/infra/database"

	"olimpo/app/domain"
	"olimpo/app/repository"

	usecase "olimpo/app/usecase/org_usecase"

	"github.com/gin-gonic/gin"
)

func NewGetMyOrgRouter(env *bootstrap.Env, timeout time.Duration, db database.Database, group *gin.RouterGroup) {
	or := repository.NewOrgRepository(db, domain.CollectionOrg)
	ur := repository.NewUserRepository(db, domain.CollectionOrg)

	ou := usecase.NewGetMyOrgUseCase(env, or, ur, timeout)
	group.GET("/org", ou.GetMyOrg)
}
