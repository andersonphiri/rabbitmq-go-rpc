package common

import (
	"github.com/joho/godotenv"
)
type RabbitMq struct {
	Username string `json:"user"`
	Password string `json:"password"`
	Port string `json:"port"`
	VirtualHostName string  `json:"virtual_host_name"`
}

type AppSettings struct{
	RabbitMq RabbitMq `json:"rabbitmq"`
}

func LoadDefaultAppSettings() (map[string]string, error) {
	return godotenv.Read()
}

func LoadMultipleAppSettingsFiles(filenames ...string) (map[string]string, error) {
	return godotenv.Read(filenames...)
}