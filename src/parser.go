package src

import (
	"os"

	"github.com/rs/zerolog/log"
	yaml "gopkg.in/yaml.v2"
)

type RootFSManifest struct {
	Image           string `yaml:"image"`
	TargetDirectory string `yaml:"target_directory"`
	RootFsRegistry  string `yaml:"rootfs_registry"`
	Context         string `yaml:"context"`
	DockerfilePath  string `yaml:"docker_file"`
}

type Parser interface {
	ParseYamlFile(configFile string) (RootFSManifest, error)
}

type ParseHandler struct{}

func (parser *ParseHandler) ParseYamlFile(configFile string) (RootFSManifest, error) {
	var config RootFSManifest
	data, err := os.ReadFile("config.yaml")

	if err != nil {
		log.Err(err).Msg("Error reading YAML file:")
		return config, err
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Err(err).Msg("Error parsing YAML:")
		return config, err
	}

	return config, nil
}
