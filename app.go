package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yaput/webmaker/src/handler"

	"github.com/yaput/webmaker/src/setup"
)

func main() {
	cfg, err := setup.LoadConfig("config/config.development.json")
	if err != nil {
		panic(fmt.Sprint("Failed to load config: ", err))
	}

	// HANDLE EVEERY ENDPOINT HERE
	http.HandleFunc("/home", handler.HomePage)
	// END HANDLE
	log.Println("Serving on port:", cfg.Server.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Server.Port, nil))
}
