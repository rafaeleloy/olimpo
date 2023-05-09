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

func NewUpdateOrgNameRouter(env *bootstrap.Env, timeout time.Duration, db database.Database, group *gin.RouterGroup) {
	or := repository.NewOrgRepository(db, domain.CollectionOrg)

	ou := usecase.NewUpdateOrgNameUseCase(env, or, timeout)
	group.PUT("/org", ou.UpdateOrgName)
}
