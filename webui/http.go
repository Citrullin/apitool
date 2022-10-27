package webui

import (
	"fmt"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func Serve(port string) {
	router := gin.Default()
	router.Use(cors.Default())
	gin.SetMode(gin.ReleaseMode)

	router.GET("/", GetIndex)
	router.POST("/sign", Sign)
	router.POST("/send", Send)

	fmt.Printf("WebUI listening on http://127.0.0.1:%s\n", port)

	if err := router.Run(":" + port); err != nil {
		fmt.Println(err)
	}
}
