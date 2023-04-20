package config

import (
	"github.com/google/uuid"
)

// Configuration is the structure of config.yaml
type Configuration struct {
	RabbitMQ struct {
		Host      string `yaml:"host"`
		Port      int    `yaml:"port"`
		User      string `yaml:"user"`
		Password  string `yaml:"password"`
		QueueName string `yaml:"queuename"`
	} `yaml:"rabbitmq"`

	Redis struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Password string `yaml:"password"`
		DB       int    `yaml:"db"`
	} `yaml:"redis"`

	TransactionHistoryService struct {
		Host                  string `yaml:"host"`
		Port                  int    `yaml:"port"`
		Path                  string `yaml:"path"`
		PathUpdateStatusPart0 string `yaml:"pathupdatestatuspart0"`
		PathUpdateStatusPart1 string `yaml:"pathupdatestatuspart1"`
		StatusCompleted       int32  `yaml:"statuscompleted"`
		StatusFailed          int32  `yaml:"statusfailed"`
		TypeFailure           int32  `yaml:"typefailure"`
	} `yaml:"transactionhistoryservice"`

	AccountService struct {
		Host             string    `yaml:"host"`
		Port             int       `yaml:"port"`
		PathPart0        string    `yaml:"pathpart0"`
		PathPart1        string    `yaml:"pathpart1"`
		FailureAccountID uuid.UUID `yaml:"failureaccountid"`
	} `yaml:"accountservice"`
}
