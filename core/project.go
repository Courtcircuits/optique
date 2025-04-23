package core


import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// find the location of optique.json since it is the root of the project
func FindOptiqueJson() (string, error) {
	// first check in current level
	current_path, err := os.Getwd()

	matches, err := filepath.Glob(fmt.Sprintf("%s/%s", current_path, PROJECT_MANIFEST))
	if matches != nil && len(matches) > 0 {
		fmt.Println(matches)
		return matches[0], nil
	}
	MAX_DEPTH := 200
	for range MAX_DEPTH {
		os.Chdir("..")
		current_path, err = os.Getwd()
		fmt.Println(current_path)
		if err != nil {
			return "", err
		}
		if current_path == "/" {
			return "", ERR_NO_OPTIQUE_JSON
		}
		matches, err = filepath.Glob(fmt.Sprintf("%s/%s", current_path, PROJECT_MANIFEST))
		if err != nil {
			return "", err
		}
		if len(matches) > 0 {
			return matches[0], nil
		}
	}
	return "", ERR_NO_OPTIQUE_JSON
}

var ERR_NO_OPTIQUE_JSON = errors.New("no optique.json found")
