package stacks

import (
	"strings"
	"whm-api/utils/docker"
	dockerContainer "whm-api/utils/docker/container"

	"github.com/compose-spec/compose-go/types"
)

type Stack struct {
	ID          string                            `db:"id" json:"id"`
	Name        string                            `db:"name" json:"name"`
	Config      docker.Config                     `json:"config"`
	Containers  []dockerContainer.DockerContainer `json:"containers"`
	NetworkName string                            `db:"network_name"  json:"network_name"`
	NetworkID   string                            `json:"network_id"`
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

const StacksDirectoryPath = "/data/stacks"

func (stack Stack) DirectoryName() string {
	if stack.Url != "" {
		return stack.Url
	}

	fileName := stack.Name
	fileName = strings.ReplaceAll(fileName, " ", "_")
	fileName = strings.ToLower(fileName)

	return fileName
}

type ResponseStack struct {
	ID          string `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	NetworkName string `db:"network_name"  json:"network_name"`
	Type        string `db:"type"  json:"type"`
	Url         string `db:"url"  json:"url"`
}
