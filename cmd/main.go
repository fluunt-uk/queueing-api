package main

import (
	"gitlab.com/projectreferral/queueing-api/internal/api"
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE|os.O_CREATE, 0644)
	if err != nil {
	 log.Fatal(err)
	}
	defer f.Close()
	log.SetOutput(f)
	api.Init()
	api.SetupEndpoints()
}