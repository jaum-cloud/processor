package workflow

import (
	"github.com/jaum-cloud/processor/action"
	"github.com/jaum-cloud/processor/trigger"
)

type Workflow interface {
	SetTrigger(trigger trigger.Trigger)
	SetActions(action ...action.Action)
	Execute()
}

type WorkflowBase struct {
	ID      string
	Name    string
	Trigger trigger.Trigger
	Actions []action.Action
}

func NewWorkflowBase() *WorkflowBase {
	return &WorkflowBase{}
}

func (wf *WorkflowBase) SetTrigger(t trigger.Trigger) {
	wf.Trigger = t
}

func (wf *WorkflowBase) SetActions(actions ...action.Action) {
	wf.Actions = actions
}

func (wf *WorkflowBase) Execute() {
	var input interface{}
	var err error

	if wf.Trigger != nil {
		wf.Trigger.Start()
	}

	for _, action := range wf.Actions {
		input, err = action.Execute(input)
		if err != nil {
			break
		}
	}
}
