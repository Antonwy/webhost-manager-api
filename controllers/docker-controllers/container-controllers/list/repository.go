package listContainers

import (
	"context"
	"net/http"

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

	containers, err := r.client.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		return []types.Container{}, err.Error()
	}

	return containers, http.StatusText(http.StatusOK)
}
