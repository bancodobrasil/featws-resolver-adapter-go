package types

// ResolveInput contains all input for resolver execution
//
// Property:
//   - Context: is a map of key-value pairs that provides additional information or context for the input. It can be used to pass any relevant data that may be required for processing the input. The keys in the map are strings, and the values can be of any type.
//   - {[]string} Load: is a slice of strings that specifies the names of the modules or packages that need to be loaded before executing the code. This is useful when the code being executed depends on external libraries or modules. The `Load` property allows the code to import and use those dependencies.
type ResolveInput struct {
	Context map[string]interface{} `json:"context"`
	Load    []string               `json:"load"`
}

// ResolveOutput contais all output of resolver execution
//
// Property:
//   - Context: is a map that stores key-value pairs of information related to the resolution of a task or problem. It can contain any type of data, such as strings, numbers, arrays, or even nested objects. This property is often used to pass information between different parts of a program.The keys in the map represent the names or identifiers of the data, while the values can be of any type.
//   - Errors: is a map that contains any errors that occurred during the resolution process. The keys of the map are strings that identify the specific error, and the values are any additional information about the error.
type ResolveOutput struct {
	Context map[string]interface{} `json:"context"`
	Errors  map[string]interface{} `json:"errors"`
}
