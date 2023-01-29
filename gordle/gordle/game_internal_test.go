package gordle

import (
	"errors"
	"testing"
)

// func TestGameAsk(t *testing.T) {
// 	tt := map[string]struct {
// 		input string
// 		want  []rune
// 	}{
// 		"5 characters in english": {
// 			input: "HELLO",
// 			want:  []rune("HELLO"),
// 		},
// 		"5 characters in arabic": {
// 			input: "مرحبا",
// 			want:  []rune("مرحبا"),
// 		},
// 		"5 characters in japanese": {
// 			input: "こんにちは",
// 			want:  []rune("こんにちは"),
// 		},
// 		"3 characters in japanese": {
// 			input: "こんに\nこんにちは",
// 			want:  []rune("こんにちは"),
// 		},
// 	}

// 	for name, tc := range tt {
// 		t.Run(name, func(t *testing.T) {

// 			g := New(strings.NewReader(tc.input), string(tc.want), 0)

// 			got := g.ask()
// 			if !slices.Equal(got, tc.want) {
// 				t.Errorf("readRunes() got = %v, want %v", string(got), string(tc.want))
// 			}
// 		})
// 	}
// }

func TestGameValidateGuess(t *testing.T) {
	tt := map[string]struct {
		word     []rune
		expected error
	}{
		"nominal": {
			word:     []rune("GUESS"),
			expected: nil,
		},
		"too long": {
			word:     []rune("POCKET"),
			expected: errInvalidWordLength,
		},
		"too short": {
			word:     []rune("DOG"),
			expected: errInvalidWordLength,
		},
		"empty": {
			word:     []rune(""),
			expected: errInvalidWordLength,
		},
		"nil": {
			word:     nil,
			expected: errInvalidWordLength,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			g := Game{}

			err := g.validateGuess(tc.word)
			if !errors.Is(err, tc.expected) {
				t.Errorf("%c, expected %q, got %q", tc.word, tc.expected, err)
			}
		})
	}
}

func TestComputeFeedback(t *testing.T) {
	tt := map[string]struct {
		guess            string
		solution         string
		expectedFeedback feedback
	}{
		"nominal": {
			guess:            "hello",
			solution:         "hello",
			expectedFeedback: feedback{correctPosition, correctPosition, correctPosition, correctPosition, correctPosition},
		},
		"double character": {
			guess:            "loyal",
			solution:         "hello",
			expectedFeedback: feedback{wrongPosition, wrongPosition, absentCharacter, absentCharacter, wrongPosition},
		},
		"double character with wrong answer": {
			guess:            "jello",
			solution:         "hello",
			expectedFeedback: feedback{absentCharacter, correctPosition, correctPosition, correctPosition, correctPosition},
		},
		"two identical, but not in the right position (from left to right)": {
			guess:            "hlleo",
			solution:         "hello",
			expectedFeedback: feedback{correctPosition, wrongPosition, correctPosition, wrongPosition, correctPosition},
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			fb := computeFeedback([]rune(tc.guess), []rune(tc.solution))
			if !tc.expectedFeedback.Equal(fb) {
				t.Errorf("guess: %q, got the wrong feedback, expected %v, got %v",
					tc.guess, tc.expectedFeedback, fb)
			}
		})
	}
}
