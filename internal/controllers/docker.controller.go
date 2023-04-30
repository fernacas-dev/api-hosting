package controllers

import (
	"api-hosting/internal/models/requests"
	service "api-hosting/internal/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"strconv"
	"time"
)

type DockerController struct {
	dockerService service.IDockerService
}

type IDockerController interface {
	CreateWordpressService(c *gin.Context)
}

func NewDockerController(dockerService service.IDockerService) IDockerController {
	return &DockerController{
		dockerService: dockerService,
	}
}

func (dockerController *DockerController) CreateWordpressService(c *gin.Context) {
	createWordpressServiceRequest := requests.CreateWordpressServiceRequest{
		ContainerImage: c.PostForm("containerImage"),
		ContainerName:  c.PostForm("containerName"),
		VolumeName:     c.PostForm("volumeName"),
		NetworkName:    "database_network",
	}

	rand.Seed(time.Now().UnixNano())
	createWordpressServiceRequest.ContainerName = createWordpressServiceRequest.ContainerName + strconv.Itoa(rand.Intn(999999))
	createWordpressServiceRequest.VolumeName = createWordpressServiceRequest.ContainerName

	containerID, err := dockerController.dockerService.FindContainer(createWordpressServiceRequest.ContainerName)

	for i := 0; err == nil; i++ {
		containerID, err = dockerController.dockerService.FindContainer(createWordpressServiceRequest.ContainerName)
		createWordpressServiceRequest.ContainerName = createWordpressServiceRequest.ContainerName + strconv.Itoa(rand.Intn(999999))
		createWordpressServiceRequest.VolumeName = createWordpressServiceRequest.ContainerName
		fmt.Println("CotainerID: ", containerID)
	}

	_, err = dockerController.dockerService.FindVolume(createWordpressServiceRequest.VolumeName)

	if err != nil {
		log.Fatalln(err)
	}

	networkId, err := dockerController.dockerService.FindNetwork(createWordpressServiceRequest.NetworkName)

	if err != nil {
		panic(err)
	}

	go dockerController.dockerService.RunContainer(createWordpressServiceRequest, networkId)

	c.JSON(200, gin.H{
		"message": "pong",
	})
}
