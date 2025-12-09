package config

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

// Конструктор NewKubeConfig - создаёт новый пустой KubeConfig и возвращает указатель на него.
// Возвращает указатель, чтобы избежать копирования структуры.
func NewKubeConfig() *KubeConfig {
    return &KubeConfig{
        APIVersion: "v1",
        Kind:       "Config",
        // Можно задать дефолтные значения, если нужно
        Preferences: make(map[string]any),
        Clusters:    []ClusterEntry{},
        Contexts:    []ContextEntry{},
        Users:       []UserEntry{},
    }
}

// Валидатор конфигурации Kubernetes
// TODO: Пересмотреть необходимость данного метода, и если да - то необходимо переработать логику
// func (cfg *KubeConfig) ValidateKubeConfig(file string) error {
// 	// Пример простой валидации: проверка наличия кластеров, контекстов и пользователей
// 	if len(cfg.Clusters) == 0 {
// 		return fmt.Errorf("no clusters defined in kubeconfig")
// 	}
// 	if len(cfg.Contexts) == 0 {
// 		return fmt.Errorf("no contexts defined in kubeconfig")
// 	}
// 	if len(cfg.Users) == 0 {
// 		return fmt.Errorf("no users defined in kubeconfig")
// 	}
// 	if cfg.CurrentContext == "" {
// 		return fmt.Errorf("current-context is not set in kubeconfig")
// 	}
// 	return nil
// }

// Метод Merg - для объединения двух конфигураций Kubernetes
func (cfg *KubeConfig) Merg(other *KubeConfig) {
	cfg.mergClusters(other.Clusters)
	cfg.mergeContexts(other.Contexts)
	cfg.mergeUsers(other.Users)
	if other.CurrentContext != "" {
		cfg.CurrentContext = other.CurrentContext
	}
}

// Метод mergeClusters - для слияния кластеров
func (cfg *KubeConfig) mergClusters(from []ClusterEntry) {
	// Проходим по каждому кластеру из другой конфигурации
	for _, NewCluster := range from {
		exists := false
		// Если кластер уже существует, перезаписываем его
		for i, c := range cfg.Clusters {
			if c.Name == NewCluster.Name {
				cfg.Clusters[i] = NewCluster // перезаписываем
				exists = true
				break
			}
		}
		// Если кластер не существует, добавляем его
		if !exists {
			cfg.Clusters = append(cfg.Clusters, NewCluster)
		}
	}
}

// Метод mergeContexts - для слияния контекстов
func (cfg *KubeConfig) mergeContexts(from []ContextEntry) {
	// Проходим по каждому контексту из другой конфигурации
	for _, NewContext := range from {
		exists := false
		// Если контекст уже существует, перезаписываем его
		for i, c := range cfg.Contexts {
			if c.Name == NewContext.Name {
				cfg.Contexts[i] = NewContext // перезаписываем
				exists = true
				break
			}
		}
		// Если контекст не существует, добавляем его
		if !exists {
			cfg.Contexts = append(cfg.Contexts, NewContext)
		}
	}
}

// Метод mergeUsers - для слияния пользователей
func (cfg *KubeConfig) mergeUsers(from []UserEntry) {
	// Проходим по каждому пользователю из другой конфигурации
	for _, NewUser := range from {
		exists := false
		// Если контекст уже существует, перезаписываем его
		for i, c := range cfg.Users {
			if c.Name == NewUser.Name {
				cfg.Users[i] = NewUser // перезаписываем
				exists = true
				break
			}
		}
		// Если контекст не существует, добавляем его
		if !exists {
			cfg.Users = append(cfg.Users, NewUser)
		}
	}
}
