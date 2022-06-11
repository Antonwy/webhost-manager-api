package templates

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func Read(name string) (Template, error) {
	e, err := os.Executable()

	fmt.Println(e)

	configJson, err := os.Open(fmt.Sprintf("%s/%s/config.json", DirectoryPath(), name))

	if err != nil {
		log.Println(err)
		return Template{}, err
	}

	defer configJson.Close()

	byteValue, _ := ioutil.ReadAll(configJson)

	var template Template
	json.Unmarshal(byteValue, &template)

	return template, nil
}
