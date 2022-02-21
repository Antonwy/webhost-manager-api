package createWordPress

import (
	"context"
	"fmt"
	"net/http"
	"strings"

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
	CreateWordPressRepository(input *InputCreateWordPress) (stacks.Stack, string)
}

type repository struct {
	client *client.Client
}

func NewRepositoryCreate(client *client.Client) *repository {
	return &repository{
		client: client,
	}
}

func (r *repository) CreateWordPressRepository(input *InputCreateWordPress) (stacks.Stack, string) {

	ctx := context.Background()
	cli := r.client

	config := docker.Config{
		Context: ctx,
		Client:  cli,
	}

	wpContainerName := util.WordPressContainerName(input.Name)
	wpNetworkName := wpContainerName + "_network"
	wpDatabaseName := wpContainerName + "_db"
	stackType := "WordPress"

	stack := stacks.Stack{
		ID:          stacks.GenerateStackID(),
		Name:        input.Name,
		Config:      config,
		NetworkName: wpNetworkName,
		Type:        stackType,
		Url:         input.Url,
	}

	if err := stack.Create(); err != nil {
		fmt.Println(err.Error())
		stack.Remove()
		return stacks.Stack{}, "Couldn't create new container stack with name " + input.Name
	}

	mariadb := dockerContainer.DockerContainer{
		Config: config,
		Name:   wpDatabaseName,
		Image:  mariadbImage,
		Env: []string{
			fmt.Sprintf("MYSQL_ROOT_PASSWORD=%s", input.DBPassword),
			"MYSQL_DATABASE=wordpress",
			fmt.Sprintf("MYSQL_USER=%s", input.DBUsername),
			fmt.Sprintf("MYSQL_PASSWORD=%s", input.DBPassword),
		},
		Volumes: map[string]struct{}{
			"/" + wpDatabaseName: {},
		},
		StackID: stack.ID,
	}

	virtualHostUrls := []string{stack.Url, "www." + stack.Url}
	virtualHostsJoined := strings.Join(virtualHostUrls, ",")

	envs := []string{
		fmt.Sprintf("WORDPRESS_DB_USER=%s", input.DBUsername),
		fmt.Sprintf("WORDPRESS_DB_PASSWORD=%s", input.DBPassword),
		"WORDPRESS_DB_NAME=wordpress",
		fmt.Sprintf("WORDPRESS_DB_HOST=%s", mariadb.Name),
		fmt.Sprintf("VIRTUAL_HOST=%s", virtualHostsJoined),
	}

	if input.SSLEmail != "" {
		envs = append(envs, fmt.Sprintf("LETSENCRYPT_HOST=%s", virtualHostsJoined))
		envs = append(envs, fmt.Sprintf("LETSENCRYPT_EMAIL=%s", input.SSLEmail))
	}

	wp := dockerContainer.DockerContainer{
		Config:    config,
		Name:      wpContainerName,
		Image:     wordPressImage,
		Env:       envs,
		NetworkID: mariadb.NetworkID,
		PortBindings: map[nat.Port]nat.Port{
			"0": "80/tcp",
		},
		StackID: stack.ID,
	}

	stack.Containers = []dockerContainer.DockerContainer{
		mariadb, wp,
	}

	if err := stack.StackStart(); err != nil {
		stack.Remove()
		return stacks.Stack{}, "Couldn't create new container stack because " + err.Error()
	}

	return stack, http.StatusText(http.StatusOK)
}
