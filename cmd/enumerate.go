package cmd

import (
	"github.com/spf13/cobra"
	"github.com/Ubunfu/stringtool/enumerate"
	"log"
)

func init() {
	rootCmd.AddCommand(enumCmd)
	enumCmd.Flags().IntVarP(&minLen, "min-length", "n", 1, "Minimum string length")
	enumCmd.Flags().IntVarP(&maxLen, "max-length", "x", 3, "Maximum string length")
	enumCmd.Flags().StringVarP(&begin, "begin", "b", "", "Starting point for enumeration. E.g. Ry4")
	enumCmd.Flags().StringVarP(&end, "end", "e", "", "Ending point for enumeration. E.g. ccc")
	enumCmd.Flags().StringVarP(&output, "output", "o", "strings.out", "Output file for enumerated strings")
	enumCmd.Flags().IntVarP(&flushInterval, "flushInterval", "f", 10000, "Number of strings that will be enumerated before flushing to disk")
}

// Flag vars
var minLen int
var maxLen int
var begin string
var end string
var output string
var flushInterval int

var enumCmd = &cobra.Command{
	Use: "enumerate",
	Aliases: []string{"e"},
	Short: "emumarate strings",
	Long: `Enumerate all alpha-numeric strings of the given length parameters. 
		Brute-force style.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Enumerate all the strings
		err := enumerate.Enumerate(minLen, maxLen, begin, end, output, flushInterval)
		if err != nil {
			log.Fatalln(err)
		}
	},
}