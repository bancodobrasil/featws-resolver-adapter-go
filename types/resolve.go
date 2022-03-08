package types

type ResolveInput struct {
	Context map[string]interface{} `json:"context"`
	Load    []string               `json:"load"`
}

// ResolveOutput contais all output of resolver execution
type ResolveOutput struct {
	Context map[string]interface{} `json:"context"`
	Errors  map[string]interface{} `json:"errors"`
}
