package templates

import (
	"errors"
	"fmt"
	"github.com/compose-spec/compose-go/cli"
	"github.com/compose-spec/compose-go/types"
	"html/template"
	"log"
	"os"
	"whm-api/utils/db/stacks"
)

func (tmpl Template) Exec(input TemplateInput) (types.Project, error) {
	stackDirectory := fmt.Sprintf("%s/%s", stacks.DirectoryPath(), input.Url)
	templateDirectory := fmt.Sprintf("%s/%s", DirectoryPath(), tmpl.ID)
	err := os.MkdirAll(stackDirectory, os.ModePerm)
	if err != nil {
		return types.Project{}, errors.New("couldn't create stacks directory")
	}

	for _, file := range tmpl.Files {
		f, err := os.Create(fmt.Sprintf("%s/%s", stackDirectory, file.File))

		if err != nil {
			log.Println(err)
			return types.Project{}, errors.New("Couldn't create file: " + file.File)
		}

		temp, err := template.ParseFiles(fmt.Sprintf("%s/%s", templateDirectory, file.Template))

		if err != nil {
			log.Println(err)
			return types.Project{}, errors.New("couldn't parse template")
		}

		execErr := temp.Execute(f, input)

		f.Close()

		if execErr != nil {
			log.Println(execErr)
			return types.Project{}, errors.New("couldn't execute template")
		}
	}

	dockerCompose, err := os.Open(fmt.Sprintf("%s/docker-compose.yaml", stackDirectory))
	if err != nil {
		log.Println(err)
		return types.Project{}, errors.New("couldn't open docker-compose")
	}

	defer dockerCompose.Close()

	options := cli.ProjectOptions{
		WorkingDir:  stackDirectory,
		Environment: map[string]string{},
	}

	err = cli.WithDefaultConfigPath(&options)
	if err != nil {
		log.Println(err)
		return types.Project{}, errors.New("couldn't find docker-compose files in the working directory")
	}

	project, err := cli.ProjectFromOptions(&options)

	if err != nil {
		log.Println(err)
		return types.Project{}, errors.New("couldn't parse docker-compose from options")
	}

	return *project, nil
}
