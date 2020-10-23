package cmd

import (
	"fmt"
	cobra "github.com/bobbaicloudwithpants/bobra"
)

var edu = &cobra.Command{
	Use: "edu",
	Short: "edu stores and displays you educational background",
	Long: "edu reads the data put into the program in advance, and dynamically chooses which item to show based on the given parameters.",
	Run: func(cmd *cobra.Command, args []string) {
		if ok, _ := cmd.Flags().GetBool("college");ok {
			fmt.Println("SYSU")
		} else if ok, _ := cmd.Flags().GetBool("middle");ok{
			fmt.Println("THSchool")
		} else if ok, _ := cmd.Flags().GetBool("primary");ok{
			fmt.Println("TS2FX")
		}
	},
}

func init(){
	edu.Flags().BoolP("college", "c", false, "whether show college")
	edu.Flags().BoolP("middle", "m", false, "whether show middle school")
	edu.Flags().BoolP("primary", "p", false, "whether show primary")
	resume.AddCommand(edu)
}

