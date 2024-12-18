package filemanipulator

import (
	"encoding/json"
	"fmt"
	"os"
)

const fileName = "settings.json"

type Settings struct {
	From        string	`json:"from"`
	Destination string	`json:"destination"`
	Extension	string 	`json:"extension"`
}

func ReadSettings() Settings {
	settings, err := os.ReadFile(fileName)

	if err != nil {
		fmt.Println(err)
	}

	var settingsReaded Settings;

	json.Unmarshal(settings, &settingsReaded)

	return settingsReaded
}
