package main

import (
	"encoding/json"
	"fmt"
	sysconfig "libraries/lib/system/config"
	"libraries/lib/system/database"
	"libraries/lib/system/email"
	"libraries/lib/system/recaptcha"
	"libraries/lib/system/server"
	"libraries/lib/system/session"
	"libraries/lib/system/view"
)

func main() {
	//basePath, err := os.Getwd()
	sysconfig.Load(config)

	fmt.Println(config.Database.MySQL.Username)
	//fmt.Println("configs"+string(os.PathSeparator)+"development/application.json")
}

// *****************************************************************************
// Application Settings
// *****************************************************************************

// config the settings variable
var config = &configuration{}

// configuration contains the application settings
type configuration struct {
	Database  database.Info   `json:"Database"`
	Email     email.SMTPInfo  `json:"Email"`
	Recaptcha recaptcha.Info  `json:"Recaptcha"`
	Server    server.Server   `json:"Server"`
	Session   session.Session `json:"Session"`
	Template  view.Template   `json:"Template"`
	View      view.View       `json:"View"`
}

// ParseJSON unmarshals bytes to structs
func (c *configuration) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}