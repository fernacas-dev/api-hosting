package main

import (
	"api-hosting/internal/controllers"
	"api-hosting/internal/models/requests"
	service "api-hosting/internal/services"
	"context"

	"github.com/docker/docker/client"
)

func main() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	createWordpressServiceRequest := requests.CreateWordpressServiceRequest{
		ContainerImage: "wordpress",
		ContainerName:  "wordpress-web",
		VolumeName:     "wordpress-web",
		NetworkName:    "database_network",
	}

	dockerService := service.NewDockerService(ctx, cli)
	dockerController := controllers.NewDockerController(dockerService)
	dockerController.CreateWordpressService(createWordpressServiceRequest)
}
