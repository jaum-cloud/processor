package params

import (
	"encoding/json"
	"os"
)

func NewParams(value interface{}) {
	// Serializa os inputs em JSON
	inputJSON, _ := json.Marshal(value)
	os.Setenv("WORKFLOW_PARAMS", string(inputJSON))
}

func Params(decode interface{}) {
	// Serializa os inputs em JSON
	value := os.Getenv("WORKFLOW_PARAMS")
	json.Unmarshal([]byte(value), decode)
}
