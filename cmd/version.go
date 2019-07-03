package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use: "version",
	Short: "show version",
	Long: "Display the current version number of the stringtool CLI tool",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("stringtool - v1.0")
		fmt.Println("Ryan Allen (rallen3882@gmail.com)")
		fmt.Println("June 21, 2019")
	},
}