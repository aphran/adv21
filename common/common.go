package common

import (
    "errors"
    "fmt"
    "io/ioutil"
    "log"
    "os"

    "gopkg.in/yaml.v2"
)

type Config struct {
    SessionCookie string `yaml:"SessionCookie,omitempty"`
    InputPath     string `yaml:"InputPath,omitempty"`
}

const (
    InputURL = "https://adventofcode.com/2021/day/%d/input"
    ConfigFile = "../config.yaml"
)

var (
    ConfigData Config
)

func LoadConfig() {
    rawConfig, err := ioutil.ReadFile(ConfigFile)
    if err != nil {
        log.Fatalln("Fatal error,", err)
    }
    err = yaml.Unmarshal(rawConfig, &ConfigData)
    if err != nil {
        log.Fatalln("Error,", err)
    }
}

func OpenDayData(day int) (*os.File, error) {
    var err error
    if day < 1 || day > 24 {
        err = errors.New("The day number must be between 1 and 25")
        return nil, err
    }

    if ConfigData == (Config{}) {
        LoadConfig()
    }
    filePlace := fmt.Sprintf("%s/%d", ConfigData.InputPath, day)

    file, err := os.Open(filePlace)
    if err != nil {
        fmt.Println(err)
    }

    return file, err
}
