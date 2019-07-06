package hash

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io"
	"log"
	"os"
	"strings"
)

// FileHash accepts two string file paths, to an input and output.
// The input file is read line by line, and a new file is written in the form
// ^{string} : {hash}$
func FileHash(inFilePath string, outFilePath string, algorithm string, encoding string) error {
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
	// (Prime the pump)
	bytesRead, err := reader.ReadBytes('\n') // Read until "\n" is found (inclusive)
	if err != nil {
		return err
	}

	// (Keep pumping 'til EOF)
	for len(bytesRead) > 0 {

		// Hash the bytes
		bytesRead = bytesRead[:len(bytesRead)-1] // trim "\n" byte
		hashedBytes, err := hashBytes(bytesRead, algorithm)
		if err != nil {
			return err
		}

		// encode them for readability
		encodedHash, err := encode(hashedBytes, encoding)
		if err != nil {
			return err
		}

		// Write to the output buffer
		stringBuilder, err = bufferHash(bytesRead, encodedHash, stringBuilder)
		if err != nil {
			return err
		}

		// Read the next line from the file
		bytesRead, err = reader.ReadBytes('\n')
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return err
		}
	}

	// flush the output buffer to file
	_, err = io.WriteString(outFile, stringBuilder.String())
	if err != nil {
		log.Fatalln(err)
	}
	return nil
}

// encode will apply an encoding format to data to make it more readable
func encode(bytes []byte, encoding string) ([]byte, error) {

	switch encoding {
	case "hex":
		encodedBytes := make([]byte, hex.EncodedLen(len(bytes)))
		hex.Encode(encodedBytes, bytes)
		return encodedBytes, nil
	case "base64":
		err := errors.New("base64 encoding is not implemented yet")
		return nil, err
	default:
		err := errors.New("Invalid encoding format.  Supported formats are: hex, base64")
		return nil, err
	}

}

// hashBytes accepts a byte slice and a string describing the algorithm to use for hashing.
// It will return a byte slice containing the hashed bytes, and an error if pertinent
func hashBytes(bytes []byte, algorithm string) ([]byte, error) {

	var hashedBytes []byte
	// decide which hashing algorithm to use
	switch algorithm {
	case "md5":
		hashedBytes16 := md5.Sum(bytes)
		hashedBytes = hashedBytes16[:]
	case "sha1":
		err := errors.New("sha1 hashing is not implemented yet")
		return nil, err
	case "sha512":
		err := errors.New("sha512 hashing is not implemented yet")
		return nil, err
	default:
		err := errors.New("Invalid hash algorithm.  Supported algorithms are: md5, sha1, sha512")
		return nil, err
	}

	return hashedBytes, nil
}

// bufferHash writes the rainbow table data to the output buffer in the proper format
// It will return nothing but an error if pertinent
func bufferHash(plainBytes []byte, hashedBytes []byte, buffer strings.Builder) (strings.Builder, error) {
	_, err := buffer.Write(plainBytes)
	if err != nil {
		return buffer, err
	}
	_, err = buffer.WriteString(" : ")
	if err != nil {
		return buffer, err
	}
	_, err = buffer.Write(hashedBytes)
	if err != nil {
		return buffer, err
	}
	_, err = buffer.WriteString("\n")
	if err != nil {
		return buffer, err
	}
	return buffer, nil
}
