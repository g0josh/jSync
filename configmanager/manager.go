package configmanager

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Directories []struct {
		Name        string `yaml:"name,omitempty"`
		Path        string `yaml:"path,omitempty"`
		Description string `yaml:"description,omitempty"`
	} `yaml:"directories,omitempty"`
}

func (c *Config) ParseConfig(path string) (*Config, error) {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return c, err
	}
	if err = yaml.Unmarshal(yamlFile, &c); err != nil {
		return c, err
	}
	return c, nil
}

func (c *Config) CreateEmptyConfig(path string) (*Config, error) {
	bytes, err := yaml.Marshal(&Config{})

	if err != nil {
		return c, err
	}
	dir, _ := filepath.Split(path)
	if _, statErr := os.Stat(dir); os.IsNotExist(statErr) {
		if err = os.MkdirAll(dir, 0760); err != nil {
			return c, err
		}
	}
	if err = ioutil.WriteFile(path, bytes, 0660); err != nil {
		log.Println("Creating")
		return c, err
	}
	return c, nil
}
