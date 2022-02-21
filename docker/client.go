package docker

import (
	"github.com/docker/docker/client"
)

// returns docker client
func dockerClient() *client.Client {
	cli, err := client.NewClientWithOpts(client.FromEnv)

	if err != nil {
		panic(err)
	}
	return cli
}
