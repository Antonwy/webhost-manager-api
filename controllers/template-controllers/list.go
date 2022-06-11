package templateControllers

import (
	"io/ioutil"
	"log"
	"net/http"
	"whm-api/utils/templates"
)

func (c controller) ListTemplates() ([]templates.TemplateResponse, string) {
	files, err := ioutil.ReadDir(templates.Directory)

	if err != nil {
		log.Println(err)
		return nil, "Couldn't find templates directory!"
	}

	var tmplts []templates.TemplateResponse

	for _, file := range files {
		template, err := templates.Read(file.Name())

		if err != nil {
			log.Println(err)
			continue
		}

		tmplts = append(tmplts, templates.TemplateResponse{
			Name: template.Name,
			ID:   template.ID,
		})
	}

	return tmplts, http.StatusText(http.StatusOK)
}
