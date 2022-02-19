package stacks

import (
	"whm-api/utils/docker/network"
)

func (stack *Stack) StackStart() error {
	for i := range stack.Containers {
		if stack.NetworkID != "" {
			stack.Containers[i].NetworkName = stack.NetworkName
		}

		if stack.Url != "" {
			stack.Containers[i].ConnectToProxyNetwork = true
		}

		if err := stack.Containers[i].Validate(); err != nil {
			return err
		}
	}

	if networkName := stack.NetworkName; networkName != "" {
		id, err := network.Create(networkName, stack.Config)

		if err != nil {
			return err
		}

		stack.NetworkID = id
	}

	for i := range stack.Containers {
		stack.Containers[i].NetworkID = stack.NetworkID
		if err := stack.Containers[i].QuickStart(); err != nil {
			stack.Remove()
			return err
		}

	}

	return nil
}
