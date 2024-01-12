package export

import (
	"encoding/json"
	"os"

	"github.com/jaum-cloud/processor/action"
)

func Export(value interface{}) {
	// Serializa os inputs em JSON
	inputJSON, _ := json.Marshal(action.StepData{Data: value})
	os.Setenv("WORKFLOW_PARAMS", string(inputJSON))
}
