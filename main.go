package main

import (
	"flag"
	"github.com/godsun/proxypoolCheck/api"
	"github.com/godsun/proxypoolCheck/config"
	"github.com/godsun/proxypoolCheck/internal/app"
	"github.com/godsun/proxypoolCheck/internal/cron"
	"log"
	"net/http"
)

var configFilePath = ""

func main()  {
	go func() {
		http.ListenAndServe("0.0.0.0:6061", nil)
	}()

	// fetch configuration
	flag.StringVar(&configFilePath, "c", "", "path to config file: config.yaml")
	flag.Parse()
	if configFilePath == "" {
		configFilePath = "config.yaml"
	}
	err := config.Parse(configFilePath)
	if err != nil {
		log.Fatal(err, "\n\"Config file err. Exit\"")
		return
	}

	go app.InitApp()
	go cron.Cron()
	// Run
	api.Run()


}
