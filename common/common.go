package common

import (
    "log"
    "io/ioutil"

    "gopkg.in/yaml.v2"
)

type Config struct {
    SessionCookie string `yaml:"SessionCookie,omitempty"`
}

const (
    InputPath = "input"
    InputURL = "https://adventofcode.com/2021/day/%d/input"
    SecretFile = "secrets.yaml"
)

var (
    ConfigData Config
)

func LoadConfig() {
    rawConfig, err := ioutil.ReadFile(SecretFile)
    if err != nil {
        log.Fatalln("Fatal error,", err)
    }
    err = yaml.Unmarshal(rawConfig, &ConfigData)
    if err != nil {
        log.Fatalln("Error,", err)
    }
}

