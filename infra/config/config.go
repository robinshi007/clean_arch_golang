package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"

	"clean_arch/infra"
)

// NewConfig  -
func NewConfig() (*infra.Config, error) {
	env := ReadEnv("APP_MODE", "dev")
	return ReadConfigFromYAML(fmt.Sprintf("config/config.%s.yml", env))
}

// ReadEnv -
func ReadEnv(env string, valueDefault string) string {
	value := os.Getenv(env)
	if value != "" {
		return value
	}
	return valueDefault
}

// ReadConfigFromYAML -reads the file of the filename (in the same folder) and put it into the Config
func ReadConfigFromYAML(filename string) (*infra.Config, error) {
	var c infra.Config
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.Wrap(err, "read error")
	}
	err = yaml.Unmarshal(file, &c)

	if err != nil {
		return nil, errors.Wrap(err, "unmarshal")
	}
	err = ValidateConfig(c)
	if err != nil {
		return nil, errors.Wrap(err, "validate config")
	}

	ReadEnvConfig(&c)

	return &c, nil
}

// ReadEnvConfig -
func ReadEnvConfig(cfg *infra.Config) error {
	// re write mode using env
	cfg.Mode = ReadEnv("APP_MODE", "dev")
	return nil
}
