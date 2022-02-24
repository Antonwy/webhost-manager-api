package createWordPress

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	util "whm-api/utils"

	"whm-api/utils/db/stacks"
	"whm-api/utils/docker"

	"github.com/compose-spec/compose-go/types"
	"github.com/docker/docker/client"
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
	wpVolumeDatabaseName := wpContainerName + "_db_volume"
	wpVolumeName := wpContainerName + "_volume"
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

	virtualHostUrls := []string{stack.Url, "www." + stack.Url}
	virtualHostsJoined := strings.Join(virtualHostUrls, ",")

	envs := []string{
		fmt.Sprintf("WORDPRESS_DB_USER=%s", input.DBUsername),
		fmt.Sprintf("WORDPRESS_DB_PASSWORD=%s", input.DBPassword),
		"WORDPRESS_DB_NAME=wordpress",
		fmt.Sprintf("WORDPRESS_DB_HOST=%s", wpDatabaseName),
		fmt.Sprintf("VIRTUAL_HOST=%s", virtualHostsJoined),
	}

	if input.SSLEmail != "" {
		envs = append(envs, fmt.Sprintf("LETSENCRYPT_HOST=%s", virtualHostsJoined))
		envs = append(envs, fmt.Sprintf("LETSENCRYPT_EMAIL=%s", input.SSLEmail))
	}

	phpConfigPath := fmt.Sprintf("%s/%s/uploads.ini", stacks.StacksDirectoryPath, stack.DirectoryName())
	log.Println(phpConfigPath)
	if _, err := os.Create(phpConfigPath); err != nil {
		stack.Remove()
		return stacks.Stack{}, "Couldn't create php-conf directory because " + err.Error()
	}

	wpPorts, _ := types.ParsePortConfig("80")
	project := types.Project{
		Services: types.Services{
			{
				Name:          wpDatabaseName,
				Image:         mariadbImage,
				ContainerName: wpDatabaseName,
				Restart:       "always",
				Environment: types.NewMappingWithEquals([]string{
					fmt.Sprintf("MYSQL_ROOT_PASSWORD=%s", input.DBPassword),
					"MYSQL_DATABASE=wordpress",
					fmt.Sprintf("MYSQL_USER=%s", input.DBUsername),
					fmt.Sprintf("MYSQL_PASSWORD=%s", input.DBPassword),
				}),
				Volumes: []types.ServiceVolumeConfig{
					{
						Source: wpVolumeDatabaseName,
						Target: "/var/lib/mysql",
						Type:   "volume",
					},
				},
			},
			{
				Name:          wpContainerName,
				Image:         wordPressImage,
				ContainerName: wpContainerName,
				Restart:       "always",
				Environment:   types.NewMappingWithEquals(envs),
				Volumes: []types.ServiceVolumeConfig{
					{
						Source: wpVolumeName,
						Target: "/var/www/html",
						Type:   "volume",
					},
					{
						Source: ".",
						Target: "/usr/local/etc/php/conf.d:ro",
						Type:   "bind",
					},
				},
				Ports: wpPorts,
			},
		},
		Volumes: types.Volumes{
			wpVolumeDatabaseName: types.VolumeConfig{},
			wpVolumeName:         types.VolumeConfig{},
		},
		Networks: types.Networks{
			"default": {
				Name: docker.ProxyNetworkName,
				External: types.External{
					External: true,
				},
			},
		},
	}

	stack.Project = project

	if err := stack.StackStart(); err != nil {
		stack.Remove()
		// os.RemoveAll(fmt.Sprintf("%s/%s", stacks.StacksDirectoryPath, stack.DirectoryName()))
		return stacks.Stack{}, "Couldn't create new stack because " + err.Error()
	}

	return stack, http.StatusText(http.StatusOK)
}
