package main

import (
	"./cmd"
	"fmt"
	"github.com/kurolz/cobra"
	"os"
)

var RootCmd = &cobra.Command{
	Use:  "scout_manager",
	RunE: startListen,
}

var port int

func init() {
	RootCmd.AddCommand(cmd.Version)
	RootCmd.AddCommand(cmd.Status)
	RootCmd.AddCommand(cmd.Update)
	RootCmd.AddCommand(cmd.ConfigUpdate)

	RootCmd.PersistentFlags().IntVarP(&port, "port", "p", 8577,
		"http listen port")
}

func startListen(c *cobra.Command, args []string) error {
	return c.Usage()
}

func main() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
