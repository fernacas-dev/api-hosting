package main

import (
	"api-hosting/internal/controllers"
	service "api-hosting/internal/services"
	"context"

	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
)

func main() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	dockerService := service.NewDockerService(ctx, cli)
	dockerController := controllers.NewDockerController(dockerService)

	r := gin.Default()
	r.GET("/healthcheck", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "OK",
		})
	})
	r.GET("/ping", dockerController.CreateWordpressService)
	r.Run("0.0.0.0:8088") // listen and serve on 0.0.0.0:8080
}
