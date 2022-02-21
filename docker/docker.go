package docker

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"strings"
)

type Container struct {
	Name   string
	Id     string
	Image  string
	Status string
}

// returns docker client
func dockerClient() *client.Client {
	cli, err := client.NewClientWithOpts(client.FromEnv)

	if err != nil {
		panic(err)
	}
	return cli
}

// ContainerList : Get all containers
// it returns a list of Container structs
func ContainerList() []Container {
	var containerNames []Container

	cli := dockerClient()
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		ctn := Container{Name: strings.Replace(container.Names[0], "/", "", -1), Id: container.ID[:10], Image: container.Image, Status: container.Status}
		containerNames = append(containerNames, ctn)
	}

	return containerNames
}
