package route

import (
	"time"

	"olimpo/bootstrap"

	"olimpo/app/http/middleware"

	"olimpo/infra/database"

	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db database.Database, gin *gin.Engine) {
	publicRouter := gin.Group("")
	// All Public APIs
	NewSignupRouter(env, timeout, db, publicRouter)
	NewLoginRouter(env, timeout, db, publicRouter)
	NewRefreshTokenRouter(env, timeout, db, publicRouter)

	// TODO: Add functions to get, update and delete orgs,
	// do the same for users and campaigns
	protectedRouterOrgAdmin := gin.Group("")
	protectedRouterOrgAdmin.Use(
		middleware.JwtAuthMiddleware(env.AccessTokenSecret),
		middleware.IsOrgAdminMiddleware(),
	)
	NewProfileRouter(env, timeout, db, protectedRouterOrgAdmin)
	NewCreateOrgRouter(env, timeout, db, protectedRouterOrgAdmin)
}
