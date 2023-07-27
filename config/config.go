package config

import (
	"flag"
	"github.com/spf13/viper"
	"log"
)

var (
	Listen   string
	SmtpHost string
	SmtpPort string
	SmtpUser string
	SmtpPass string
	Url      string
	CfgFile  string
)

func Init() {
	flag.StringVar(&CfgFile, "c", "config.yaml", "config file")
	flag.Parse()
	if CfgFile != "" {
		viper.SetConfigFile(CfgFile)
	} else {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath("./")
	}

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("read config failed:", err)
	}

	Listen = viper.GetString("default.listen")
	SmtpHost = viper.GetString("smtp.host")
	SmtpPort = viper.GetString("smtp.port")
	SmtpUser = viper.GetString("smtp.user")
	SmtpPass = viper.GetString("smtp.pass")
	Url = viper.GetString("domain.url")
}
