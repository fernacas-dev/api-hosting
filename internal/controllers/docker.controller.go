package controllers

import (
	"api-hosting/internal/models/requests"
	service "api-hosting/internal/services"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type DockerController struct {
	dockerService service.IDockerService
}

type IDockerController interface {
	CreateWordpressService(c *gin.Context)
	GetWordpressService(c *gin.Context)
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
		"message": createWordpressServiceRequest.ContainerName,
	})
}

func (dockerController *DockerController) GetWordpressService(c *gin.Context) {
	name, _ := c.Params.Get("name")
	containerInfo, _ := dockerController.dockerService.DescribeContainer(name)
	c.JSON(200, gin.H{
		"message": containerInfo,
	})
}
