package test

import (
	"fmt"
	"os"
)

// buils a filesystem based on an interface
// e.g the following interface :
// {
// "config": {"config.go": "package config"},
// "optique.json" : "optique.json"
// }
// will induce the following actions:
// - create folder config
// - create file config.go
// - create file optique.json
func BuildFilesystem(filesytem any) error {
	mobj, ok := filesytem.(map[string]any)
	if !ok {
		return nil
	}
	for key, value := range mobj {
		switch value.(type) {
		case string:
			err := os.WriteFile(key, []byte(value.(string)), 0644)
			if err != nil {
				return err
			}
		case map[string]any:
			err := os.MkdirAll(key, 0755)
			if err != nil {
				return err
			}
			os.Chdir(key)
			err = BuildFilesystem(value)
			os.Chdir("..")
			if err != nil {
				return err
			}
		default:
			return fmt.Errorf("invalid filesytem")
		}
	}
	return nil
}

type FSBuilder struct {
	root_path string
	first_path string
}

func NewFSBuilder(root_path string) *FSBuilder {
	current_path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return &FSBuilder{root_path: root_path, first_path: current_path}
}

func (b *FSBuilder) Build(filesytem any) error {
	os.MkdirAll(b.root_path, 0755)
	os.Chdir(b.root_path)
	if err := BuildFilesystem(filesytem); err != nil {
		os.Chdir(b.first_path)
		return err
	}
	os.Chdir(b.first_path)
	return nil
}

func (b *FSBuilder) GoTo(path string) {
	os.Chdir(b.root_path)
	os.Chdir(path)
	current_path, _ := os.Getwd()
	fmt.Println(current_path)
}

func (b *FSBuilder) Clean() {
	os.Chdir(b.first_path)
	os.RemoveAll(b.root_path)
}
