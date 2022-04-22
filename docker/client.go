package docker

import (
	"fmt"
	"github.com/docker/docker/client"
)

// returns docker client
func DockerClient() *client.Client {
	cli, err := client.NewClientWithOpts(client.FromEnv)

	if err != nil {
		fmt.Println(err)
	}
	return cli
}
