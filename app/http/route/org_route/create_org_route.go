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

func NewCreateOrgRouter(env *bootstrap.Env, timeout time.Duration, db database.Database, group *gin.RouterGroup) {
	or := repository.NewOrgRepository(db, domain.CollectionOrg)

	lc := usecase.NewCreateOrgUseCase(env, or, timeout)
	group.POST("/org/create", lc.CreateOrg)
}
