package executor

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"
)

// ExecuteGoCode executa um código Go e retorna o resultado ou um erro.
func ExecuteGoCode(code string) (string, error) {
	executionTimeout := 20 * time.Second // Defina o timeout conforme necessário

	// Criar um diretório temporário para o módulo
	moduleDir, err := os.MkdirTemp("", "go-module-")
	if err != nil {
		return "", err
	}
	defer os.RemoveAll(moduleDir) // Limpar após a execução

	// Criar um arquivo temporário para o código dentro do diretório do módulo
	tmpFile, err := os.Create(moduleDir + "/main.go")
	if err != nil {
		return "", err
	}
	defer os.Remove(tmpFile.Name())

	if err := os.Chdir(moduleDir); err != nil {
		return "", err
	}

	// Escrever o código no arquivo temporário
	if _, err := tmpFile.Write([]byte(code)); err != nil {
		return "", err
	}
	if err := tmpFile.Close(); err != nil {
		return "", err
	}

	cmdModInit := exec.Command("go", "mod", "init", "tempmodule")
	// cmdModInit.Dir = moduleDir // Defina o diretório de trabalho
	if err := cmdModInit.Run(); err != nil {
		return "", fmt.Errorf("erro ao inicializar o módulo Go: %w", err)
	}

	cmdModTidy := exec.Command("go", "mod", "tidy")
	// cmdModTidy.Dir = moduleDir // Defina o diretório de trabalho
	if err := cmdModTidy.Run(); err != nil {
		return "", fmt.Errorf("erro ao instalar dependencias no módulo Go: %w", err)
	}

	// dir, _ := os.ReadDir(moduleDir)

	// for _, file := range dir {
	// 	content, _ := os.ReadFile(file.Name())
	// 	fmt.Println(file.Name())
	// 	fmt.Println(string(content))
	// }

	// Definindo um contexto com timeout
	ctx, cancel := context.WithTimeout(context.Background(), executionTimeout)
	defer cancel()

	// Executar o código Go com o contexto
	cmd := exec.CommandContext(ctx, "go", "run", "main.go")
	// cmd := exec.CommandContext(ctx, "cat go.mod")
	// cmd.Dir = moduleDir // Definir o diretório do módulo como diretório de trabalho

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return stderr.String(), err
	}

	return stdout.String(), nil
}
