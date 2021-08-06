package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var addCommand = &cobra.Command{
	Use: "add",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:
	
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fstatus, _ := cmd.Flags().GetBool("float")

		if fstatus {
			addFloat(args)
		} else {
			addInt(args)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCommand)
	addCommand.Flags().BoolP("float", "f", false, "Add floating numbers")
}

func addFloat(args []string) {
	var sum float64

	for _, val := range args {
		temp, err := strconv.ParseFloat(val, 64)
		if err != nil {
			fmt.Println(err.Error())
		}
		sum += temp
	}
	fmt.Printf("La suma de %s es %f\n", args, sum)
}

func addInt(args []string) {
	var sum int

	for _, val := range args {
		temp, err := strconv.Atoi(val)
		if err != nil {
			fmt.Println(err.Error())
		}
		sum += temp
	}
	fmt.Printf("La suma de %s es %d\n", args, sum)
}
