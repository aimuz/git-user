package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type User struct {
	Email        string `yaml:"email"`
	Name         string `yaml:"name"`
	IdentityFile string `yaml:"identity_file"`
}

type Users map[string]User

func (u Users) Update() error {
	data, err := yaml.Marshal(u)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(dataPath, data, 0666)
}
