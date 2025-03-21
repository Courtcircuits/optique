package actions

import (
	"encoding/json"
	"os"
)

func AddModule(repo_url string, path string) {
	data := SetUpSparseModule(repo_url, path)
	CleanUpSparseModule()
	goBack()
	goBack()
	MoveModule(".optique/tmp", data.Name)
	CleanUpOptique()
}

func SetUpSparseModule(repo_url string, path string) *ModuleTemplate {
	//create temp folder
	os.Mkdir(".optique", os.ModePerm)
	os.Chdir(".optique")
	os.Mkdir("tmp", os.ModePerm)
	os.Chdir("tmp")
	ExecWithLoading("Initializing module", "git", "init")
	ExecWithLoading("Creating sparse module", "git", "remote", "add", "origin", repo_url)
	ExecWithLoading("Sparse-checking module", "git", "sparse-checkout", "init", "--cone")
	ExecWithLoading("Setting up module", "git", "sparse-checkout", "set", path)
	ExecWithLoading("Pulling module", "git", "pull", "origin", "main")
	return ParseModuleData()
}

func CleanUpSparseModule() {
	os.RemoveAll(".git")
}

func ParseModuleData() *ModuleTemplate {
	// read config.json
	fd, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	var data ModuleTemplate
	err = json.NewDecoder(fd).Decode(&data)
	if err != nil {
		panic(err)
	}

	return &data
}

func MoveModule(path string, destination string) {
	ExecWithLoading("Moving module", "mv", path, ".")
	ExecWithLoading("Moving module", "mv", "tmp",destination)
}

func CleanUpOptique() {
	os.RemoveAll(".optique")
}
