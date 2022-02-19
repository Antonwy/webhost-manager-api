package docker

import (
	"context"

	"github.com/docker/docker/client"
)

type Config struct {
	Context context.Context
	Client  *client.Client
}

const ProxyNetworkName = "reverse-proxy-net"
