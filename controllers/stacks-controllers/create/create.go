package createStackController

import (
	"context"
	"fmt"
	"log"
	"net/http"
	util "whm-api/utils"
	"whm-api/utils/db/stacks"
	"whm-api/utils/docker"
	"whm-api/utils/templates"
)

func (c controller) CreateStack(input CreateStackInput) (stacks.Stack, string) {

	ctx := context.Background()
	cli := c.client

	config := docker.Config{
		Context: ctx,
		Client:  cli,
	}

	tmplReq := templates.TemplateInput{
		Name:    util.LoweredAndUnderscored(input.Name),
		Options: make(map[string]string),
		Envs:    make(map[string]KeyValue),
		Url:     input.Domain,
	}

	fmt.Println(input)

	for _, opt := range input.Options {
		if opt.Type == "env" {
			if tmplReq.Envs[opt.Service] == nil {
				tmplReq.Envs[opt.Service] = make(map[string]string)
			}
			tmplReq.Envs[opt.Service][opt.Variable] = opt.Value
		} else if opt.Type == "option" {
			tmplReq.Options[opt.Variable] = opt.Value
		}
	}

	//virtualHostUrls := []string{stack.Url, "www." + stack.Url}
	//virtualHostsJoined := strings.Join(virtualHostUrls, ",")

	template, err := templates.Read(input.Template)
	if err != nil {
		log.Println(err)
		return stacks.Stack{}, "Couldn't read template with name: " + input.Name
	}

	project, err := template.Exec(tmplReq)
	if err != nil {
		log.Println(err)
		return stacks.Stack{}, "Couldn't convert template to project"
	}

	stack := stacks.Stack{
		ID:          stacks.GenerateStackID(),
		Name:        input.Name,
		Config:      config,
		NetworkName: docker.ProxyNetworkName,
		Type:        template.ID,
		Url:         input.Domain,
		Project:     project,
	}

	if err := stack.Create(); err != nil {
		fmt.Println(err.Error())
		stack.Remove()
		return stacks.Stack{}, "Couldn't create new container stack with name " + input.Name
	}

	//envs := []string{
	//	fmt.Sprintf("WORDPRESS_DB_USER=%s", input.DBUsername),
	//	fmt.Sprintf("WORDPRESS_DB_PASSWORD=%s", input.DBPassword),
	//	"WORDPRESS_DB_NAME=wordpress",
	//	fmt.Sprintf("WORDPRESS_DB_HOST=%s", wpDatabaseName),
	//	fmt.Sprintf("VIRTUAL_HOST=%s", virtualHostsJoined),
	//}

	//if input.SSLEmail != "" {
	//	envs = append(envs, fmt.Sprintf("LETSENCRYPT_HOST=%s", virtualHostsJoined))
	//	envs = append(envs, fmt.Sprintf("LETSENCRYPT_EMAIL=%s", input.SSLEmail))
	//}

	//phpConfigPath := fmt.Sprintf("%s/%s/uploads.ini", stacks.StacksDirectoryPath, stack.DirectoryName())
	//log.Println(phpConfigPath)
	//if _, err := os.Create(phpConfigPath); err != nil {
	//	stack.Remove()
	//	return stacks.Stack{}, "Couldn't create php-conf directory because " + err.Error()
	//}

	//wpPorts, _ := types.ParsePortConfig("80")
	//project := types.Project{
	//	Services: types.Services{
	//		{
	//			Name:          wpDatabaseName,
	//			Image:         mariadbImage,
	//			ContainerName: wpDatabaseName,
	//			Restart:       "always",
	//			Environment: types.NewMappingWithEquals([]string{
	//				fmt.Sprintf("MYSQL_ROOT_PASSWORD=%s", input.DBPassword),
	//				"MYSQL_DATABASE=wordpress",
	//				fmt.Sprintf("MYSQL_USER=%s", input.DBUsername),
	//				fmt.Sprintf("MYSQL_PASSWORD=%s", input.DBPassword),
	//			}),
	//			Volumes: []types.ServiceVolumeConfig{
	//				{
	//					Source: wpVolumeDatabaseName,
	//					Target: "/var/lib/mysql",
	//					Type:   "volume",
	//				},
	//			},
	//		},
	//		{
	//			Name:          wpContainerName,
	//			Image:         wordPressImage,
	//			ContainerName: wpContainerName,
	//			Restart:       "always",
	//			Environment:   types.NewMappingWithEquals(envs),
	//			Volumes: []types.ServiceVolumeConfig{
	//				{
	//					Source: wpVolumeName,
	//					Target: "/var/www/html",
	//					Type:   "volume",
	//				},
	//				// {
	//				// 	Source: ".",
	//				// 	Target: "/usr/local/etc/php/conf.d:ro",
	//				// 	Type:   "bind",
	//				// 	Bind:   &types.ServiceVolumeBind{CreateHostPath: true},
	//				// },
	//			},
	//			Ports: wpPorts,
	//		},
	//	},
	//	Volumes: types.Volumes{
	//		wpVolumeDatabaseName: types.VolumeConfig{},
	//		wpVolumeName:         types.VolumeConfig{},
	//	},
	//	Networks: types.Networks{
	//		"default": {
	//			Name: docker.ProxyNetworkName,
	//			External: types.External{
	//				External: true,
	//			},
	//		},
	//	},
	//}

	if err := stack.StackStart(); err != nil {
		stack.Remove()
		// os.RemoveAll(fmt.Sprintf("%s/%s", stacks.StacksDirectoryPath, stack.DirectoryName()))
		return stacks.Stack{}, "Couldn't create new stack because " + err.Error()
	}

	return stack, http.StatusText(http.StatusOK)
}
