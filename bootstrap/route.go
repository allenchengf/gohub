package bootstrap

import (
	"gohub/routes"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// SetupRoute路由初始化
func SetupRoute(router *gin.Engine) {
	//註冊全局中間件
	registerGlobalMiddleWare(router)

	//註冊API路由
	routes.RegisterAPIRoutes(router)

	//配置404路由
	setup404Handler(router)
}

func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(
		gin.Logger(),
		gin.Recovery(),
	)
}

func setup404Handler(router *gin.Engine) {
	//處理404請求
	router.NoRoute(func(c *gin.Context) {
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			//如果是HTML
			c.String(http.StatusNotFound, "頁面返回404")
		} else {
			// 默認返回 JSON
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定義，請確認url和請求方法是否正確",
			})
		}
	})
}
