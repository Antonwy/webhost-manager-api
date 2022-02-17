package stacks

import "whm-api/utils/docker/network"

func (stack Stack) StackStart() error {
	for _, c := range stack.Containers {
		if stack.NetworkID != "" {
			c.NetworkName = stack.NetworkName
		}

		if err := c.Validate(); err != nil {
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

	for _, c := range stack.Containers {
		c.NetworkID = stack.NetworkID
		if err := c.QuickStart(); err != nil {
			stack.Remove()
			return err
		}
	}

	return nil
}
