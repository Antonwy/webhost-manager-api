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
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return []types.Container{}, err.Error()
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		return []types.Container{}, err.Error()
	}

	return containers, http.StatusText(http.StatusOK)
}
