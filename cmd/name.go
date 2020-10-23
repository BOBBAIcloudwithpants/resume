package cmd

import (
	"fmt"
	cobra "github.com/bobbaicloudwithpants/bobra"
)
var name = &cobra.Command{
	Use : "name",
	Short: "name screens user's name",
	Long: "name displays the user's name put into the memory in advance.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("bjd")
	},
}

func init() {
	resume.AddCommand(name)
}
