//go:generate go run -tags=dev data_generate.go

package config

import (
	"io/ioutil"
	"os"
	"sort"

	yaml "gopkg.in/yaml.v2"
)

// Host represents host information
type Host struct {
	Name     string
	Host     string
	Platform string
}

// Config represents yaml config with list of servers and ports
type Config struct {
	Tokens map[string]string
	PC     map[string]string
	XBOX   map[string]string
	Ports  []int
}

// AllHosts returns all hostnames to scan
func (cfg Config) AllHosts() []Host {
	result := make([]Host, 0)

	for name, host := range cfg.PC {
		result = append(result, Host{Name: name, Host: host, Platform: "PC"})
	}

	for name, host := range cfg.XBOX {
		result = append(result, Host{Name: name, Host: host, Platform: "XBOX"})
	}

	sort.Slice(result, func(i, j int) bool { return result[i].Name < result[j].Name })

	return result
}

func (cfg Config) MainToken() string {
	token := os.Getenv("MAIN_TOKEN")
	if token != "" {
		return token
	}

	return cfg.Tokens["main"]
}

// ReadYAML will read config yaml and parse it
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
