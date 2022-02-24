package removeStack

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"whm-api/utils/cli"
	"whm-api/utils/db"
	"whm-api/utils/db/stacks"

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
	stack, err := stacks.GetFromID(id)
	if err != nil {
		fmt.Println(err)
		return "Couldn't get stack with ID: " + id
	}

	// if id, err := network.IdFromName(stack.NetworkName, config); err == nil {
	// 	if err := network.Remove(id, config); err != nil {
	// 		return "Couldn't remove network " + stack.NetworkName
	// 	}
	// } else {
	// 	return err.Error()
	// }

	dirName := stack.DirectoryName()

	cmd := exec.Command("docker-compose", "down", "-v", "--remove-orphans")
	cmd.Dir = fmt.Sprintf("%s/%s", stacks.StacksDirectoryPath, dirName)

	downErr := cli.Run(cmd)

	if downErr != nil {
		return "Couldn't remove compose file!"
	}

	errRemoveDir := os.RemoveAll(fmt.Sprintf("%s/%s", stacks.StacksDirectoryPath, dirName))

	if errRemoveDir != nil {
		return "Couldn't remove stack directory!"
	}

	if _, err := db.DB.Exec("delete from stacks where id = $1;", id); err != nil {
		fmt.Println(err)
		return "Error: Couldn't remove stack!"
	}

	return http.StatusText(http.StatusOK)
}
