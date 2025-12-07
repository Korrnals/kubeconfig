package cli

import (
	"github.com/spf13/pflag"
)

type Options struct {
	FilePath	string	`yaml:"filePath"` // Path to the directory containing Kubernetes configuration files
	FilesDir	string	`yaml:"filesDir"` // Directory containing Kubernetes configuration files
}

// Конструктор NewOptions - создает и возвращает новый экземпляр Options
func NewOptions() *Options {
	return &Options{}
}

// Метод AddFlags - добавляет флаги командной строки в указанный FlagSet
func (o *Options) AddFlags() {
	pflag.StringVarP(&o.FilePath, "filePath", "f", ".kube/config", "Path to the directory containing single Kubernetes configuration file")
	pflag.StringVarP(&o.FilesDir, "filesDir", "d", "", "Directory containing Kubernetes configuration files")
}