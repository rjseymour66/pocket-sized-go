package gordle

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// ErrCorpusIsEmpty is returned when the copus file does not contain data
const ErrCorpusIsEmpty = corpusError("corpus is empty")

// ReadCorpus reads the file located at the given path
// and returns a list of words.
func ReadCorpus(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open %q for reading: %w", path, err)
	}

	if len(data) == 0 {
		return nil, ErrCorpusIsEmpty
	}

	// we expect the corups to be a line- or space-separated list of words
	words := strings.Fields(string(data))

	return words, nil
}

// PickWord randomly selects a word from a slice of strings
func PickWord(corpus []string) string {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(corpus))
	return corpus[index]
}
