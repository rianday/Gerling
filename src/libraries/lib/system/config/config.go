package config

import (
	"encoding/json"
	"io"
	"io/ioutil"

	// "libraries/lib/system/database"
	// "libraries/lib/system/email"
	// "libraries/lib/system/recaptcha"
	// "libraries/lib/system/server"
	// "libraries/lib/system/session"
	// "libraries/lib/system/view"
	"log"
	"os"
	"path/filepath"
)

// Parser must implement ParseJSON
type Parser interface {
	ParseJSON([]byte) error
}

var configFile string = "src/configs" + string(os.PathSeparator) + "development/application.json"

// Load the JSON config file
func Load(p Parser) {
	var err error
	var absPath string
	var input = io.ReadCloser(os.Stdin)
	if absPath, err = filepath.Abs(configFile); err != nil {
		log.Fatalln(err)
	}

	if input, err = os.Open(absPath); err != nil {
		log.Fatalln(err)
	}

	// Read the config file
	jsonBytes, err := ioutil.ReadAll(input)
	input.Close()
	if err != nil {
		log.Fatalln(err)
	}

	// Parse the config
	if err := p.ParseJSON(jsonBytes); err != nil {
		log.Fatalln(err)
	}
}

// *****************************************************************************
// Application Settings
// *****************************************************************************

// config the settings variable
var Cfg = &configuration{}

// configuration contains the application settings
type configuration struct {
	Database DbInfo `json:"Database"`
	// Email     email.SMTPInfo  `json:"Email"`
	// Recaptcha Info `json:"Recaptcha"`
	// Server    server.Server   `json:"Server"`
	// Session   session.Session `json:"Session"`
	// Template  view.Template   `json:"Template"`
	// View      view.View       `json:"View"`
}

// ParseJSON unmarshals bytes to structs
func (c *configuration) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}

func init() {
	Load(Cfg)
}
