package main

import (
	"crudwithgolang/item"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var appName = "App Example"

func main() {
	var router = gin.Default()
	router.Use(cors.New(cors.Config{
		AllowMethods:    []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Authorization", "Content-type"},
		AllowAllOrigins: true,
	}))

	router.POST("/", item.AddItem)
	router.GET("/", item.GetItem)
	router.GET("/:id", item.GetItemBySlug)
	router.PATCH("/:id", item.EditItemBySlug)
	router.DELETE("/:id", item.DeleteItemBySlug)

	router.Run(os.Getenv("BIND_ADDR")) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
