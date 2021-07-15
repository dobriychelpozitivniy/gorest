package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"gorest2/internal/app/apiserver"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")

}

func main() {
	flag.Parse()

	c := apiserver.NewConfig()
	_, err := toml.DecodeFile(configPath, c)
	if err != nil {
		log.Fatal(err)
	}

	if err := apiserver.Start(c); err != nil {
		log.Fatal(err)
	}

}

