package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rejrak/url-shortener/routes"
	docs "github.com/rejrak/url-shortener/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func setupRoutes(r *gin.Engine) {
	r.GET("/:url", routes.ResolveURL)
	r.POST("/api/v1", routes.ShortenURL)
}

// @title SnM GO UrlShortener
// @version 1.0
// @description Just a Cherry.
// @termsOfService http://swagger.io/terms/
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /api/v1
// @schemes http
func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerfiles.Handler,
		ginSwagger.URL("http://localhost:3000/swagger/doc.json"),
	))
	router.GET("/api/v1/health-check", HealthCheck)

	router.Use(gin.Logger())

	setupRoutes(router)

	log.Fatal(router.Run(os.Getenv("APP_PORT")))

}

// HealthCheck
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /health-check [get]
func HealthCheck(c *gin.Context) {
	res := map[string]interface{}{
		"data": "Server is up and running",
	}
	c.JSON(http.StatusOK, res)
}
