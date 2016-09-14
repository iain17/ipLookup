package main

import (
	"runtime"
	"github.com/iain17/ipLookup/src/utils/logger"
	"github.com/iain17/ipLookup/src/config"
	"github.com/iain17/ipLookup/src/environment"
	"github.com/iain17/ipLookup/src/rest"
	"github.com/oschwald/geoip2-golang"
)

func main() {
	// Add the Stdout logger
	logger.AddOutput(logger.Stdout{
		MinLevel: logger.ERROR, //logger.DEBUG,
		Colored:  true,
	})

	// Load settings from config.toml in working directory
	settings := config.Load("./config.toml")

	// Maximize CPU usage for maximum performance
	cores := runtime.NumCPU()
	logger.Infof("Running a server with: %d CPU cores. \n", cores)
	runtime.GOMAXPROCS(cores)

	//Load maxmind db
	db, err := geoip2.Open("GeoIP2-City.mmdb")
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	environment.SetEnvironment(&environment.Environment{
		Config:				settings,
		Maxmind:			db,
		//MongoDB:			mongoSession,
		//Redis:				redis,
	})

	//Starts the rest part of our program
	go rest.Start()

	// select{} stops execution of the program until all goroutines close
	select {}
}