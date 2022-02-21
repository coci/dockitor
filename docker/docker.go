package dumitor

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)


func dockerClient() *client.Client {
    cli, err := client.NewClientWithOpts(client.FromEnv)

	if err != nil {
		panic(err)
    }
    return cli
}

func containerList() []string {
    cli := dockerClient()
    containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
    if err != nil {
        panic(err)
    }
    var containerNames []string
    for _, container := range containers {
        fmt.Println(containerNames, container.Names)
    }
    return containerNames
}


