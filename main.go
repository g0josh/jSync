package main

import (
	"flag"
	"log"
	"os"
	"os/user"

	"github.com/g0josh/jsync/configmanager"
	// "github.com/karrick/godirwalk"
)

var configPath string
var defaultConfigPath string
var config *configmanager.Config

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

	// Parse config file passed in via flags
	// If it does not exist use the default config file at ~/.config/jsync/config.yaml
	// If default config file does not exist create a new empty one
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		if configPath != defaultConfigPath {
			log.Printf("%v not found. Using %v instead", configPath, defaultConfigPath)
		} else {
			log.Printf("%v not found. Creating a new one", defaultConfigPath)
		}
		// make a new config file
		if config, err = config.CreateEmptyConfig(defaultConfigPath); err != nil {
			log.Fatalf("Error while creating config file- %v", err)
		}
		configPath = defaultConfigPath
	}

	log.Printf("Using config file %v", configPath)
	config, err := config.ParseConfig(configPath)
	if err != nil {
		log.Fatalf("Error reading config file - %v", err)
	}
	log.Println(*config)
}
