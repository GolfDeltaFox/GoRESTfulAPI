package main

import (
	"log"
	"net/http"

	"app/route"

	"app/shared/database"
	"encoding/json"
	"app/shared/jsonconfig"
	"os"
)

func main() {
	// Load the configuration file
	jsonconfig.Load("config"+string(os.PathSeparator)+"config.json", config)
	// Connect to database
	database.Database.Connect(config.Database)

	database.Database.Init(config.Database)

	router := route.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}


// *****************************************************************************
// Application Settings
// *****************************************************************************

// config the settings variable
var config = &configuration{}

// configuration contains the application settings
type configuration struct {
	Database  database.Database   `json:"Database"`
}

// ParseJSON unmarshals bytes to structs
func (c *configuration) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}
