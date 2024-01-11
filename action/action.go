package action

import (
	"fmt"
	"github_automation/executor"
	"github_automation/github"
)

type Action interface {
	Execute(input interface{}) (output interface{}, err error)
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

func NewBaseAction(id string, order int) *BaseAction {
	return &BaseAction{
		ID:    id,
		Name:  id,
		Order: order,
	}
}

func (ba *BaseAction) Execute(input interface{}) (output interface{}, err error) {
	code, err := github.FetchGoCodeFromGist(ba.ID)
	if err != nil {
		return nil, nil
	}

	result, err := executor.ExecuteGoCode(code)
	if err != nil {
		return nil, nil
	}

	fmt.Println(ba.Order, input)

	return result, nil
}
