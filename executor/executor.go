package executor

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"
)

const executionTimeout = 5 * time.Second // Limite de tempo para a execução do código

// ExecuteGoCode executa um código Go e retorna o resultado ou um erro.
func ExecuteGoCode(code string) (string, error) {
	// Salva o código em um arquivo temporário
	tmpFile, err := os.CreateTemp("", "*.go")
	if err != nil {
		return "", err
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.Write([]byte(code)); err != nil {
		return "", err
	}
	if err := tmpFile.Close(); err != nil {
		return "", err
	}

	// Definindo um contexto com timeout
	ctx, cancel := context.WithTimeout(context.Background(), executionTimeout)
	defer cancel()

	// Executa o código Go com o contexto
	cmd := exec.CommandContext(ctx, "go", "run", tmpFile.Name())

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {

		fmt.Println(err.Error())
		return stderr.String(), err
	}

	return stdout.String(), nil
}
