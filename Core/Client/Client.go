package Client

import (
	"encoding/json"
	"fmt"
	"os"
)

var (
	ConfigFile = "./Assets/JSON/Config.json"
)

type Config struct {
	Keys []string
}

func GetConfig(file string) Config {
	var config Config
	ConfigFile, err := os.Open(file)
	if err != nil {
		fmt.Println("Config file not found." + err.Error())
	}
	defer ConfigFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(ConfigFile)
	jsonParser.Decode(&config)
	return config
}
