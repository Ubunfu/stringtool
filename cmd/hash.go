package cmd

import (
	"github.com/spf13/cobra"
	"github.com/Ubunfu/stringtool/hash"
	"log"
)

func init() {
	rootCmd.AddCommand(hashCmd)
	hashCmd.Flags().StringVarP(&inFilePath, "in-file", "i", "", "File with strings to hash")
	hashCmd.Flags().StringVarP(&outFilePath, "out-file", "o", "hashes.out", "File to write the strings and hashes")
	hashCmd.MarkFlagRequired("in-file")
}

var inFilePath string
var outFilePath string

var hashCmd = &cobra.Command{
	Use: "hash",
	Aliases: []string{"h"},
	Short: "hash your strings",
	Long: "hash strings from a file and writes the {string, hash} pairs to a file",
	Run: func (cmd *cobra.Command, args []string) {
		err := hash.FileMD5(inFilePath, outFilePath)
		if err != nil {
			log.Fatalln(err)
		}
	},
}