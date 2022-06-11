package createStackController

import (
	"github.com/docker/docker/client"
	"whm-api/utils/db/stacks"
)

type Controller interface {
	CreateStack(input CreateStackInput) (stacks.Stack, string)
}

type controller struct {
	client *client.Client
}

func NewController() *controller {
	return &controller{}
}
