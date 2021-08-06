package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("Hola CLI")
	},
}

func init() {
	//fmt.Println("inside init")
	cobra.OnInitialize(initConfig)

}

func initConfig() {
	//fmt.Println("inside initConfig")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
