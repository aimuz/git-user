package main

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type User struct {
	Email        string   `yaml:"email"`
	Name         string   `yaml:"name"`
	IdentityFile string   `yaml:"identity_file"`
	GPGKey       string   `yaml:"gpg_key"`
	Default      bool     `yaml:"default"`
	AutoEnable   []string `yaml:"auto_enable"`
}

type Users map[string]User

func (u Users) Update() error {
	data, err := yaml.Marshal(u)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(dataPath, data, 0666)
}
