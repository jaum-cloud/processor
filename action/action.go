package action

import (
	"encoding/json"
	"fmt"
	"github_automation/executor"
	"github_automation/github"
	"log"
	"os"
)

type Action interface {
	Execute(input interface{}) (output interface{}, err error)
	SetOrder(order int)
}

type StepData struct {
	Data  interface{}
	Error error
}

type BaseAction struct {
	Order  int
	ID     string
	Name   string
	Input  StepData
	Output StepData
}

func NewBaseAction(id string) *BaseAction {
	return &BaseAction{
		ID:   id,
		Name: id,
	}
}

func (ba *BaseAction) Execute(input interface{}) (output interface{}, err error) {
	ba.Input.Data = input

	code, err := github.FetchGoCodeFromGist(ba.ID)
	if err != nil {
		ba.Output.Error = err
		return nil, err
	}

	// Serializa os inputs em JSON
	inputJSON, err := json.Marshal(ba.Input)
	if err != nil {
		log.Printf("Erro ao serializar input: %v\n", err)
		return "", err
	}

	// Define a variável de ambiente com os inputs serializados
	os.Setenv("GIST_INPUT", string(inputJSON))

	result, err := executor.ExecuteGoCode(code)
	if err != nil {
		log.Printf("Erro ao serializar input: %v\n", err)
		return nil, err
	}

	fmt.Println(result)

	ba.Output.Data = result
	return result, nil
}

func (ba *BaseAction) SetOrder(order int) {
	ba.Order = order
}
