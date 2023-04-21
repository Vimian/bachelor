package config

// Configuration is the structure of config.yaml
type Configuration struct {
	Server struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"server"`

	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
		SSLMode  string `yaml:"sslmode"`
	} `yaml:"database"`

	RabbitMQ struct {
		Host      string `yaml:"host"`
		Port      int    `yaml:"port"`
		User      string `yaml:"user"`
		Password  string `yaml:"password"`
		QueueName string `yaml:"queuename"`
	} `yaml:"rabbitmq"`

	AccountService struct {
		Host                       string `yaml:"host"`
		Port                       int    `yaml:"port"`
		GetAccountsByTimestampPath string `yaml:"getaccountsbytimestamppath"`
	} `yaml:"accountservice"`

	QueueingLoopInterval     int `yaml:"queueingloopinterval"`
	EnqueueDelay             int `yaml:"enqueuedelay"`
	EnqueueAccountsPerFectch int `yaml:"enqueueaccountsperfetch"`
	FetchAccountRewind       int `yaml:"fetchaccountrewind"`
}
