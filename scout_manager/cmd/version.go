package cmd

import (
	"../utils"
	"fmt"
	"github.com/kurolz/cobra"
)

var Version = &cobra.Command{
	Use:           "version",
	Short:         "Show Plugin Version",
	RunE:          version,
	SilenceUsage:  true,
	SilenceErrors: true,
}

func version(c *cobra.Command, args []string) error {
	fmt.Printf(utils.VERSION)
	return nil
}
