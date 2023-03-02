package utils

import (
	"gopkg.in/yaml.v3"
	"log"
)

type AppConfig struct {
	Minio struct {
		Endpoint        string `yaml:"endpoint"`
		AccessKeyId     string `yaml:"accessKeyId"`
		SecretAccessKey string `yaml:"secretAccessKey"`
	}
	Server struct {
		HostName     string   `yaml:"hostName"`
		RemoteEnable bool     `yaml:"remoteEnable"`
		FileAddress  string   `yaml:"fileAddress"`
		M3u8Address  string   `yaml:"m3u8Address"`
		DataPath     string   `yaml:"dataPath"`
		Debug        bool     `yaml:"debug"`
		Nodes        []string `yaml:"nodes"`
	}
}

var c = AppConfig{}

func initConfig() {
	log.Println("init config")
	content := ReadFile(AppPath("config.yml"))
	err := yaml.Unmarshal(content, &c)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
func GetConfig() AppConfig {
	return c
}
