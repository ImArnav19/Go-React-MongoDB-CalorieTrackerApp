package main

import (
	"os"

	"github.com/ImArnav19/go-react-calorieTracker/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(cors.Default())

	router.POST("/entry/create", routes.AddEntry)

	router.GET("/entries", routes.AllEntry)
	router.GET("/entry/:id", routes.GetEntry)
	router.GET("/ingredient/:id", routes.GetEntriesByIngredient)

	router.PUT("/entry/update/:id", routes.UpdateEntry)
	router.PUT("/ingredients/update/:id", routes.UpdateIngredient)

	router.DELETE("/delete/:id", routes.DeleteEntry)
	router.Run(":" + port)

}
