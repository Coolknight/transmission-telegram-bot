package dockerhandler

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func RestartContainer(containerName string) error {
	ctx := context.Background()

	// Initialize Docker client
	dockerClient, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return fmt.Errorf("error creating Docker client: %v", err)
	}

	opts := container.StopOptions{Signal: "5"}
	// Stop the container
	if err := dockerClient.ContainerRestart(ctx, containerName, opts); err != nil {
		return fmt.Errorf("error restarting container: %v", err)
	}

	return nil
}
