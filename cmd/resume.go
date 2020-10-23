package cmd

import (
	cobra "github.com/bobbaicloudwithpants/bobra"
)

var resume = &cobra.Command{
	Use: "resume",			// Use 指定了这个命令的名字
	Short: "resume is a simple self-introduction cli program",	// Short 是对于该命令的简短介绍
	Long: "resume makes you organize your personal resume properly, and display in a user-friendly and cleary way.",	// Long 是命令的比较完整的介绍
	Run: func(c *cobra.Command, args []string) {
		c.Usage()
	},
}


func Execute() {
	//if err := rootCmd.Execute(); err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	resume.Execute()
}
