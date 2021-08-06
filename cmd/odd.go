package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var oddCommand = &cobra.Command{
	Use: "impar",
	Run: func(cmd *cobra.Command, args []string) {
		var oddSum int
		for _, val := range args {
			temp, _ := strconv.Atoi(val)
			if temp%2 != 0 {
				oddSum += temp
			}
		}
		fmt.Printf("La suma de impares de %s es %d\n", args, oddSum)
	},
}

func init() {
	addCommand.AddCommand(oddCommand)

}
