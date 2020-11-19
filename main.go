package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// creating a router
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/user", func(c *gin.Context) {
		firstname := c.Query("firstname")
		lastname := c.Query("lastname")

		c.JSON(200, gin.H{"firstname": firstname,
			"lastname": lastname})
	})

	r.Run(":9000")
}
