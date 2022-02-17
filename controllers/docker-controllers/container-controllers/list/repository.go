package listContainers

import (
	"context"
	"net/http"
	"whm-api/utils/docker"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type Repository interface {
	ListContainersRepository() ([]types.Container, string)
}

type repository struct {
	client *client.Client
}

func NewRepositoryCreate(client *client.Client) *repository {
	return &repository{
		client: client,
	}
}

func (r *repository) ListContainersRepository() ([]types.Container, string) {

	ctx := context.Background()

	config := docker.Config{
		Context: ctx,
		Client:  r.client,
	}

	containers, err := config.ListContainers()

	if err != nil {
		return []types.Container{}, err.Error()
	}

	return containers, http.StatusText(http.StatusOK)
}
