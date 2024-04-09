package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	r := gin.Default()
	r.LoadHTMLGlob("./internal/views/*")
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, world"})
	})

	err := r.Run(":8080")
	if err != nil {
		fmt.Printf("server run err:%s\n", err.Error())
	}
}
