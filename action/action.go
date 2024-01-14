package action

import (
	"log"

	"github.com/jaum-cloud/processor/executor"
	"github.com/jaum-cloud/processor/github"
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

	result, err := executor.ExecuteGoCode(code)
	if err != nil {
		log.Printf("Erro ao serializar input: %v\n", err)
		return nil, err
	}

	ba.Output.Data = result
	return result, nil
}

func (ba *BaseAction) SetOrder(order int) {
	ba.Order = order
}
