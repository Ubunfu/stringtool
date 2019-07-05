package hash

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"io"
	"log"
	"os"
	"strings"
)

// FileMD5 accepts two string file paths, to an input and output.
// The input file is read line by line, and a new file is written in the form
// ^{string} : {md5hash}$
func FileMD5(inFilePath string, outFilePath string) error {
	// Open the input file for reading, and the output file for writing
	inFile, err := os.Open(inFilePath)
	if err != nil {
		return err
	}
	defer inFile.Close()

	// NOTE: output file will be truncated if it already exists!!
	outFile, err := os.Create(outFilePath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	var stringBuilder strings.Builder // write buffer
	reader := bufio.NewReader(inFile) // read buffer

	// Read the input file a line at a time
	bytesRead, err := reader.ReadBytes('\n')
	if err != nil {
		return err
	}

	for len(bytesRead) > 0 {
		if err != nil {
			return err
		}

		bytesRead = bytesRead[:len(bytesRead)-1] // trim newline
		// log.Printf("Read: %q bytes of length %d.", bytesRead, len(bytesRead))
		hashedBytes16 := md5.Sum(bytesRead)
		hashedBytes := hashedBytes16[:] // converts [16]byte to []byte.  Omit the trailing \n
		stringBuilder.Write(bytesRead)
		stringBuilder.WriteString(" : ")
		stringBuilder.WriteString(hex.EncodeToString(hashedBytes))
		stringBuilder.WriteString("\n")

		bytesRead, err = reader.ReadBytes('\n')
	}

	// flush the buffer to the file
	_, err = io.WriteString(outFile, stringBuilder.String())
	if err != nil {
		log.Fatalln(err)
	}

	return nil
}
