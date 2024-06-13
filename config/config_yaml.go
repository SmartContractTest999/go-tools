package config

import (
	"errors"
	"gopkg.in/yaml.v3"
	"os"
)

func LoadYaml(fileName string, v any) error {
	yamlFile, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yamlFile, v)
	if err != nil {
		return err
	}

	return nil
}

func UpdateYaml(fileName string, v any) error {
	if v == nil {
		return errors.New("config data is nil")
	}
	data, err := yaml.Marshal(v)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, data, 0777)
}
