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
	hashCmd.Flags().StringVarP(&algorithm, "algorithm", "a", "", "Algorithm to use for hashing the strings")
	hashCmd.Flags().StringVarP(&encoding, "encoding", "e", "hex", "Encoding to use for writing the hashed strings")
	hashCmd.MarkFlagRequired("in-file")
	hashCmd.MarkFlagRequired("algorithm")
}

var inFilePath string
var outFilePath string
var algorithm string
var encoding string

var hashCmd = &cobra.Command{
	Use: "hash",
	Aliases: []string{"h"},
	Short: "hash your strings",
	Long: "hash strings from a file and writes the {string, hash} pairs to a file",
	Run: func (cmd *cobra.Command, args []string) {
		err := hash.FileHash(inFilePath, outFilePath, algorithm, encoding)
		if err != nil {
			log.Fatalln(err)
		}
	},
}