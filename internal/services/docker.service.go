package service

import (
	"api-hosting/internal/models/requests"
	"context"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types/volume"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type DockerService struct {
	ctx context.Context
	cli *client.Client
}

type IDockerService interface {
	ListContainers()
	FindContainer(name string) (containerID string, err error)
	DescribeContainer(name string) (container types.Container, err error)
	FindNetwork(networkName string) (networkId string, err error)
	FindVolume(volumeName string) (volumeId string, err error)
	RunContainer(createWordpressServiceRequest requests.CreateWordpressServiceRequest, networkId string)
	RemoveContainer(containerID string)
	RemoveVolume(volumeName string)
	CreateDB(containerName string)
	DeleteDB(containerName string)
}

func NewDockerService(ctx context.Context, cli *client.Client) IDockerService {
	return &DockerService{
		ctx: ctx,
		cli: cli,
	}
}

func (dockerService *DockerService) ListContainers() {
	containers, err := dockerService.cli.ContainerList(dockerService.ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Println(container.ID)
	}
}

func (dockerService *DockerService) FindNetwork(networkName string) (networkId string, err error) {
	args := filters.NewArgs(filters.KeyValuePair{
		Key:   "name",
		Value: networkName,
	})

	networks, err := dockerService.cli.NetworkList(dockerService.ctx, types.NetworkListOptions{
		Filters: args,
	})

	if err != nil {
		return "", err
	}

	for _, network := range networks {
		fmt.Println("Network ID: ", network.ID)
		fmt.Println("Network Name: ", network.Name)
	}

	return networks[0].ID, nil
}

func (dockerService *DockerService) FindVolume(volumeName string) (volumeId string, err error) {
	args := filters.NewArgs(filters.KeyValuePair{
		Key:   "name",
		Value: volumeName,
	})

	volumes, err := dockerService.cli.VolumeList(dockerService.ctx, args)

	if err != nil {
		return "", err
	}

	for _, volume := range volumes.Volumes {
		fmt.Println("Volume Name: ", volume.Name)
		fmt.Println("Volume Usage Data: ", volume.UsageData)
	}

	if len(volumes.Volumes) == 0 {
		return "", nil
	}

	return volumes.Volumes[0].Name, nil
}

func (dockerService *DockerService) RunContainer(createWordpressServiceRequest requests.CreateWordpressServiceRequest, networkId string) {

	fmt.Println("Start running wpInstance Container")
	out, err := dockerService.cli.ImagePull(dockerService.ctx, createWordpressServiceRequest.ContainerImage, types.ImagePullOptions{All: false})
	if err != nil {
		panic(err)
	}
	defer out.Close()
	io.Copy(os.Stdout, out)

	if createWordpressServiceRequest.VolumeName == "" {
		volume, err := dockerService.cli.VolumeCreate(dockerService.ctx, volume.CreateOptions{Name: createWordpressServiceRequest.ContainerName})
		if err != nil {
			panic(err)
		}
		createWordpressServiceRequest.VolumeName = volume.Name
	}

	config := &container.Config{
		Image: createWordpressServiceRequest.ContainerImage,
		ExposedPorts: nat.PortSet{
			"80/tcp": struct{}{},
		},
		Labels: map[string]string{
			"traefik.http.routers." + createWordpressServiceRequest.ContainerName + ".entrypoints": "appwrite_web",
			"traefik.http.routers." + createWordpressServiceRequest.ContainerName + ".service":     createWordpressServiceRequest.ContainerName,
			"traefik.http.routers." + createWordpressServiceRequest.ContainerName + ".rule":        "Host(`" + createWordpressServiceRequest.ContainerName + "`)",
			"traefik.enable":                 "true",
			"traefik.docker.network":         "appwrite",
			"traefik.constraint-label-stack": "appwrite",
			"traefik.http.services." + createWordpressServiceRequest.ContainerName + ".loadbalancer.server.port": "80",
		},
		Env: []string{
			"WORDPRESS_DB_NAME=" + createWordpressServiceRequest.ContainerName,
			"WORDPRESS_DB_USER=root",
			"WORDPRESS_DB_PASSWORD=DontTouchMyDbServer2021*",
			"WORDPRESS_DB_HOST=172.26.0.23",
		},
	}

	hostConfig := &container.HostConfig{
		/*PortBindings: nat.PortMap{
			"80/tcp": []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: "0",
				},
			},
		},*/
		Resources: container.Resources{
			MemoryReservation: 512 * 1024 * 1024,
			CPUQuota:          10000,
		},
		Mounts: []mount.Mount{
			{
				Type:   mount.TypeVolume,
				Source: createWordpressServiceRequest.VolumeName,
				Target: "/var/www/html",
			},
		},
	}

	netConfig := network.EndpointSettings{
		NetworkID: networkId,
	}

	networkConfig := network.NetworkingConfig{
		EndpointsConfig: map[string]*network.EndpointSettings{
			"appwriteconfig_appwrite": &netConfig,
		},
	}

	resp, err := dockerService.cli.ContainerCreate(dockerService.ctx, config, hostConfig, &networkConfig, nil, createWordpressServiceRequest.ContainerName)
	if err != nil {
		panic(err)
	}

	if err := dockerService.cli.ContainerStart(dockerService.ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

}

func (dockerService *DockerService) RemoveContainer(containerID string) {
	err := dockerService.cli.ContainerStop(dockerService.ctx, containerID, container.StopOptions{})
	if err != nil {
		panic(err)
	}

	err = dockerService.cli.ContainerRemove(dockerService.ctx, containerID, types.ContainerRemoveOptions{RemoveVolumes: false, Force: true})
	if err != nil {
		panic(err)
	}
	fmt.Println("Container removed")
}

func (dockerService *DockerService) RemoveVolume(volumeName string) {

	volumeId, err := dockerService.FindVolume(volumeName)
	if err != nil {
		panic(err)
	}

	err = dockerService.cli.VolumeRemove(dockerService.ctx, volumeId, true)
	if err != nil {
		panic(err)
	}

	fmt.Println("Volume removed")
}

func (dockerService *DockerService) FindContainer(name string) (containerID string, err error) {
	args := filters.NewArgs(filters.KeyValuePair{
		Key:   "name",
		Value: name,
	})

	containers, err := dockerService.cli.ContainerList(dockerService.ctx, types.ContainerListOptions{
		Filters: args,
	})

	if err != nil {
		fmt.Println(err)
	}

	for _, container := range containers {
		fmt.Println(container.ID)
	}

	if len(containers) > 0 {
		return containers[0].ID, nil
	}

	return "", errors.New("Container not found")
}

func (dockerService *DockerService) DescribeContainer(name string) (container types.Container, err error) {
	args := filters.NewArgs(filters.KeyValuePair{
		Key:   "name",
		Value: name,
	})

	containers, err := dockerService.cli.ContainerList(dockerService.ctx, types.ContainerListOptions{
		Filters: args,
	})

	if err != nil {
		fmt.Println(err)
	}

	for _, container := range containers {
		fmt.Println(container.ID)
	}

	if len(containers) > 0 {
		return containers[0], nil
	}

	return types.Container{}, errors.New("Container not found")
}

func (dockerService *DockerService) CreateDB(containerName string) {

	fmt.Println("Starting to create database " + containerName)

	db, err := sql.Open("mysql", "root:DontTouchMyDbServer2021*@tcp(172.29.0.24:3306)/dbtest")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	res, err := db.Exec("CREATE DATABASE " + containerName)

	fmt.Println(res)
	fmt.Println(err)

	fmt.Println("Database " + containerName + " created")
}

func (dockerService *DockerService) DeleteDB(containerName string) {

	fmt.Println("Starting to delete database " + containerName)

	db, err := sql.Open("mysql", "root:DontTouchMyDbServer2021*@tcp(172.29.0.24:3306)/dbtest")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	res, err := db.Exec("DROP DATABASE " + containerName)

	fmt.Println(res)
	fmt.Println(err)

	fmt.Println("Database " + containerName + " deleted")
}
