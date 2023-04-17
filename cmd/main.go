package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

func listContainers(ctx context.Context, cli *client.Client) {
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Println(container.ID)
	}
}

func runContainer(ctx context.Context, cli *client.Client, containerName string, containerImage string) {

	out, err := cli.ImagePull(ctx, containerImage, types.ImagePullOptions{All: false})
	if err != nil {
		panic(err)
	}
	defer out.Close()
	io.Copy(os.Stdout, out)

	//volume, err := cli.VolumeCreate(ctx, volume.CreateOptions{Name: containerName})

	if err != nil {
		panic(err)
	}

	config := &container.Config{
		Image: containerImage,
		Volumes: map[string]struct{}{
			//volume.Name + ":/var/www/html": {},
			"wordpress-web:/var/www/html": {},
		},
		ExposedPorts: nat.PortSet{
			"80/tcp": struct{}{},
		},
		Labels: map[string]string{
			"traefik.http.routers.wordpress.rule": "Host(`wordpress.docker.localhost`)",
		},
	}

	hostConfig := &container.HostConfig{
		PortBindings: nat.PortMap{
			"80/tcp": []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: "4140",
				},
			},
		},
		Resources: container.Resources{
			MemoryReservation: 512 * 1024 * 1024,
			CpusetCpus:        "0,5",
			CPUQuota:          10000,
		},
		Mounts: []mount.Mount{
			{
				Type:   mount.TypeVolume,
				Source: "wordpress-web",
				Target: "/var/www/html",
			},
		},
	}

	netConfig := network.EndpointSettings{
		NetworkID: "d37b9e8b6ff6aec262eea96d1a32a9893a94e8adee800c721c8735a20b6c9f4b",
	}

	networkConfig := network.NetworkingConfig{
		EndpointsConfig: map[string]*network.EndpointSettings{
			"traefik_default": &netConfig,
		},
	}

	resp, err := cli.ContainerCreate(ctx, config, hostConfig, &networkConfig, nil, containerName)
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	fmt.Println(resp.ID)
}

func removeContainer(ctx context.Context, cli *client.Client, containerID string) {
	err := cli.ContainerStop(ctx, containerID, container.StopOptions{})
	if err != nil {
		panic(err)
	}

	err = cli.ContainerRemove(ctx, containerID, types.ContainerRemoveOptions{RemoveVolumes: false, Force: true})
	if err != nil {
		panic(err)
	}
	fmt.Println("Container removed")

}

func findContainer(ctx context.Context, cli *client.Client, name string) (containerID string, err error) {
	args := filters.NewArgs(filters.KeyValuePair{
		Key:   "name",
		Value: name,
	})

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{
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

func main() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	containerID, err := findContainer(ctx, cli, "wordpress")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ContainerID: ", containerID)
		removeContainer(ctx, cli, containerID)
	}

	runContainer(ctx, cli, "wordpress-web", "wordpress")

}
