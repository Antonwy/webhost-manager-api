package listStacks

import (
	"fmt"
	"net/http"
	"whm-api/utils/db"
	"whm-api/utils/db/stacks"

	"github.com/docker/docker/client"
)

type Repository interface {
	ListStacksRepository() ([]stacks.Stack, string)
}

type repository struct {
	client *client.Client
}

func NewRepository(client *client.Client) *repository {
	return &repository{
		client: client,
	}
}

func (r *repository) ListStacksRepository() ([]stacks.Stack, string) {
	stacks := []stacks.Stack{}
	if err := db.DB.Select(&stacks, "select * from stacks;"); err != nil {
		fmt.Println(err)
		return nil, "Error: Cannot get stacks!"
	}

	return stacks, http.StatusText(http.StatusOK)
}
