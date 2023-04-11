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
		PathCreate            string `yaml:"pathcreate"`
		PathUpdateStatusPart0 string `yaml:"pathupdatestatuspart0"`
		PathUpdateStatusPart1 string `yaml:"pathupdatestatuspart1"`
		StatusCompleted       int32  `yaml:"statuscompleted"`
		StatusFailed          int32  `yaml:"statusfailed"`
	} `yaml:"transactionhistoryservice"`

	AccountService struct {
		Host      string `yaml:"host"`
		Port      int    `yaml:"port"`
		PathPart0 string `yaml:"pathpart0"`
		PathPart1 string `yaml:"pathpart1"`
	} `yaml:"accountservice"`

	FailureAccountID uuid.UUID `yaml:"failureaccountid"`
}
