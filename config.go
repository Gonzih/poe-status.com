package main

import (
	yaml "gopkg.in/yaml.v2"
)

type data struct {
	Pc    map[string]string
	Xbox  map[string]string
	Ports []int
}

func readServersYaml() *data {
	fdata, err := slurp("servers.yaml")
	must(err)

	data := &data{}
	err = yaml.Unmarshal(fdata, &data)
	must(err)

	return data
}
