package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Go uses concurrency so it can handle many requests at once
	// creating a router

	r := gin.Default()
	r.Use(cors.Default())

	// use r.Group for grouping routes
	r.GET("/", HelloHandler)
	r.GET("/user", GetName)
	r.GET("/profile/:id", GetId)
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

// H is map[string]interface{}
// so value can be anything
// c.JSON writes it to the response handler

func HelloHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message":  true,
		"message2": "hello world",
	})
}

// localhost:9000/user?firstname=naveen&lastname=garg
func GetName(c *gin.Context) {
	firstName := c.Query("firstname")
	lastName := c.Query("lastname")

	c.JSON(http.StatusOK, gin.H{
		"firstname": firstName,
		"lastname":  lastName,
	})
}

func GetId(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
	// c.JSON(200, id)
}
