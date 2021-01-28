package main

import (
	"flag"
	"log"
	"os/user"
	"path/filepath"
)

func main() {

	var conf *Config

	usr, err := user.Current()
	if err != nil {
		log.Println(err)
	}

	// default config location
	configFile := filepath.Join(usr.HomeDir, ".factom-anchormaker/config.yaml")

	// check if custom config location passed as flag
	flag.StringVar(&configFile, "c", configFile, "config.yaml path")

	flag.Parse()

	log.Printf("Using config: %s\n", configFile)

	// load config
	if conf, err = NewConfig(configFile); err != nil {
		log.Fatal(err)
	}

	log.Printf("Starting anchoring system\n")

}
