// main.go
package main

import (
	"github_automation/trigger"
)

func main() {
	// Exemplo de inicialização e uso de um trigger
	t := trigger.NewBaseTrigger() // Aqui, estamos usando um TimeTrigger como exemplo

	t.SetActions(
		trigger.ExecuteGistCode("gandarfh/4855b65a615993d1903d48a4f43f1505"),
		trigger.ExecuteGistCode("gandarfh/4855b65a615993d1903d48a4f43f1505"),
	)

	// Iniciar um trigger
	t.Start()
}
