package main

import (
	"fmt"
	"log"
	"os"

	"kubeconfig/internal/cli"
	"kubeconfig/internal/models/config"
	"kubeconfig/internal/utils"

	"github.com/spf13/pflag"
)

func main() {
	// Инициализация и парсинг флагов командной строки
	opts := cli.NewOptions()
	opts.AddFlags()
	pflag.Parse()

	// Создание нового экземпляра KubeConfig для хранения данных
	cfg := config.NewKubeConfig()

	// // Чтение и разбор kubeconfig файла в структуру KubeConfig
	// if err := cli.GetSingleConf(opts.FilePath, cfg); err != nil {
	// 	log.Fatalf("Reading kubeconfig file is failed: %v", err)
	// }

	// Генерация объединенного kubeconfig на основе прочитанных данных и опций
	if opts.DirPath != "" {
		if err := cli.GenKubeconfig(cfg, opts); err != nil {
			log.Fatalf("Generating kubeconfig is failed: %v", err)
		}
	}
	// Вывод версии утилиты
	if opts.Version {
        fmt.Println("kubeconfig v1.0.0")
        os.Exit(0)
    }

	// Выводит информацию о кластерах и контексте
	if opts.List {
    if err := utils.PrintSummary(cfg); err != nil {
        log.Fatalf("Failed to list: %v", err)
    }
    return
}
}
