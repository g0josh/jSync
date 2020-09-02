package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/g0josh/jsync/configmanager"
	// "github.com/karrick/godirwalk"
)

var configPath string
var defaultConfigPath string

func init() {
	user, err := user.Current()
	if err != nil {
		log.Fatalf("Could not get user's home directory - %v", err)
	}
	defaultConfigPath = user.HomeDir + "/.config/jsync/config.yaml"
	flag.StringVar(&configPath, "config", defaultConfigPath, "Config file to use")
}

func main() {
	flag.Parse()
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		if configPath != defaultConfigPath {
			log.Printf("%v not found. Using %v instead", configPath, defaultConfigPath)
		} else {
			log.Printf("%v not found. Creating a new one", defaultConfigPath)
		}
		// make a new config file
		if confErr := configmanager.CreateEmptyConfig(defaultConfigPath); confErr != nil {
			log.Fatalf("Error while creating config file- %v", confErr)
		}
	}
	fmt.Println(configPath)
}
