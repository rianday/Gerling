package main

import (
	"fmt"
	"libraries/models/test"
)

// "libraries/lib/system/config"
// "libraries/lib/system/database"
// "libraries/lib/system/email"
// "libraries/lib/system/recaptcha"
// "libraries/lib/system/server"
// "libraries/lib/system/session"
// "libraries/lib/system/view"
// "libraries/models/test"

func main() {
	var tesst = test.Mahasiswa{}

	result, err := tesst.Delete()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("delete successful : ", result)
	}

	result, err = tesst.Create()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("insert successful : ", result)
	}

	result, err = tesst.Update()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("update successful : ", result)
	}

	result2 := tesst.Read()
	for i, data := range result2 {
		fmt.Printf("index (%d) : %s\n", i, data.Name)
	}
	// fmt.Println(result2[0].Name)
}

// // *****************************************************************************
// // Application Settings
// // *****************************************************************************

// // config the settings variable
// var config = &configuration{}

// // configuration contains the application settings
// type configuration struct {
// 	Database  database.Info   `json:"Database"`
// 	Email     email.SMTPInfo  `json:"Email"`
// 	Recaptcha recaptcha.Info  `json:"Recaptcha"`
// 	Server    server.Server   `json:"Server"`
// 	Session   session.Session `json:"Session"`
// 	Template  view.Template   `json:"Template"`
// 	View      view.View       `json:"View"`
// }

// // ParseJSON unmarshals bytes to structs
// func (c *configuration) ParseJSON(b []byte) error {
// 	return json.Unmarshal(b, &c)
// }
