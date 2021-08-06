package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var evenCommand = &cobra.Command{
	Use: "par",
	Run: func(cmd *cobra.Command, args []string) {
		var evenSum int
		for _, val := range args {
			temp, _ := strconv.Atoi(val)
			if temp%2 == 0 {
				evenSum += temp
			}
		}
		fmt.Printf("La suma de pares de %s es %d\n", args, evenSum)
	},
}

func init() {
	addCommand.AddCommand(evenCommand)

}
