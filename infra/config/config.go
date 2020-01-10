package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"gopkg.in/yaml.v2"

	"clean_arch/infra"
)

// NewConfig  -
func NewConfig(dir string) (*infra.Config, error) {
	appMode := ReadEnv("APP_MODE", "dev")

	if dir == "" {
		dir, _ = os.Getwd()
	}

	cfg, err := LoadConfigFromYAML(fmt.Sprintf("%s/config/config.%s.yml", dir, appMode))
	if err != nil {
		return nil, err
	}
	// set config mode
	cfg.Mode = appMode
	if !strings.HasPrefix(cfg.Log.FileName, "/") {
		cfg.Log.FileName = path.Join(dir, cfg.Log.FileName)
	}
	//fmt.Printf("APP MODE:%s\n", cfg.Mode)
	return cfg, nil
}

// ReadEnv -
func ReadEnv(env string, valueDefault string) string {
	value := os.Getenv(env)
	if value != "" {
		return value
	}
	return valueDefault
}

// LoadConfigFromYAML -reads the file of the filename (in the same folder) and put it into the Config
func LoadConfigFromYAML(filename string) (*infra.Config, error) {
	var c infra.Config
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("readFile: %w", err)
	}
	err = yaml.Unmarshal(file, &c)

	if err != nil {
		return nil, fmt.Errorf("unmarshal: %w", err)
	}
	err = ValidateConfig(c)
	if err != nil {
		return nil, fmt.Errorf("validateConfig: %w", err)
	}
	return &c, nil
}

// ReadEnvConfig -
func ReadEnvConfig(cfg *infra.Config) error {
	// re write mode using env
	// cfg.Mode = ReadEnv("APP_MODE", "dev")
	return nil
}
