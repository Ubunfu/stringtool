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
	hashCmd.Flags().StringVarP(&algorithm, "algorithm", "a", "", "Algorithm to use for hashing the strings: [ md5 | sha1 | sha512 ]")
	hashCmd.Flags().StringVarP(&encoding, "encoding", "e", "hex", "Encoding to use for writing the hashed strings: [ hex | base64 ]")
	hashCmd.Flags().IntVarP(&rounds, "rounds", "r", 1, "Number of rounds of hashing to perform on the input strings")
	hashCmd.Flags().IntVarP(&flushInt, "flushInterval", "f", 10000, "Number of strings that will be hashed before flushing to disk")
	hashCmd.MarkFlagRequired("in-file")
	hashCmd.MarkFlagRequired("algorithm")
}

var inFilePath string
var outFilePath string
var algorithm string
var encoding string
var rounds int
var flushInt int

var hashCmd = &cobra.Command{
	Use: "hash",
	Aliases: []string{"h"},
	Short: "hash your strings",
	Long: "hash strings from a file and writes the {string, hash} pairs to a file",
	Run: func (cmd *cobra.Command, args []string) {
		err := hash.FileHash(inFilePath, outFilePath, flushInt, algorithm, rounds, encoding)
		if err != nil {
			log.Fatalln(err)
		}
	},
}