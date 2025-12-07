package main

import (
	"fmt"
	"log"
	"os"

	"kubeconfig/internal/cli"
	"kubeconfig/internal/models/config"

	"github.com/spf13/pflag"
	"gopkg.in/yaml.v3"
)

// Функция getFromFile - для получения содержимого файла по указанному пути
func getFromFile(cfg *config.KubeConfig, FilePath string) (config.KubeConfig, error) {
	// Получаем домашнюю директорию пользователя
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return config.KubeConfig{}, err
	}
	// Формируем полный путь к файлу
	path := fmt.Sprintf("%s/%s", homeDir, FilePath)
	// Читаем содержимое файла
	data, err := os.ReadFile(path)
	if err != nil {
		return config.KubeConfig{}, err
	}
	// Разбираем YAML содержимое в структуру KubeConfig
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return config.KubeConfig{}, err
	}
	// Возвращаем заполненную структуру KubeConfig
	return *cfg, nil
}

func main() {
	// Инициализация и парсинг флагов командной строки
	opts := cli.NewOptions()
	opts.AddFlags()
	pflag.Parse()

	// Создание нового экземпляра KubeConfig для хранения данных
	kubeconfig := config.NewKubeConfig()

	//TODO: Залить проект в GitHub под именем kubeconfig
	//TODO: Вести дальнейшую разработку в отдельной ветке feature/*
	//TODO: Добавить флаг и логику для получения версии утилиты (например, --version)
	//TODO: Добавить проверку на opts.FilesDir и обработку нескольких файлов
	//TODO: Реализовать логику объединения нескольких kubeconfig файлов в один '.kube/config'
	//TODO: Добавить логику сохранения объединенного kubeconfig в один общий файл '.kube/config'
	//TODO: Добавить логику сохранения текущего контекста и переключения между ними
	//TODO: Добавить обработку ошибок и логирование
	//TODO: Реализовать другие источники получения конфигурации Kubernetes (например, из Vault)
	//TODO: Актуализировать README.md с учетом новых возможностей

	// Чтение и разбор kubeconfig файла в структуру KubeConfig
	cfg, err := getFromFile(kubeconfig, opts.FilePath)
	if err != nil {
		log.Fatalf("Reading kubeconfig file is failed: %v", err)
	}
	// Валидация прочитанной конфигурации Kubernetes
	if err := cfg.ValidateKubeConfig(); err != nil {
		log.Fatalf("Invalid kubeconfig: %v", err)
	}
	// Форматирование и вывод конфигурации Kubernetes
	if err := cfg.PrettyYAML(); err != nil {
		log.Fatalf("Formatting kubeconfig is failed: %v", err)
	}
}



// Валидатор файлов конфигурации
// func isValidateFile(configFile string) bool {
// 	// Логика валидации конфигурационного файла
// 	validate, err := getFromFile(configFile)
// 	if err != nil {
// 		fmt.Errorf("Error reading config file:", err)
// 		return false
// 	}
	

// 	if validate == "" {
// 		return false
// 	}

// 	return true
// }


// Функция configFilesParser - для парсинга конфигурационных файлов
// func configFilesParser(configsDir string) (string, error) {
// 	// Цикл обработки файлов в директории configsDir

// 	return "", nil
// }
