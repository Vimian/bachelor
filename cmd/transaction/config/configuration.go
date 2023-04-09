package config

// Configuration is the structure of config.yaml
type Configuration struct {
	Server struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"server"`

	RabbitMQ struct {
		Host      string `yaml:"host"`
		Port      int    `yaml:"port"`
		User      string `yaml:"user"`
		Password  string `yaml:"password"`
		QueueName string `yaml:"queuename"`
	} `yaml:"rabbitmq"`
}
