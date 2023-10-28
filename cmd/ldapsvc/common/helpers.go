package common

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func IsDir(path string) bool {
	fileInfo, err := os.Stat(path)

	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}

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
			panic(err)
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
