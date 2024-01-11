package trigger

import (
	"github_automation/action"
	"time"
)

type Trigger interface {
	Start()
	SetActions(action ...action.Action)
}

type BaseTrigger struct {
	Actions []action.Action
}

func NewBaseTrigger() *BaseTrigger {
	return &BaseTrigger{}
}

func (bt *BaseTrigger) ExecuteActions(input interface{}) {
	var err error
	for _, action := range bt.Actions {
		input, err = action.Execute(input)
		if err != nil {
			break
		}
	}
}

func (bt *BaseTrigger) SetActions(actions ...action.Action) {
	bt.Actions = actions
}

func (bt *BaseTrigger) Start() {
	time.Sleep(3 * time.Second)
	// bt.ExecuteActions(nil)
}
