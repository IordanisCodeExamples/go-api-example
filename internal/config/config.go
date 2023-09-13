// Package config contains the logic for the configuration of the service
package config

/*
   return stuct of a yaml config file
*/

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

/*
	Example of a config file
	server:
		port: ":8888"
*/

// AppConf contains all main structs
type AppConf struct {
	Server ServerConfig `yaml:"server"`
	DB     DB           `yaml:"db"`
	Kafka  Kafka        `yaml:"kafka"`
	GRPC   GRPC         `yaml:"grpc"`
}

// ServerConfig contains the data for the micro servise server
type ServerConfig struct {
	Port string `yaml:"port"`
}

// DB contains the data for the elastic search cluster connection
type DB struct {
	URL      string `yaml:"url"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

// Kafka contains the data for the kafka cluster connection
type Kafka struct {
	GroupID          string `yaml:"group_id"`
	Server           string `yaml:"server"`
	InsertMovieTopic string `yaml:"insert_movie_topic"`
}

// GRPC contains the data for the grpc server
type GRPC struct {
	Port string `yaml:"port"`
}

// GetAppConfig reads a spefic file and return the yaml format of it
// return ServerConfig struct yaml format of the config file
func GetAppConfig(path string) (*AppConf, error) {
	var c AppConf

	yamlFile, openfileError := os.ReadFile(path)
	if openfileError != nil {
		return nil, fmt.Errorf(
			"internal_config_getappconfig_open_file %s %w",
			path,
			openfileError,
		)
	}

	err := yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		return nil, fmt.Errorf("internal_config_getappconfig_yaml_unmarshal %w", err)
	}

	return &c, nil
}
