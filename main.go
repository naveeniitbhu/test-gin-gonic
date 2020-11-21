package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Go uses concurrency so it can handle many requests at once
	// creating a router

	r := gin.Default()
	r.Use(cors.Default())

	// use r.Group for grouping routes
	r.GET("/", getHandler)
	// r.GET("/user", func(c *gin.Context) {
	// 	firstname := c.Query("firstname")
	// 	lastname := c.Query("lastname")

	// 	c.JSON(200, gin.H{"firstname": firstname,
	// 		"lastname": lastname})
	// })

	if err := r.Run(":9000"); err != nil {
		log.Fatalln(err)
	}
}

func getHandler(c *gin.Context) {
	// H is map[string]interface{}
	// so value can be anything
	// c.JSON writes it to the response handler

	c.JSON(200, gin.H{
		"message":  true,
		"message2": "hello world",
	})
}
