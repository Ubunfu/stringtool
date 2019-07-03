package cmd

import (
	"os"
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{
	Use: "stringtool",
	Short: "Do things with strings",
	Long: "Enumerate or reverse strings",
	Run: func(cmd *cobra.Command, args []string) {
		// Some stuff
		if len(args) == 0 {
			log.Fatalln("Choose a sub-command to execute!")
		}
	},
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
}