package jsonutils

import (
	"encoding/json"
	"io"
	"os"
)

func Save(object interface{}, path string) error {
	jsonData, err := json.MarshalIndent(object, "", " ")
	if err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		return err
	}

	return nil
}

func Load(object interface{}, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	jsonData, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonData, object)
	if err != nil {
		return err
	}

	return nil
}
