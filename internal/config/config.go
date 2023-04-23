package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

type Config struct {
	MmsUrl      string `yaml:"mms_url"`
	SmsFile     string `yaml:"sms_file"`
	ViclFile    string `yaml:"vicl_file"`
	MailFile    string `yaml:"mail_file"`
	BillFile    string `yaml:"bill_file"`
	SupportUrl  string `yaml:"support_url"`
	IncidentUrl string `yaml:"incident_url"`
	Port        string `yaml:"port"`
}

func GetConfig() *Config {

	config := &Config{}
	if err := cleanenv.ReadConfig("config.yml", config); err != nil {
		help, _ := cleanenv.GetDescription(config, nil)
		log.Println(help)

	}
	return config
}
