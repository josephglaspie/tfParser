package main

import (
	"fmt"
	"log"

	"github.com/hashicorp/hcl/v2/hclsimple"
)

type NewConfig struct {
	Variables []VariablesConfig `hcl:"variable,block"`
}

type VariablesConfig struct {
	ProjectName string `hcl:"projectname,label"`
	Default     string `hcl:"default"`
}

type Config struct {
	Service ServiceConfig `hcl:"service,block"`
}

type ServiceConfig struct {
	Protocol   string          `hcl:"protocol,label"`
	Type       string          `hcl:"type,label"`
	ListenAddr string          `hcl:"listen_addr"`
	Foo        string          `hcl:"foo"`
	Processes  []ProcessConfig `hcl:"process,block"`
}

type ProcessConfig struct {
	Type    string   `hcl:"type,label"`
	Command []string `hcl:"command"`
}

func main() {
	var config Config
	var config2 NewConfig
	err := hclsimple.DecodeFile("config.hcl", nil, &config)
	if err != nil {
		log.Fatalf("Failed to load configuration: %s", err)
	}
	err = hclsimple.DecodeFile("workspace_configurations.hcl", nil, &config2)
	if err != nil {
		log.Fatalf("Failed to load configuration: %s", err)
	}
	log.Printf("Configuration is %#v", config)
	fmt.Println(config.Service.Type)
	fmt.Println(config.Service.Foo)

	fmt.Println(config2.Variables)

}
