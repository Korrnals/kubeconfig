package utils

import (
	"bytes"
	"fmt"
	"kubeconfig/internal/models/config"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

// Вспомогательная функция PrettyYAML - для форматированного вывода конфигурации Kubernetes
func PrettyYAML(cfg *config.KubeConfig) error {
	enc := yaml.NewEncoder(os.Stdout)
	enc.SetIndent(2)
	err := enc.Encode(cfg)
	if err != nil {
		return err
	}
	defer enc.Close()

	return nil
}

// Вспомогательная функция PrettyYAMLBytes - для форматированной записи в файл
func PrettyYAMLBytes(cfg *config.KubeConfig) ([]byte, error) {
    var buf bytes.Buffer
    enc := yaml.NewEncoder(&buf)
    enc.SetIndent(2)
    if err := enc.Encode(cfg); err != nil {
        return nil, err
    }
    enc.Close()
    return buf.Bytes(), nil
}

// Вспомогательная функция isKubeConfigFile - определяет, является ли файл Kubeconfig'ом
func IsKubeConfigFile(content []byte) (bool, error) {
	// Получение контента и преобразование в строку
	file := string(content)
	// Проверка, если есть содержимое, то возвращается значение 'true'
	hasAPIVersion := strings.Contains(file, "apiVersion: v1")
	hasKind := strings.Contains(file, "kind: Config")

	return hasAPIVersion && hasKind, nil
}

// Вспомогательная функция ReplaceUserName - заменяет базовое имя пользователя - на уникальное
func ReplaceUserName(cfg *config.KubeConfig) {
	if cfg.CurrentContext == "" {
        return
    }

    newName := cfg.CurrentContext

    for i, user := range cfg.Users {
        if user.Name == "kubernetes-admin" {
            cfg.Users[i].Name = newName

            // Обновляем ссылки в контекстах
            for j, ctx := range cfg.Contexts {
                if ctx.Context.User == "kubernetes-admin" {
                    cfg.Contexts[j].Context.User = newName
                }
            }
        }
    }
}

// Вспомогательная функция PrintSummary - формирует вывод для флага '--list'
// PrintSummary — выводит список кластеров и контекстов
func PrintSummary(cfg *config.KubeConfig) error {
    // Если cfg пустой — пытаемся загрузить из ~/.kube/config
    if cfg == nil || (len(cfg.Clusters) == 0 && len(cfg.Contexts) == 0 && len(cfg.Users) == 0 && cfg.CurrentContext == "") {
        home, err := os.UserHomeDir()
        if err != nil {
            return fmt.Errorf("cannot get home directory: %w", err)
        }
        defaultPath := filepath.Join(home, ".kube", "config")

        data, err := os.ReadFile(defaultPath)
        if err != nil {
            return fmt.Errorf("no config loaded and %s not found", defaultPath)
        }

        cfg = config.NewKubeConfig()
        if err := yaml.Unmarshal(data, cfg); err != nil {
            return fmt.Errorf("failed to parse %s: %w", defaultPath, err)
        }

        log.Printf("Loaded current config from: %s\n---\n", defaultPath)
    }

    if cfg.CurrentContext == "" {
        fmt.Println("Current context: <none>")
    } else {
        fmt.Printf("Current context: %s\n", cfg.CurrentContext)
    }
    fmt.Println()

    fmt.Println("Clusters:")
    if len(cfg.Clusters) == 0 {
        fmt.Println("  <none>")
    }
    for _, c := range cfg.Clusters {
        fmt.Printf("  • %s → %s\n", c.Name, c.Cluster.Server)
    }

    fmt.Println("\nContexts:")
    if len(cfg.Contexts) == 0 {
        fmt.Println("  <none>")
    }
    for _, ctx := range cfg.Contexts {
        mark := ""
        if ctx.Name == cfg.CurrentContext {
            mark = " ← current"
        }
        fmt.Printf("  • %s%s\n", ctx.Name, mark)
    }

    return nil
}