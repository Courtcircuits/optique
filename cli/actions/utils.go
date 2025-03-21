package actions

import (
	"context"
	"os"
	"os/exec"
	"os/signal"

	"github.com/charmbracelet/huh/spinner"
)

func ExecWithLoading(label string, name string, commands ...string) error {
	cmd := context.Background()
	ctx, cancel := context.WithCancel(cmd)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	defer func() {
		signal.Stop(c)
		cancel()
	}()
	go func() {
		if err := spinner.New().Title(label).Context(ctx).Run(); err != nil {
		}
		select {
		case <-c:
			cancel()
			panic("Err")
		case <-ctx.Done():
		}
	}()
	if err := exec.CommandContext(ctx, name, commands...).Run(); err != nil {
		return err
	}
	return nil
}
