package cmd

import (
	"errors"
	"fmt"
	"github.com/kurolz/cobra"
	"github.com/toolkits/file"
	"os/exec"
	"runtime"
)

var ConfigUpdate = &cobra.Command{
	Use:           "config.update",
	Short:         "config update",
	RunE:          configUpdate,
	Args:          cobra.MinimumNArgs(1),
	SilenceUsage:  true,
	SilenceErrors: true,
}

func configUpdate(c *cobra.Command, args []string) error {
	if runtime.GOOS == "linux" {
		newConfigFile := args[0]
		configFile := "/usr/local/troop-scout/conf/config.ini"

		if !file.IsExist(newConfigFile) {
			errLog := fmt.Sprintln("no new configuration file was detected on the server:", newConfigFile)
			return errors.New(errLog)
		}

		cmdCp := exec.Command("cp", newConfigFile, configFile)
		errCp := cmdCp.Run()
		if errCp != nil {
			fmt.Printf("Cp configuration file failed. order: %s", cmdCp.Args)
			return errCp
		}
		cmd3 := exec.Command("systemctl", "restart", "troop-scoutd.service")
		err3 := cmd3.Run()
		if err3 != nil {
			fmt.Printf("systemctl restart troop-scoutd.service failed. ")
			return err3
		}
		fmt.Print("successfully")
	} else {
		return errors.New("only support linux")
	}
	return nil
}
