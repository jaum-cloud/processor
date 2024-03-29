package trigger

import (
	"github.com/jaum-cloud/processor/action"
	"time"
)

type Trigger interface {
	Start()
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

func (bt *BaseTrigger) Start() {
	time.Sleep(1 * time.Second)
}
