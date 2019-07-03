package cmd

import (
	"github.com/spf13/cobra"
	"stringtool/strev"
	"log"
)

func init() {
	rootCmd.AddCommand(revCmd)
	revCmd.Flags().StringVarP(&str, "string", "s", "", "String to reverse")
}

var str string

var revCmd = &cobra.Command {
	Use: "reverse",
	Short: "Reverse a string",
	Long: "Reverse a string lexigraphically",
	Run: func(cmd *cobra.Command, args []string) {

		// Check for input
		if len(str) <= 0 {
			log.Fatalln("You must pass a string to reverse!")
		}

		log.Println(strev.Rev(str))
	},
}