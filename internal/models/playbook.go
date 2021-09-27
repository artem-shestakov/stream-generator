package models

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Playbook struct {
	Name       string
	Hosts      string
	RemoteUser string `yaml:"remote_user"`
	Become     bool
	Roles      []Role
}

type Role struct {
	Role string
	Vars map[string]interface{}
}

type Playbooks struct {
	playbooks []Playbook
}

func (p *Playbooks) ReadFile() *Playbooks {
	var data []Playbook
	yamlFile, err := ioutil.ReadFile("example.yml")
	if err != nil {
		log.Printf("yamlFile get error %s", err.Error())
	}

	err = yaml.Unmarshal(yamlFile, &data)
	if err != nil {
		log.Printf("Unmarshall error %s", err.Error())
	}

	p.playbooks = data
	return p
}

type Stream struct {
	Name    string
	Listen  []string
	Servers []string
}

type Streams struct {
	Streams []Stream
}
