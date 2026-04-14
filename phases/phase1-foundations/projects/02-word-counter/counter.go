package wordcounter

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

// WordFreq holds a word and its count.
type WordFreq struct {
	Word  string
	Count int
}

// CountWords reads all text from r, tokenizes by whitespace,
// lowercases each word, and returns a frequency map.
// TODO: use bufio.Scanner with ScanWords, strings.ToLower, map[string]int
func CountWords(r io.Reader) (map[string]int, error) {
	return nil, fmt.Errorf("not implemented")
}

// TopN returns the top n words by frequency, sorted desc by count then asc by word.
// If n <= 0 or n > len(freq), returns all words sorted.
// TODO: build []WordFreq from map, sort, slice to n
func TopN(freq map[string]int, n int) []WordFreq {
	return nil
}

// keep imports
var _, _, _ = bufio.NewScanner, sort.Slice, strings.ToLower
