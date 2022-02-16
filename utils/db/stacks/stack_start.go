package stacks

func (stack Stack) StackStart() error {

	for _, c := range stack.Containers {
		if err := c.Validate(); err != nil {
			return err
		}

		if err := c.QuickStart(); err != nil {
			stack.Remove()
			return err
		}
	}

	return nil
}
