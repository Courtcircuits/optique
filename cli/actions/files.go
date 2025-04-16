package actions

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func ClearIgnoredFiles() error {
	ignore, err := ReadIgnoreSection()
	if err != nil {
		return err
	}
	for _, file := range ignore {
		err = DeleteFile(file)
		if err != nil {
			return err
		}
	}
	fmt.Println("Ignored files cleared")
	return nil
}

func ReadIgnoreSection() ([]string, error) {
	type Ignore struct {
		Ignore []string `json:"ignore"`
	}

	config := Ignore{}
	optiqueConfigFile := "optique.json"
	optiqueConfig, err := os.ReadFile(optiqueConfigFile)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(optiqueConfig, &config)
	if err != nil {
		return nil, err
	}
	return config.Ignore, nil
}

func DeleteFile(glob string) error {
	matches, err := filepath.Glob(glob)
	if err != nil {
		return err
	}
	for _, match := range matches {
		err = os.Remove(match)
		if err != nil {
			return err
		}
	}
	return nil
}

func ReplaceInAllFiles(old string, new string) error {
	return filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println("----")
		fmt.Println(path)
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			panic(err)
		}
		data = bytes.ReplaceAll(data, []byte(old), []byte(new))
		fmt.Println(string(data))
		err = os.WriteFile(path, data, 0644)
		if err != nil {
			return err
		}
		return nil
	})
}
