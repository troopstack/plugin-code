package cmd

import (
	"../utils"
	"fmt"
	"github.com/kurolz/cobra"
	"os/exec"
	"runtime"
)

var Start = &cobra.Command{
	Use:           "start",
	Short:         "Start Troop-Scout Service",
	RunE:          start,
	SilenceUsage:  true,
	SilenceErrors: true,
}

func ExecStart() error {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "net", "start", "troop-scout")
	} else {
		cmd = exec.Command("/bin/bash", "-c", "service troop-scout start")
	}
	//cmd.Stdout = os.Stdout
	//cmd.Stderr = os.Stderr
	return cmd.Run()
}

func start(c *cobra.Command, args []string) error {
	if utils.IsRunning() {
		fmt.Print("troop-scout service is already running")
		return nil
	}
	if err := ExecStart(); err != nil {
		return err
	}

	fmt.Print("successfully")
	return nil
}
