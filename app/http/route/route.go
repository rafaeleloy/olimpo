package route

import (
	"time"

	"olimpo/bootstrap"

	"olimpo/app/http/middleware"
	"olimpo/app/http/route/org_route"
	"olimpo/app/http/route/user_route"

	"olimpo/infra/database"

	"github.com/gin-gonic/gin"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db database.Database, gin *gin.Engine) {
	publicRouter := gin.Group("")
	// All Public APIs
	NewSignupRouter(env, timeout, db, publicRouter)
	NewLoginRouter(env, timeout, db, publicRouter)
	user_route.NewRefreshTokenRouter(env, timeout, db, publicRouter)

	// TODO: Add functions to get, update and delete orgs,
	// do the same for users and campaigns
	protectedRouterOrgAdmin := gin.Group("")
	protectedRouterOrgAdmin.Use(
		middleware.JwtAuthMiddleware(env.AccessTokenSecret),
		middleware.IsOrgAdminMiddleware(),
	)
	user_route.NewProfileRouter(env, timeout, db, protectedRouterOrgAdmin)
	org_route.NewCreateOrgRouter(env, timeout, db, protectedRouterOrgAdmin)
}
