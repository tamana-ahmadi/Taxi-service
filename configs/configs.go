package configs

import (
	"Taxi_service/models"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
)

var AppSettings models.Configs

func ReadString() error {
	fmt.Println("Starting reading settings file")
	configfile, err := os.Open("configs/configs.json")
	if err != nil {
		return errors.New(fmt.Sprintf("Couldn't open config file. Error is: %s", err.Error()))
	}
	defer func(configfile *os.File) {
		err = configfile.Close()
		if err != nil {
			log.Fatal("Couldn't close config file. Error is: ", err.Error())
		}

	}(configfile)
	fmt.Println("Starting decoding settings file")
	err = json.NewDecoder(configfile).Decode(&AppSettings)
	if err != nil {
		return errors.New(fmt.Sprintf("Couldn't decode settings json file. Error is: %s", err.Error()))
	}
	return nil

}
