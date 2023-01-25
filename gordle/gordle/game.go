package gordle

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/exp/slices"
)

// Game holds all the information we need to play a game of gordle.
type Game struct {
	reader      *bufio.Reader
	solution    []rune
	maxAttempts int
}

// New returns a Game, which can be used to Play!
func New(playerInput io.Reader, solution string, maxAttempts int) *Game {
	g := &Game{
		reader:      bufio.NewReader(playerInput),
		solution:    splitToUppercaseCharacter(solution),
		maxAttempts: maxAttempts,
	}
	return g
}

// Play runs the game.
func (g *Game) Play() {
	fmt.Println("Welcome to Gordle!")

	for currentAttempt := 1; currentAttempt <= g.maxAttempts; currentAttempt++ {
		// ask for a valid word
		guess := g.ask()

		if slices.Equal(guess, g.solution) {
			fmt.Printf("🎉 You won! You found it in %d attempts(s)! The word was: %s\n", currentAttempt, string(g.solution))
			return
		}
	}

	fmt.Printf("😞 You lost! The solution was: %s.\n", string(g.solution))
}

const wordLength = 5

// ask reads input until a valid suggestion is made (and returned)
func (g *Game) ask() []rune {
	fmt.Printf("Enter a %d-character guess:\n", len(g.solution))

	for {
		playerInput, _, err := g.reader.ReadLine()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Gordle failed to read your guess: %s\n", err.Error())
			continue
		}

		guess := splitToUppercaseCharacter(string(playerInput))

		err = g.validateGuess(guess)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Your attempt is invalid with Gordle's solution: %s.\n", err.Error())
		} else {
			return guess
		}
	}
}

// errInvalidWordLength is returned when the guess has the wrong number of characters
var errInvalidWordLength = fmt.Errorf("invalid guess. The word " +
	"does not have the correct number of characters")

func (g *Game) validateGuess(guess []rune) error {
	if len(guess) != wordLength {
		return fmt.Errorf("expected %d, got %d, %w", len(g.solution), len(guess), errInvalidWordLength)
	}
	return nil
}

// splitToUppercaseCharacter is a naive implementation that turns a
// string into a list of characters
func splitToUppercaseCharacter(input string) []rune {
	return []rune(strings.ToUpper(input))
}
