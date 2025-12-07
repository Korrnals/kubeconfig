package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Структура KubeConfig для представления конфигурации Kubernetes
type KubeConfig struct {
	APIVersion    	string            `yaml:"apiVersion"`
	Clusters      	[]ClusterEntry	  `yaml:"clusters"`
	Contexts      	[]ContextEntry	  `yaml:"contexts"`
	CurrentContext	string            `yaml:"current-context"`
	Kind          	string            `yaml:"kind"`
	Preferences   	map[string]any    `yaml:"preferences"`
	Users         	[]UserEntry		  `yaml:"users"`
}

// Структура Cluster для представления информации о кластере
type Cluster struct {
	CertAuthData	string	`yaml:"certificate-authority-data"`
	Server    		string	`yaml:"server"`
}

type ClusterEntry struct {
	Name 		string		`yaml:"name"`
	Cluster		Cluster		`yaml:"cluster"`
}

// Структура Context для представления контекста Kubernetes
type Context struct {
	Cluster 	string	`yaml:"cluster"`
	User    	string	`yaml:"user"`
}

type ContextEntry struct {
	Name		string	`yaml:"name"`
	Context		Context	`yaml:"context"`
}

// Структура User для представления информации о пользователе
type User struct {
	ClientCertData	string	`yaml:"client-certificate-data"`
	ClientKeyData	string	`yaml:"client-key-data"`
}

type UserEntry struct {
	Name	string	`yaml:"name"`
	User	User	`yaml:"user"`
}

// ------------------------------------------------------- //
// Методы структур можно добавлять здесь при необходимости //

// Конструктор KubeConfig
func NewKubeConfig() *KubeConfig {
	return &KubeConfig{}
}

// Валидатор конфигурации Kubernetes
func (cfg *KubeConfig) ValidateKubeConfig() error {
	// Пример простой валидации: проверка наличия кластеров, контекстов и пользователей
	if len(cfg.Clusters) == 0 {
		return fmt.Errorf("no clusters defined in kubeconfig")
	}
	if len(cfg.Contexts) == 0 {
		return fmt.Errorf("no contexts defined in kubeconfig")
	}
	if len(cfg.Users) == 0 {
		return fmt.Errorf("no users defined in kubeconfig")
	}
	if cfg.CurrentContext == "" {
		return fmt.Errorf("current-context is not set in kubeconfig")
	}
	return nil
}

// Функция PrettyYAML - для форматирования и вывода конфигурации Kubernetes
func (cfg *KubeConfig) PrettyYAML() error {
	enc := yaml.NewEncoder(os.Stdout)
	enc.SetIndent(2)
	err := enc.Encode(cfg)
	if err != nil {
		return err
	}
	defer enc.Close()

	return nil
}