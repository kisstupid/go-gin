package middlewares

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kisstupid/go-gin/internal/database"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()

		var greeting string
		db.QueryRow(context.Background(), "select 'Hello, word!'").Scan(&greeting)

		fmt.Println(greeting)

		c.Next()
	}
}
