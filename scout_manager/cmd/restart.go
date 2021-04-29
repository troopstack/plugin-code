package cmd

import (
	"fmt"
	"github.com/kurolz/cobra"
	"os"
	"os/exec"
	"runtime"
)

var Restart = &cobra.Command{
	Use:           "restart",
	Short:         "Restart Troop-Scout Service",
	RunE:          restart,
	SilenceUsage:  true,
	SilenceErrors: true,
}

func ExecRestart() error {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("net", "stop", "troop-scout")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
		}

		cmdStart := exec.Command("net", "start", "troop-scout")
		cmdStart.Stdout = os.Stdout
		cmdStart.Stderr = os.Stderr
		return cmdStart.Run()

	} else {
		cmd = exec.Command("/bin/bash", "-c", "service troop-scout restart")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	}
}

func restart(c *cobra.Command, args []string) error {
	if err := ExecRestart(); err != nil {
		return err
	}

	fmt.Print("successfully")
	return nil
}
