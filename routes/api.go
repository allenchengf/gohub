package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		// 註冊一個路由
		v1.GET("/", func(c *gin.Context) {
			// 以 JSON 格式響應
			c.JSON(http.StatusOK, gin.H{
				"Hello": "World!",
			})
		})
	}
}
