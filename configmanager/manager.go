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
		Description string `yaml:"description,omitempty"`
		Path        string `yaml:"path,omitempty"`
	} `yaml:"directories,omitempty"`
}

func ReadConfig(path string) (Config, error) {
	var result Config
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return result, err
	}
	if err = yaml.Unmarshal(yamlFile, &result); err != nil {
		return result, err
	}
	return result, nil
}

func CreateEmptyConfig(path string) error {
	bytes, err := yaml.Marshal(&Config{})
	if err != nil {
		return err
	}
	dir, _ := filepath.Split(path)
	if _, statErr := os.Stat(dir); os.IsNotExist(statErr) {
		if err = os.MkdirAll(dir, 0760); err != nil {
			return err
		}
	}
	if err = ioutil.WriteFile(path, bytes, 0660); err != nil {
		log.Println("Creating")
		return err
	}
	return nil
}
