package stacks

import (
	"fmt"
	"os"
	"os/exec"

	"whm-api/utils/cli"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

func (stack Stack) StackStart() error {

	res, err := yaml.Marshal(stack.Project)

	if err != nil {
		return errors.Wrap(err, "Couldn't create YAML!")
	}

	dirName := stack.DirectoryName()

	errCreateFile := os.WriteFile(fmt.Sprintf("%s/%s/docker-compose.yaml", StacksDirectoryPath, dirName), res, 0644)

	if errCreateFile != nil {
		return errors.Wrap(err, "Couldn't create compose file!")
	}

	cmd := exec.Command("docker-compose", "up", "-d")
	cmd.Dir = fmt.Sprintf("%s/%s", StacksDirectoryPath, dirName)

	if err := cli.Run(cmd); err != nil {
		return errors.Wrap(err, "Couldn't start compose file!")
	}

	return nil
}
