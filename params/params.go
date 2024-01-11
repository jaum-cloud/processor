package params

import (
	"encoding/json"
	"os"
)

var paramsMap map[string]interface{}

func init() {
	// Carrega os parâmetros a partir de um arquivo JSON
	paramsFile := os.Getenv("PARAMS_FILE")
	data, err := os.ReadFile(paramsFile)
	if err != nil {
		panic(err)
	}

	paramsMap = make(map[string]interface{})
	if err := json.Unmarshal(data, &paramsMap); err != nil {
		panic(err)
	}
}

// Parse preenche uma struct com os parâmetros.
func Parse(v interface{}) error {
	data, err := json.Marshal(paramsMap)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}


// Trigger
// Run first action
// Safe output action inside json



