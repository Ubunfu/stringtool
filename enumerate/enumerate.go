package enumerate

import (
	"io"
	"log"
	"math"
	"os"
	"strings"
)

var maxLen int
var minLen int
var outPath string

// Enumerate brute-forces strings of given lengths
func Enumerate(minLen int, maxLen int, outPath string) {

	// Build a starting point
	start := ""
	for index := 0; index < minLen; index++ {
		start = start + string("0")
	}

	// Convert the starting string into runes
	runes := []rune(start)

	increments := calcIncrements(minLen, maxLen, 62)
	log.Printf("Enumerating %d strings to %s ...", increments, outPath)

	// try to open the file for write and append
	file, err := os.OpenFile(outPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	// Increment until max length exceeded
	for idx := 0; idx < increments; idx++ {
		runes = increment(runes)

		var stringBuilder strings.Builder
		stringBuilder.WriteString(string(runes))
		stringBuilder.WriteString("\n")

		_, err = io.WriteString(file, stringBuilder.String())
		if err != nil {
			log.Fatalln(err)
		}
	}
	file.Sync()
	log.Println("Done.")
}

// calcIncrements pre-calculates the number of times that we will need to increment
// our string in order to brute-force all of the possible permutations
func calcIncrements(minLen int, maxLen int, dictSize int) int {
	totalIncrements := 0
	for idx := minLen; idx <= maxLen; idx++ {
		totalIncrements = totalIncrements + int(math.Pow(float64(dictSize), float64(idx)))
	}
	totalIncrements--
	return totalIncrements
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
		// Check if there are any more strings of this length to enumerate
		if stringIsMaxed(runes) {
			return extendRunes(runes)
		}

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

// stringIsMaxed checks to see if there are any more strings that can be enumerated
func stringIsMaxed(runes []rune) bool {
	for _, b := range runes {
		if b != rune(122) {
			return false
		}
	}
	return true
}
