package cmd

import (
	"github.com/spf13/cobra"
	"stringtool/enumerate"
)

func init() {
	rootCmd.AddCommand(enumCmd)
	enumCmd.Flags().IntVarP(&minLen, "min-length", "n", 1, "Minimum string length")
	enumCmd.Flags().IntVarP(&maxLen, "max-length", "x", 3, "Maximum string length")
	enumCmd.Flags().StringVarP(&output, "output", "o", "strings.out", "Output file for enumerated strings")
}

// Flag vars
var minLen int
var maxLen int
var output string

var enumCmd = &cobra.Command{
	Use: "enumerate",
	Short: "emumarate strings",
	Long: `Enumerate all alpha-numeric strings of the given length parameters. 
		Brute-force style.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Enumerate all the strings
		enumerate.Enumerate(minLen, maxLen, output)
	},
}