package common

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

// FileExists returns true if filename exists, false otherwise
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// IsDir return true if path is a directory, false otherwise
func IsDir(path string) bool {
	fileInfo, err := os.Stat(path)

	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}

// SaveJson saves data to fn JSON file, returns an error if the save operation fails
func SaveJson(fn string, data interface{}) error {
	if file, err := json.MarshalIndent(data, "", " "); err != nil {
		return fmt.Errorf("error marhsaling json data (%v)", err.Error())
	} else {
		if err = os.WriteFile(fn, file, 0644); err != nil {
			return fmt.Errorf("error saving json file %s (%v)", fn, err.Error())
		}
	}
	return nil
}

// ReadJson reads the json file specified by fn
func ReadJson(fn string, data interface{}) error {
	// defer the closing of our jsonFile so that we can parse it later on
	jsonFile, err := os.Open(fn)
	defer func() {
		if err := jsonFile.Close(); err != nil {
			log.Printf("error closing file %s (%v)", fn, err)
		}
	}()

	if err != nil {
		return err
	}
	if byteValue, err := io.ReadAll(jsonFile); err != nil {
		return err
	} else {
		return json.Unmarshal(byteValue, &data)
	}

}
