package stacks

import (
	"fmt"
	"os/exec"

	"whm-api/utils/cli"

	"github.com/pkg/errors"
)

func (stack Stack) StackStart() error {
	dirName := stack.DirectoryName()
	stacksDirectory := DirectoryPath()

	cmd := exec.Command("docker-compose", "up", "-d")
	cmd.Dir = fmt.Sprintf("%s/%s", stacksDirectory, dirName)

	if err := cli.Run(cmd); err != nil {
		return errors.Wrap(err, "Couldn't start compose file!")
	}

	return nil
}
