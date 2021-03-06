package enumerate

import (
	"errors"
	"io"
	"log"
	"os"
	"strings"
)

var maxLen int
var minLen int
var begin string
var end string
var outPath string

// Enumerate brute-forces strings of given lengths
func Enumerate(minLen int, maxLen int, begin string, end string, outPath string, flushInterval int) error {

	// Input validation
	if minLen > maxLen {
		return errors.New("Minimum length must not exceed the maximum length!")
	}

	// Default and optionally override the starting point
	if len(begin) == 0 {
		for index := 0; index < minLen; index++ {
			begin = begin + string("0")
		}
	} else if len(begin) > maxLen {
		return errors.New("Starting point for enumeration is longer than the max length!")
	}

	// Default and optionally override the end point
	if len(end) == 0 {
		for index := 0; index < maxLen; index++ {
			end = end + string("z")
		}
	} else if len(end) < minLen {
		return errors.New("Stopping point for enumeration is shorter than the min length!")
	}

	if begin > end {
		return errors.New("Start string must appear prior to the end staring when ordered lexigraphically.")
	}

	// Convert the starting string into runes
	runes := []rune(begin)

	// try to open the file for write and append
	file, err := os.Create(outPath)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	// Create a buffer to write strings to
	var stringBuilder strings.Builder

	// loop forever until stopped by the embedded guard clause
	for enumCount := 0; true; enumCount++ {

		// Periodically flush output buffer to disk
		if enumCount % flushInterval == 0 {
			// flush the buffer to the file
			_, err = io.WriteString(file, stringBuilder.String())
			if err != nil {
				log.Fatalln(err)
			}
			file.Sync()
			stringBuilder.Reset()
		}

		// write strings to the buffer to save time
		stringBuilder.WriteString(string(runes))
		stringBuilder.WriteString("\n")

		// Guard clause to exit gracefully if done enumerating
		if string(runes) == end {
			break
		}
		runes = increment(runes)
	}

	// flush the buffer to the file
	_, err = io.WriteString(file, stringBuilder.String())
	if err != nil {
		log.Fatalln(err)
	}

	file.Sync()

	return nil
}

// increment will increment the last rune in the array and return the whole array
func increment(runes []rune) []rune {
	lastRune := runes[len(runes)-1]
	newRuneInt := int(lastRune) + 1

	// If the last rune is betwen 9 and A, skip ahead
	// If the last rune is betwen Z and a, skip ahead
	if newRuneInt == 58 {
		newRuneInt = newRuneInt + 7
	} else if newRuneInt == 91 {
		newRuneInt = newRuneInt + 6
	}

	// New Rune is past z by 1
	if newRuneInt == 123 {
		if stringIsMaxed(runes) { // Is current length string fully incremented?
			return extendRunes(runes)
		}

		// The current length string is not maxed, so roll the runes by incrementing the substring [0:n-1]
		// E.g. if the current string is aaz:
		// (1) Call increment(aa) => ab
		// (2) Re-connect the full string: ab + 0 => ab0
		// We have effectively turned aaz => ab0
		newRuneInt = 48
		rolledRunes := increment(runes[0 : len(runes)-1])
		runes = append(rolledRunes, rune(newRuneInt))

	} else {
		newRune := rune(newRuneInt)
		runes[len(runes)-1] = newRune
	}
	return runes
}

// extendRunes returns a new slice of runes, initialized with 0's, of length one greater
// 		than the slice passed into it
func extendRunes(runes []rune) []rune {
	extendedRunes := []rune{}
	for len(extendedRunes) <= len(runes) {
		extendedRunes = append(extendedRunes, rune(48))
	}
	return extendedRunes
}

// stringIsMaxed checks whether the current length string can be incremented any further
// Returns true if all characters are the highest acceptable value (e.g. zzz)
// Otherwise returns false (e.g. aaz)
func stringIsMaxed(runes []rune) bool {

	// If there are any non-"z" runes in the slice, return false
	for _, b := range runes {
		if b != rune(122) {
			return false
		}
	}
	return true
}
