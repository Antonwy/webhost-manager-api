package createWordPress

import (
	"context"
	"fmt"
	"net/http"

	util "whm-api/utils"
	"whm-api/utils/db/stacks"
	"whm-api/utils/docker"
	dockerContainer "whm-api/utils/docker/container"

	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

const mariadbImage = "mariadb:latest"
const wordPressImage = "wordpress:latest"

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

	config := docker.Config{
		Context: ctx,
		Client:  cli,
	}

	wpContainerName := util.WordPressContainerName(input.Name)

	stack, err := stacks.Create(input.Name)

	if err != nil {
		return nil, "Couldn't create new container stack: " + err.Error()
	}

	mariadb := dockerContainer.DockerContainer{
		Config: config,
		Name:   wpContainerName + "_db",
		Image:  mariadbImage,
		Env: []string{
			fmt.Sprintf("MYSQL_ROOT_PASSWORD=%s", input.DBPassword),
			"MYSQL_DATABASE=wordpress",
			fmt.Sprintf("MYSQL_USER=%s", input.DBUsername),
			fmt.Sprintf("MYSQL_PASSWORD=%s", input.DBPassword),
		},
		Volumes: map[string]struct{}{
			"/" + wpContainerName + "_db": {},
		},
		NetworkName: wpContainerName + "_network",
		StackID:     stack.ID,
	}

	wp := dockerContainer.DockerContainer{
		Config: config,
		Name:   wpContainerName,
		Image:  wordPressImage,
		Env: []string{
			fmt.Sprintf("WORDPRESS_DB_USER=%s", input.DBUsername),
			fmt.Sprintf("WORDPRESS_DB_PASSWORD=%s", input.DBPassword),
			"WORDPRESS_DB_NAME=wordpress",
			fmt.Sprintf("WORDPRESS_DB_HOST=%s", mariadb.Name),
		},
		NetworkID: mariadb.NetworkID,
		PortBindings: map[nat.Port]nat.Port{
			nat.Port(input.Port): "80/tcp",
		},
		StackID: stack.ID,
	}

	stack.Containers = []dockerContainer.DockerContainer{
		mariadb, wp,
	}

	stack.StackStart()

	return nil, http.StatusText(http.StatusOK)
}
