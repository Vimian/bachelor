package config

import (
	"github.com/google/uuid"
)

// Configuration is the structure of config.yaml
type Configuration struct {
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

	TransactionHistoryService struct {
		Host            string `yaml:"host"`
		Port            int    `yaml:"port"`
		Path            string `yaml:"path"`
		StatusCompleted int32  `yaml:"statuscompleted"`
		TypeFailure     int32  `yaml:"typefailure"`
	} `yaml:"transactionhistoryservice"`

	AccountService struct {
		Host             string    `yaml:"host"`
		Port             int       `yaml:"port"`
		Path             string    `yaml:"path"`
		PathBalance      string    `yaml:"pathbalance"`
		FailureAccountID uuid.UUID `yaml:"failureaccountid"`
	} `yaml:"accountservice"`

	TransactionService struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
		Path string `yaml:"path"`
	} `yaml:"transactionservice"`
}
