package controllers

import (
	"fmt"
	"log"
	"vendor/internal/models/requests"
	service "vendor/internal/services"
)

type DockerController struct {
	dockerService service.IDockerService
}

type IDockerController interface {
	CreateWordpressService(requests.CreateWordpressServiceRequest)
}

func NewDockerController(dockerService service.IDockerService) IDockerController {
	return &DockerController{
		dockerService: dockerService,
	}
}

func (dockerController *DockerController) CreateWordpressService(createWordpressServiceRequest requests.CreateWordpressServiceRequest) {
	containerID, err := dockerController.dockerService.FindContainer(createWordpressServiceRequest.ContainerName)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ContainerID: ", containerID)
		dockerController.dockerService.RemoveContainer(containerID)
	}

	_, err = dockerController.dockerService.FindVolume(createWordpressServiceRequest.VolumeName)

	if err != nil {
		log.Fatalln(err)
	}

	networkId, err := dockerController.dockerService.FindNetwork(createWordpressServiceRequest.NetworkName)

	if err != nil {
		panic(err)
	}

	dockerController.dockerService.RunContainer(createWordpressServiceRequest, networkId)
}
