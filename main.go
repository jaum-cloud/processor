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
		action.NewBaseAction("gandarfh/4855b65a615993d1903d48a4f43f1505"),
		action.NewBaseAction("gandarfh/a2ae888087e4b2c96823c54bbc0c62a0"),
	)

	// Executar o Workflow
	workflow.Execute()
}
