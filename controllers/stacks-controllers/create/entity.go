package createStackController

type CreateStackInputOption struct {
	Type     string `json:"type"`
	Variable string `json:"variable"`
	Value    string `json:"value"`
	Service  string `json:"service"`
}

type CreateStackInputVolumes struct {
	Type   string `json:"type"`
	Source string `json:"source"`
	Target string `json:"target"`
}

type CreateStackInput struct {
	Domain   string
	Options  []CreateStackInputOption  `json:"options"`
	Name     string                    `json:"name"`
	Template string                    `json:"template"`
	Volumes  []CreateStackInputVolumes `json:"volumes"`
}

type KeyValue = map[string]string
type StackConfig struct {
	Options KeyValue
	Name    string
	Envs    map[string]KeyValue
}
