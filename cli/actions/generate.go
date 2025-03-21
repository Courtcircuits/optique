// command to bootstrap a new module
package actions

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/Courtcircuits/optique/cli/views"
)

func GenerateFromForm(name string) {
	view, err := views.LaunchGenForm()
	if err != nil {
		fmt.Println("Error launching form", err)
		fmt.Println("Error launching form")
		os.Exit(1)
	}
	if err := Generate(name, view.Type, view.URL); err != nil {
		fmt.Println("Error generating module 2", err)
		fmt.Println("Error launching form")
		os.Exit(1)
	}
}

func Generate(name string, rtype string, url string) error {
	if err := CreateModuleFolder(name); err != nil {
		return err
	}

	if err := CreateRepositoryManifestFile(name, rtype, url); err != nil {
		return err
	}

	if err := CreateRepositoryCode(name); err != nil {
		return err
	}

	return nil
}

func CreateModuleFolder(name string) error {
	err := os.Mkdir(name, os.ModePerm)
	if err != nil {
		return err
	}
	return os.Chdir(name)
}

type ModuleTemplate struct {
	NameCapitalized string
	Name            string `json:"name"`
	Type            string `json:"type"`
	URL             string `json:"url"`
}

const MODULE_TPL = `package {{.Name}}

// please implement the Repository interface

type {{.NameCapitalized }} struct {}

func New{{.NameCapitalized}}() (*{{.NameCapitalized}}, error) {
  panic("implement me")
}

func (m *{{.NameCapitalized}}) Bootstrap() error {
  panic("implement me")
}

func (m *{{.NameCapitalized}}) Stop() error {
  panic("implement me")
}
`

func CreateRepositoryCode(name string) error {
	template_content := ModuleTemplate{
		Name:            name,
		NameCapitalized: strings.Title(name),
	}

	tpl, err := template.New("module").Parse(MODULE_TPL)
	if err != nil {
		return err
	}
	f, err := os.Create(name + ".go")
	if err != nil {
		return err
	}

	return tpl.Execute(f, template_content)
}

func CreateRepositoryManifestFile(name string, rtype string, url string) error {
	template_content := ModuleTemplate{
		Name: name,
		Type: rtype,
		URL:  url,
	}

	template, err := json.Marshal(&template_content)
	if err != nil {
		return err
	}

	f, err := os.Create("config.json")
	defer f.Close()
	if err != nil {
		return err
	}

	_, err = f.Write(template)

	return err
}
