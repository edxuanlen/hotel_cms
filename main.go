package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, Geektutu")
	})
	err := r.Run() // listen and serve on 0.0.0.0:8080
	fmt.Print(err)
	
}