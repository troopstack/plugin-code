package cmd

import (
	"../utils"
	"fmt"
	"github.com/kurolz/cobra"
)

var Status = &cobra.Command{
	Use:           "status",
	Short:         "Get Service Status",
	RunE:          status,
	SilenceUsage:  true,
	SilenceErrors: true,
}

func status(c *cobra.Command, args []string) error {
	if utils.IsRunning() {
		fmt.Print("up")
	} else {
		fmt.Print("down")
	}
	return nil
}
