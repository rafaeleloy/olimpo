package middleware

import (
	"net/http"
	"strings"

	"olimpo/app/domain"
	"olimpo/app/http/response"

	"olimpo/internal/tokenutil"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) == 2 {
			authToken := t[1]
			authorized, err := tokenutil.IsAuthorized(authToken, secret)
			if authorized {
				userInformation, err := tokenutil.ExtractUserInformationFromToken(authToken, secret)
				if err != nil {
					c.JSON(http.StatusUnauthorized, response.ErrorResponse{Message: err.Error()})
					c.Abort()
					return
				}
				c.Set("x-user-info", userInformation)
				c.Next()
				return
			}
			c.JSON(http.StatusUnauthorized, response.ErrorResponse{Message: err.Error()})
			c.Abort()
			return
		}
		c.JSON(http.StatusUnauthorized, response.ErrorResponse{Message: "Not authorized"})
		c.Abort()
	}
}

func IsOrgAdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		value, exists := c.Get("x-user-info")
		if !exists {
			c.JSON(http.StatusUnauthorized, response.ErrorResponse{Message: "Not authorized"})
			c.Abort()
		}

		userInformation := value.(*tokenutil.UserInformation)
		if userInformation.ProfileRole != domain.OrgAdmin {
			c.JSON(http.StatusUnauthorized, response.ErrorResponse{Message: "Not authorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}
