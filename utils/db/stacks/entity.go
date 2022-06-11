package stacks

import (
	util "whm-api/utils"
	"whm-api/utils/constants"
	"whm-api/utils/docker"
	dockerContainer "whm-api/utils/docker/container"

	"github.com/compose-spec/compose-go/types"
)

type Stack struct {
	ID          string                            `db:"id" json:"id"`
	Name        string                            `db:"name" json:"name"`
	Config      docker.Config                     `json:"config"`
	Containers  []dockerContainer.DockerContainer `json:"-"`
	NetworkName string                            `db:"network_name"  json:"network_name"`
	NetworkID   string                            `json:"-"`
	Type        string                            `db:"type"  json:"type"`
	Url         string                            `db:"url"  json:"url"`
	Project     types.Project                     `json:"-"`
}

func (stack Stack) Response() ResponseStack {
	return ResponseStack{
		ID:          stack.ID,
		Name:        stack.Name,
		NetworkName: stack.NetworkName,
		Type:        stack.Type,
		Url:         stack.Url,
	}
}

func DirectoryPath() string {
	return constants.BasePath() + "/stacks"
}

func (stack Stack) DirectoryName() string {
	if stack.Url != "" {
		return stack.Url
	}

	fileName := stack.Name

	return util.LoweredAndUnderscored(fileName)
}

type ResponseStack struct {
	ID          string `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	NetworkName string `db:"network_name"  json:"network_name"`
	Type        string `db:"type"  json:"type"`
	Url         string `db:"url"  json:"url"`
}
