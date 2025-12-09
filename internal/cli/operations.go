package cli

import (
	"fmt"
	"kubeconfig/internal/models/config"
	"kubeconfig/internal/utils"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

func GenKubeconfig(cfg *config.KubeConfig, opts *Options) error {
	// Получаем список файлов в указанной директории
	filesList, err := os.ReadDir(opts.DirPath)
	if err != nil {
		return fmt.Errorf("error read directory: %s", err)
	}
	if len(filesList) == 0 {
		return fmt.Errorf("no files found in directory: %s", opts.DirPath)
	}
	// Инициализация временной структуры для работы в цикле
	tempContent := config.NewKubeConfig()

	// Проходим по каждому файлу в директории, записывая его содержимое в общую структуру KubeConfig
	for _, file := range filesList {
		// Получаем имя файла
		f := file.Name()
		// Формирование полного пути до файла
		path := filepath.Join(opts.DirPath, f)
		// Получение одиночной конфигурации в формате YAML
		if err := GetSingleConf(path, tempContent); err != nil {
			return fmt.Errorf("error in 'cli.GetFromFile()' func: %s", err)
		}
		// Переименование базового имени пользователя (если не уникальное)
		utils.ReplaceUserName(tempContent)

		// Добавление временной структуры - в общую.
		cfg.Merg(tempContent)
		
	}
	// Форматируем готовый конфиг в корректный формат с отступами
	kubeconfig, err := utils.PrettyYAMLBytes(cfg)
	if err != nil {
		return fmt.Errorf("error formating []bytes yaml config: %s", err)
	}

	// Если флаг '--output' передан и не пустой (по умолчанию - '$HOME/.kube/config')
	if opts.OutputFile != "" {
		if strings.HasPrefix(opts.OutputFile, "~") {
			homeDir, _ := os.UserHomeDir()
			opts.OutputFile = filepath.Join(homeDir, ".kube", "config")
		}
		// То записывает в указанный файл
		if err := os.WriteFile(opts.OutputFile, kubeconfig, 0644); err != nil {
			return fmt.Errorf("error write []bytes yaml config to file: %s", err)
		}
		fmt.Printf("The '%s' file has been created or updated!\n", opts.OutputFile)
	} else {
		// Преобразовывем содержимое структуры в корректный вид с правильными отступами
		if err := utils.PrettyYAML(cfg); err != nil {
			return fmt.Errorf("error 'prettyYAML()' func: %s", err)
		}
	}
	
	return nil
}


// GetFromFile - функция для получения содержимого файла по указанному пути
func GetSingleConf(filePath string, cfg *config.KubeConfig) error {
	// Читаем содержимое файла
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("error read file: %s", err)
	}

	// Проверка (это kubeconfig?) содержимого файла, является ли он конфигом или нет
    if ok, err := utils.IsKubeConfigFile(data); err != nil {
        return fmt.Errorf("validation error for %s: %w", filePath, err)
    } else if !ok {
        log.Printf("Skipping %s: not a kubeconfig", filePath)
        return nil  // ← просто пропускаем!
    }

	// Разбираем YAML содержимое в структуру KubeConfig
	if err = yaml.Unmarshal(data, cfg); err != nil {
		return fmt.Errorf("error unmarshaling to 'GetFromFile()' func: %s", err)
	}
	// Логируем факт загрузки конфига
	// log.Printf("Loaded config from: %s\n---\n", filePath)
	
	// // Преобразовывем содержимое структуры в корректный вид с правильными отступами
	// if err := utils.PrettyYAML(cfg); err != nil {
	// 	return fmt.Errorf("error 'prettyYAML()' func: %s", err)
	// }

	return nil
}
