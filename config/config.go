package config

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)


// ConfigRedis : ..
type ConfigRedis struct {
	Host   string `yaml:"HOST"`
	Enable bool   `yaml:"ENABLE"`
}

// ConfigDatabase : ..
type ConfigDatabase struct {
	Driver   string `yaml:"DRIVER"`
	Port     string `yaml:"PORT"`
	Host     string `yaml:"HOST"`
	Username string `yaml:"USERNAME"`
	Password string `yaml:"PASSWORD"`
	Name     string `yaml:"NAME"`
}

// Config : ..
type Config struct {
	App         string         `yaml:"app"`
	Author      string         `yaml:"author"`
	Description string         `yaml:"description"`
	Redis       ConfigRedis    `yaml:"redis"`
	Database    ConfigDatabase `yaml:"database"`
	LogFilePath string         `yaml:"logfilepath"`
}

func (c *Config) InitConfig() *Config {

	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

// GetDatabaseURI : ..
func (c *Config) GetDatabaseURI() string {
	var uri string
	if c.Database.Driver == "mysql" {
		uri = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.Database.Username, c.Database.Password, c.Database.Host, c.Database.Port, c.Database.Name)
	} else {
		uri = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", c.Database.Host, c.Database.Port, c.Database.Username, c.Database.Name, c.Database.Password)
	}

	return uri
}
