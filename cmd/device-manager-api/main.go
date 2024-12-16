package main

import (
	"log"

	"github.com/kelseyhightower/envconfig"
	"github.com/mendelgusmao/device-manager/internal/application"
)

func main() {
	config := application.Configuration{}

	if err := envconfig.Process("DEVICEMGR", &config); err != nil {
		log.Fatal(err.Error())
	}

	app := application.NewApplication(config)
	app.Run()
}
