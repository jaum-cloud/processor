package action

type Action interface {
	Exec()
	SetInput()
	SetOutput()
}

type BaseAction struct {
}
