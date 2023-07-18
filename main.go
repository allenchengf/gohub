package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {

	//new一個Gin Engine實例
	r := gin.New()

	//註冊中間件
	r.Use(gin.Logger(), gin.Recovery())

	//註冊一個路由
	r.GET("/", func(c *gin.Context) {
		//以JSON格式響應
		c.JSON(http.StatusOK, gin.H{
			"Hello": "World!",
		})
	})

	//處理404請求
	r.NoRoute(func(c *gin.Context) {
		//獲取標頭信息的Accept信息
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			//如果是HTML的話
			c.String(http.StatusNotFound, "頁面返回404")
		} else {
			//返回JSON
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定義，請確認url和請求方法是否正確",
			})
		}
	})

	// 運行服務,指定port 8000
	r.Run(":8000")
}
