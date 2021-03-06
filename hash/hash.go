package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha512"
	"encoding/hex"
	"encoding/base64"
	"bufio"
	"errors"
	"io"
	"os"
	"strings"
)

// FileHash accepts two string file paths, to an input and output.
// The input file is read line by line, and a new file is written in the form
// ^{string} : {hash}$
func FileHash(inFilePath string, outFilePath string, flushInterval int, algorithm string, rounds int, encoding string) error {
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
	for hashCount := 0; len(bytesRead) > 0; hashCount++ {

		// Periodically flush the output buffer to disk
		if hashCount % flushInterval == 0 {
			_, err = io.WriteString(outFile, stringBuilder.String())
			if err != nil {
				return err
			}
			stringBuilder.Reset()
		}

		// Hash the bytes
		bytesRead = bytesRead[:len(bytesRead)-1] // trim "\n" byte
		hashedBytes, err := hashBytes(bytesRead, algorithm, rounds)
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
		return err
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
		encodedBytes := make([]byte, base64.StdEncoding.EncodedLen(len(bytes)))
		base64.StdEncoding.Encode(encodedBytes, bytes)
		return encodedBytes, nil
	default:
		err := errors.New("Invalid encoding format.  Supported formats are: hex, base64")
		return nil, err
	}
}

// hashBytes accepts a byte slice and a string describing the algorithm to use for hashing.
// It will return a byte slice containing the hashed bytes, and an error if pertinent
func hashBytes(bytes []byte, algorithm string, rounds int) ([]byte, error) {

	var hashedBytes []byte
	// decide which hashing algorithm to use
	switch algorithm {
	case "md5":
		hashedBytes = hashMd5(bytes, rounds)
	case "sha1":
		hashedBytes = hashSha1(bytes, rounds)
	case "sha512":
		hashedBytes = hashSha512(bytes, rounds)
	default:
		err := errors.New("Invalid hash algorithm.  Supported algorithms are: md5, sha1, sha512")
		return nil, err
	}
	return hashedBytes, nil
}

// hashMd5 hashes a byte slice using MD5 algorithm a given number of times
// It returns the resulting byte slice of hashed bytes
func hashMd5(bytes []byte, rounds int) []byte {
	hashedBytes := md5.Sum(bytes)
	for round := 1; round < rounds; round++ {
		hashedBytes = md5.Sum(hashedBytes[:])
	}
	return hashedBytes[:]
}

// hashSha1 hashes a byte slice using the SHA1 algorithm a given number of times
// it returns the resulting byte slice of hashed bytes
func hashSha1(bytes []byte, rounds int) []byte {
	hashedBytes := sha1.Sum(bytes)
	for round := 1; round < rounds; round++ {
		hashedBytes = sha1.Sum(hashedBytes[:])
	}
	return hashedBytes[:]
}

// hashSha512 hashes a byte slice using the SHA512 algorithm a given number of times
// it returns the resulting byte slice of hashed bytes
func hashSha512(bytes []byte, rounds int) []byte {
	hashedBytes := sha512.Sum512(bytes)
	for round := 1; round < rounds; round++ {
		hashedBytes = sha512.Sum512(hashedBytes[:])
	}
	return hashedBytes[:]
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
