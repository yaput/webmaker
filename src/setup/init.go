package setup

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func LoadConfig(path string) (config Configuration, err error) {
	//filename is the path to the json config file
	file, err := os.Open(path)
	if err != nil {
		err = errors.New(fmt.Sprintf("Failed to Read File: %+v", err))
		return
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		err = errors.New("Failed To Decode Config File")
		return
	}
	return
}
