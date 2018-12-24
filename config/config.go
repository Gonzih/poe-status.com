//go:generate go run -tags=dev data_generate.go

package config

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	PC    map[string]string
	XBOX  map[string]string
	Ports []int
}

func ReadYAML() (*Config, error) {
	file, err := DataDir.Open("config.yaml")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fdata, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	err = yaml.Unmarshal(fdata, &cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
