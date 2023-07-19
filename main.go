package main

import (
	"fmt"
	"gohub/bootstrap"

	"github.com/gin-gonic/gin"
)

func main() {

	// new 一個 Gin Engine 實例
	router := gin.New()

	// 初始化路由綁定
	bootstrap.SetupRoute(router)

	// 運行服務
	err := router.Run(":3000")
	if err != nil {
		// 錯誤處理
		fmt.Println(err.Error())
	}
}
