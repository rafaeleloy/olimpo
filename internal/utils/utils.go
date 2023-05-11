package utils

import (
	"olimpo/internal/tokenutil"

	"github.com/gin-gonic/gin"
)

func GetLoggedUser(c *gin.Context) *tokenutil.UserInformation {
	value, _ := c.Get("x-user-info")
	return value.(*tokenutil.UserInformation)
}
