package params

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jaum-cloud/processor/action"
)

func Params(decode interface{}) {
	// Serializa os inputs em JSON
	value := os.Getenv("WORKFLOW_PARAMS")
	json.Unmarshal([]byte(value), decode)
}

func Export(value interface{}) {
	// Serializa os inputs em JSON
	inputJSON, _ := json.Marshal(action.StepData{Data: value})
	os.Setenv("WORKFLOW_PARAMS", string(inputJSON))
	fmt.Println(string(inputJSON))
}
