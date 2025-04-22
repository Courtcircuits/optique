package test

import (
	"os"
	"testing"

	"github.com/Courtcircuits/optique/cli/actions"
)

func TestFindRootAtSameLevel(t *testing.T) {
	root_path := "test_fs"
	current_path, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	expecting := current_path + "/" + root_path + "/optique.json"
	fs := map[string]any{
		"app": map[string]any{
			"config.go": "package config",
		},
		"optique.json": "test",
	}
	builder := NewFSBuilder(root_path)
	err = builder.Build(fs)
	if err != nil {
		t.Fatal(err)
	}
	builder.GoTo("")
	defer builder.Clean()
	root, err := actions.FindOptiqueJson()
	if err != nil {
		builder.Clean()
		t.Fatal(err)
	}
	if root != expecting {
		builder.Clean()
		t.Fatalf("expected %s, got %s", expecting, root)
	}
}

func TestFindRootAtParentLevel(t *testing.T) {
	root_path := "test_fs"
	current_path, err := os.Getwd()
	fs := map[string]any{
		"app": map[string]any{
			"domain": map[string]any{
				"config.go": "package config",
			},
		},
		"optique.json": "test",
	}
	builder := NewFSBuilder(root_path)
	err = builder.Build(fs)
	if err != nil {
		t.Fatal(err)
	}
	builder.GoTo("app/domain")
	defer builder.Clean()
	root, err := actions.FindOptiqueJson()
	if err != nil {
		builder.Clean()
		t.Fatal(err)
	}
	expecting := current_path + "/test_fs/optique.json"
	if root != expecting {
		builder.Clean()
		t.Fatalf("expected %s, got %s", expecting, root)
	}
}
