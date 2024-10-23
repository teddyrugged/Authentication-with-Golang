package middleware

import (
	"fmt"
	_ "net/http"
	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Authenticating request")
	}
}

