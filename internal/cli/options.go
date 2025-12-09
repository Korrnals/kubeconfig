package cli

import (
	"github.com/spf13/pflag"
)

type Options struct {
	FilePath	string	`yaml:"filePath"` 	// Path to the directory containing Kubernetes configuration files
	DirPath		string	`yaml:"dirPath"` 	// Directory containing Kubernetes configuration files
	OutputFile	string	`yaml:"outputFile"`	// File for saving the global configuration (default - '~/.kube/config')
	Version     bool   						// Show util version
    List        bool 						// List all clusters and contexts
}

// Конструктор NewOptions - создает и возвращает новый экземпляр Options
func NewOptions() *Options {
	return &Options{}
}

// Метод AddFlags - добавляет флаги командной строки в указанный FlagSet
func (o *Options) AddFlags() {
	pflag.StringVarP(&o.FilePath, "filePath", "f", "", "Path to the directory containing single Kubernetes configuration file")
	pflag.StringVarP(&o.DirPath, "dirPath", "d", "", "Directory containing Kubernetes configuration files")
	pflag.StringVarP(&o.OutputFile, "output", "o", "~/.kube/config", "Output file path (if empty — print to stdout)")
	pflag.BoolVarP(&o.Version, "version", "v", false, "Show version")
    pflag.BoolVarP(&o.List, "list", "l", false, "List all clusters and contexts")
}
