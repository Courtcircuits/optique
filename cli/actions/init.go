package actions

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/Courtcircuits/optique/cli/views"
)

const URL = "https://github.com/Courtcircuits/optique/cli"

type Initialization struct {
	URL     string
	Name    string
	Version string
}

var DefaultInitialization = &Initialization{
	Name:    "optique",
	URL:     "https://github.com/baptistebronsin/javoue",
	Version: "latest",
}

func NewInitialization(name string) Initialization {
	DefaultInitialization.Name = name
	StartForm(name)
	return *DefaultInitialization
}

func Initialize(generation Initialization) {
	err := createProjectFolder(generation.Name)
	if err != nil {
		fmt.Println("Error creating project folder:", err)
		os.Exit(1)
	}
	err = cloneTemplate("https://github.com/Courtcircuits/optique", generation.Name)
	if err != nil {
		fmt.Println("Error cloning template:", err)
		os.Exit(1)
	}
	err = setupGoModule(&generation)
	if err != nil {
		fmt.Println("Error setting up go module:", err)
		os.Exit(1)
	}
	err = goBack()
	if err != nil {
		fmt.Println("Error going back:", err)
		os.Exit(1)
	}
}

func createProjectFolder(name string) error {
	err := os.Mkdir(name, 0755)
	if err != nil {
		return err
	}
	return nil
}

func goBack() error {
	return os.Chdir("..")
}

func cloneTemplate(url string, name string) error {
	cmd := exec.Command("git", "clone", url, name)
	views.Load(cmd, "Cloning template")

	// go to project folder
	err := os.Chdir(name)
	if err != nil {
		return err
	}

	current_dir, err := os.Getwd()
	if err != nil {
		return err
	}

	entries, err := os.ReadDir(current_dir)

	if err != nil {
		return err
	}

	folders_to_delete := []string{}
	files_to_delete := []string{}
	for _, entry := range entries {
		if entry.IsDir() {
			if entry.Name() != "template" {
				folders_to_delete = append(folders_to_delete, entry.Name())
			}
		} else {
			files_to_delete = append(files_to_delete, entry.Name())
		}
	}

	for _, entry := range folders_to_delete {
		err = os.RemoveAll(entry)
		if err != nil {
			return err
		}
	}
	for _, entry := range files_to_delete {
		err = os.Remove(entry)
		if err != nil {
			return err
		}
	}

	// go to template folder
	err = os.Chdir("template")
	if err != nil {
		return err
	}

	entries, err = os.ReadDir(".")
	for _, entry := range entries {
		val, err := exec.Command("mv", entry.Name(), current_dir).CombinedOutput()
		if err != nil {
			return err
		}
		fmt.Println(string(val))
	}
	// move all to parent folder
	err = goBack()
	if err != nil {
		return err
	}

	// remove template folder
	err = os.RemoveAll("template")
	return nil
}

func setupGoModule(config *Initialization) error {
	cmd := exec.Command("go", "mod", "init", config.URL)
	views.Load(cmd, "Initializing go module")
	cmd = exec.Command("gopls", "imports", "-w", "./main.go")
	views.Load(cmd, "Cleaning up imports")

	cmd = exec.Command("go", "mod", "tidy")
	views.Load(cmd, "Installing dependencies")

	return nil
}
