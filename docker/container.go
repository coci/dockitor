package docker

import (
	"context"
	"fmt"
	"strings"

	"github.com/docker/docker/api/types"
)

type Container struct {
	Name   string
	Id     string
	Image  string
	Status string
}

// ContainerList : Get all containers
// it returns a list of Container structs
func ContainerList() []Container {
	var containerNames []Container

	cli := DockerClient()
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	if err != nil {
		fmt.Println(err)
	}
	for _, container := range containers {
		ctn := Container{Name: strings.Replace(container.Names[0], "/", "", -1), Id: container.ID[:10], Image: container.Image, Status: container.Status}
		containerNames = append(containerNames, ctn)
	}

	return containerNames
}