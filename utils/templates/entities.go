package templates

import "whm-api/utils/constants"

func DirectoryPath() string {
	return constants.BasePath() + "/templates"
}

type Template struct {
	Name      string             `json:"name"`
	ID        string             `json:"id"`
	Files     []TemplateFile     `json:"files"`
	Questions []TemplateQuestion `json:"questions"`
}

type TemplateResponse struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

type TemplateFile struct {
	Template string `json:"template"`
	File     string `json:"file"`
}

type TemplateQuestion struct {
	Title     string             `json:"title,omitempty"`
	Text      string             `json:"text,omitempty"`
	Type      string             `json:"type,omitempty"`
	Validator string             `json:"validator,omitempty"`
	Variable  string             `json:"variable,omitempty"`
	Required  bool               `json:"required,omitempty"`
	Questions []TemplateQuestion `json:"questions,omitempty"`
}

type KeyValue = map[string]string
type TemplateInput struct {
	Options KeyValue
	Name    string
	Url     string
	Envs    map[string]KeyValue
}
