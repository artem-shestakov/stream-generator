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
	Roles      []string
	Vars       []map[string]string
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

// func (p *Playbooks) DeleteStreams() {
// 	for varIndex, playbookVar := range p.playbooks[0].Vars {
// 		if playbookVar["tcp_udp_nlb"] != nil {
// 			p.playbooks[0].Vars[varIndex] = p.playbooks[0].Vars[len(p.playbooks[0].Vars)-1]
// 			p.playbooks[0].Vars[len(p.playbooks[0].Vars)-1] = nil
// 			p.playbooks[0].Vars = p.playbooks[0].Vars[:len(p.playbooks[0].Vars)-1]
// 		}
// 	}
// }

// func (p *Playbooks) GetStreams() {
// 	var stream Stream
// 	for _, playbookVar := range p.playbooks[0].Vars {
// 		if playbookVar["tcp_udp_nlb"] != nil {
// 			switch reflect.TypeOf(playbookVar["tcp_udp_nlb"]).Kind() {
// 			case reflect.Slice:
// 				s := reflect.ValueOf(playbookVar["tcp_udp_nlb"])

// 				for i := 0; i < s.Len(); i++ {
// 					jsonString, err := json.Marshal(s.Index(i))
// 					if err != nil {
// 						log.Printf("JSON marshall error %s", err.Error())
// 					}
// 					json.Unmarshal(jsonString, &stream)
// 					fmt.Println(s.Index(i))
// 				}
// 			}
// 		}
// 	}
// }

type Stream struct {
	Name    string
	Listen  []string
	Servers []string
}

type Streams struct {
	Streams []Stream
}
