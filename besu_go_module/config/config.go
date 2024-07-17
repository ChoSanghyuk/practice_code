package config

import (
	_ "embed"
	"fmt"

	"gopkg.in/yaml.v3"
)

//go:embed config.yaml
var configByte []byte

type Configuration struct {
	Network  Network            `yaml:"network"`
	Nodes    map[string]Node    `yaml:"node"`
	Accounts map[string]Account `yaml:"account"`
}

type Network struct {
	ChainId  string `yaml:"chainId"`
	Url      string `yaml:"url"`
	GasLimit string `yaml:"gasLimit"`
}

type Node struct {
	Name       string `yaml:"name"`
	Url        string `yaml:"url"`
	Address    string `yaml:"accountAddress"`
	PublicKey  string `yaml:"nodekey"`
	PrivateKey string `yaml:"accountPrivateKey"`
}

type Account struct {
	Address    string `yaml:"address"`
	PrivateKey string `yaml:"privateKey"`
}

var Config Configuration

func init() {

	err := yaml.Unmarshal(configByte, &Config)
	if err != nil {
		fmt.Println(err)
	}
}
