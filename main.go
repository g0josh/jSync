package main

import (
	"flag"
	"fmt"
	"log"
	"os/user"
	// "gopkg.in/yaml.v2"
	// "github.com/karrick/godirwalk"
)

var configPath string

func init() {
	user, err := user.Current()
	if err != nil {
		log.Fatalf("Could not get user's home directory - %v", err)
	}
	defaultConfigPath := user.HomeDir + "/.config/jsync/config.yaml"
	flag.StringVar(&configPath, "config", defaultConfigPath, "Config file to use")
}

func main() {
	flag.Parse()
	fmt.Println(configPath)
}
