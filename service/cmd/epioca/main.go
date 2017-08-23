package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	version     string
	configFile  string
	versionFlag bool
)

func main() {
	versionString := "ePioca v" + version
	cobra.OnInitialize(func() {
		if versionFlag {
			fmt.Println(versionString)
			os.Exit(0)
		}
	})

	var RootCmd = &cobra.Command{
		Use:   "epioca",
		Short: "ePioca bidding platform",
		Long:  versionString + ``,
		Run:   RunServer,
	}
	RootCmd.Flags().StringVarP(&configFile, "config", "c", "", "Source of a configuration file")
	RootCmd.Flags().BoolVarP(&versionFlag, "version", "v", false, "Print application version")

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
