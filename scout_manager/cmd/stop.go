package cmd

import (
	"fmt"
	"github.com/kurolz/cobra"
	"os/exec"
	"runtime"
)

var Stop = &cobra.Command{
	Use:           "stop",
	Short:         "Stop Troop-Scout Service",
	RunE:          stop,
	SilenceUsage:  true,
	SilenceErrors: true,
}

func ExecStop() error {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "net", "stop", "troop-scout")
	} else {
		cmd = exec.Command("/bin/bash", "-c", "service troop-scout stop")
	}
	//cmd.Stdout = os.Stdout
	//cmd.Stderr = os.Stderr
	return cmd.Run()
}

func stop(c *cobra.Command, args []string) error {
	if err := ExecStop(); err != nil {
		return err
	}

	fmt.Print("successfully")
	return nil
}
