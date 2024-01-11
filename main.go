// main.go
package main

import (
	"github_automation/action"
	"github_automation/trigger"
	"github_automation/workflow"
)

func main() {

	// Criar uma instância de Workflow
	workflow := workflow.NewWorkflowBase()

	// Configurar Trigger e Actions
	workflow.SetTrigger(trigger.NewBaseTrigger())
	workflow.SetActions(
		action.NewBaseAction("gandarfh/4855b65a615993d1903d48a4f43f1505", 0),
		action.NewBaseAction("gandarfh/4855b65a615993d1903d48a4f43f1505", 1),
	)

	// Executar o Workflow
	workflow.Execute()
}
