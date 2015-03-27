package config

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

func ReadConfigInto(config interface{}, configPath string) error {
	configBytes, err := ioutil.ReadFile(configPath)
	if err != nil {
		return errors.New(fmt.Sprintf("Can not read config file [%s]: %s", configPath, err))
	}
	err = json.Unmarshal(configBytes, config)
	if err != nil {
		return errors.New(fmt.Sprintf("Can not parse config file %s: %s", configPath, err))
	}
	return nil
}

func ReadCsv(filepath string) ([][]string, error) {
	csvFile, err := os.Open(filepath)
	if err != nil {
		return nil, err

	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)

	reader.FieldsPerRecord = -1

	rawCSVdata, err := reader.ReadAll()

	if err != nil {
		return nil, err
	}

	return rawCSVdata, nil

}
