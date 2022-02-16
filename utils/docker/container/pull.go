package dockerContainer

import (
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types"
)

func (container DockerContainer) PullImage() error {
	reader, err := container.Config.Client.ImagePull(container.Config.Context, container.Image, types.ImagePullOptions{})
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer reader.Close()
	io.Copy(os.Stdout, reader)

	return nil
}
