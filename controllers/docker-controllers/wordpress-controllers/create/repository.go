package createWordPress

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

type Repository interface {
	CreateWordPressRepository(input *InputCreateWordPress) (*InputCreateWordPress, string)
}

type repository struct {
	client *client.Client
}

func NewRepositoryCreate(client *client.Client) *repository {
	return &repository{
		client: client,
	}
}

func (r *repository) CreateWordPressRepository(input *InputCreateWordPress) (*InputCreateWordPress, string) {

	ctx := context.Background()
	cli := r.client

	mariaDBImageName := "mariadb:latest"
	mdb_reader, err := cli.ImagePull(ctx, mariaDBImageName, types.ImagePullOptions{})
	if err != nil {
		fmt.Println(err)
		return nil, "Error pulling mariadb image"
	}

	defer mdb_reader.Close()
	io.Copy(os.Stdout, mdb_reader)

	dbContainerName := "wp_db"

	dbResp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: mariaDBImageName,
		Env: []string{
			fmt.Sprintf("MYSQL_ROOT_PASSWORD=%s", input.DBPassword),
			"MYSQL_DATABASE=wordpress",
			"MYSQL_USER=wordpress",
			"MYSQL_PASSWORD=wordpress",
		},
		Volumes: map[string]struct{}{
			"/wp_db": {},
		},
	}, nil, nil, nil, dbContainerName)
	if err != nil {
		fmt.Println(err)
		return nil, "Error creating mariadb"
	}

	if err := cli.ContainerStart(ctx, dbResp.ID, types.ContainerStartOptions{}); err != nil {
		fmt.Println(err)
		return nil, "Error starting mariadb container"
	}

	wp_reader, err := cli.ImagePull(ctx, "wordpress:latest", types.ImagePullOptions{})
	if err != nil {
		fmt.Println(err)
		return nil, "Error pulling wp image"
	}

	defer wp_reader.Close()
	io.Copy(os.Stdout, wp_reader)

	wpContainerName := "wp"

	wpResp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "wordpress:latest",
		Env: []string{
			"WORDPRESS_DB_USER=wordpress",
			"WORDPRESS_DB_PASSWORD=wordpress",
			"WORDPRESS_DB_NAME=wordpress"},
	}, &container.HostConfig{
		Links: []string{
			"wp_db:mysql",
		},
		PortBindings: nat.PortMap{
			"80/tcp": []nat.PortBinding{
				{
					HostIP:   "0.0.0.0",
					HostPort: "3002",
				},
			},
		},
	}, nil, nil, wpContainerName)
	if err != nil {
		fmt.Println(err)
		return nil, "Error creating wp"
	}

	if err := cli.ContainerStart(ctx, wpResp.ID, types.ContainerStartOptions{}); err != nil {
		fmt.Println(err)
		return nil, "Error starting wp container"
	}

	return nil, http.StatusText(http.StatusOK)
}
