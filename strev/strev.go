package strev

// Rev efficiently reverses a string
func Rev(str string) string {
	
	// parse the string out into runes, which play nicer than bytes for non-standard UTF-8 chars
	runes := []rune(str)
	
	// iterate over all runes, from inside out, swapping them in-place
	for idxHead, idxTail := 0, len(runes)-1; idxHead < len(runes)/2; idxHead, idxTail = idxHead+1, idxTail-1 {
		runes[idxHead], runes[idxTail] = runes[idxTail], runes[idxHead]
	}

	return string(runes)
	
}
