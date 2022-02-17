package removeStack

import (
	"context"
	"fmt"
	"net/http"
	"whm-api/utils/db"
	db_containers "whm-api/utils/db/containers"
	"whm-api/utils/db/stacks"
	"whm-api/utils/docker"
	"whm-api/utils/docker/network"

	"github.com/docker/docker/client"
)

type Repository interface {
	RemoveStackRepository(id string) string
}

type repository struct {
	client *client.Client
}

func NewRepository(client *client.Client) *repository {
	return &repository{
		client: client,
	}
}

func (r *repository) RemoveStackRepository(id string) string {

	ctx := context.Background()
	config := docker.Config{
		Context: ctx, Client: r.client,
	}

	stack, err := stacks.GetFromID(id)
	if err != nil {
		fmt.Println(err)
		return "Couldn't get stack with ID: " + id
	}

	containers, err := db_containers.ListContainersFromStack(id)

	if err != nil {
		fmt.Println(err)
		return "Error: Could't find any containers from stack!"
	}

	for _, container := range containers {
		if err := container.RemoveWithContainer(config); err != nil {
			fmt.Println(err)
			return "Couldn't remove container " + container.ID
		}
	}

	if id, err := network.IdFromName(stack.NetworkName, config); err == nil {
		if err := network.Remove(id, config); err != nil {
			return "Couldn't remove network " + stack.NetworkName
		}
	} else {
		return err.Error()
	}

	if _, err := db.DB.Exec("delete from stacks where id = $1;", id); err != nil {
		fmt.Println(err)
		return "Error: Couldn't remove stack!"
	}

	return http.StatusText(http.StatusOK)
}
