package trigger

type ActionFunc func(input interface{}) (output interface{}, err error)

type Trigger interface {
	Start()
	SetAction(action ActionFunc)
}

type BaseTrigger struct {
	actions []ActionFunc
}

func NewBaseTrigger() *BaseTrigger {
	return &BaseTrigger{}
}

func (bt *BaseTrigger) ExecuteActions(input interface{}) {
	var err error
	for _, action := range bt.actions {
		input, err = action(input)
		if err != nil {
			break
		}
	}
}

func (bt *BaseTrigger) SetActions(action ...ActionFunc) {
	bt.actions = append(bt.actions, action...)
}

func (bt *BaseTrigger) Start() {
	bt.ExecuteActions(nil)
}
