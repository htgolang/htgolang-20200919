package main

import (
	"context"
	"fmt"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

func main() {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	body, err := cli.ContainerCreate(context.TODO(), &container.Config{
		Tty:       true,
		OpenStdin: true,
		Image:     "nginx:latest",
	}, &container.HostConfig{
		PortBindings: nat.PortMap{nat.Port("80/tcp"): []nat.PortBinding{{"0.0.0.0", "10002"}}},
	}, nil, nil, "testnginx3")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(body)

	containerID := body.ID

	err = cli.ContainerStart(context.TODO(), containerID, types.ContainerStartOptions{})
	fmt.Println(err)
	// containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	// if err != nil {
	// 	panic(err)
	// }

	// for _, container := range containers {
	// 	fmt.Printf("%s %s\n", container.ID[:10], container.Image)
	// }
}
