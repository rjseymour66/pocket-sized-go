package main

import "testing"

func ExampleMain() {
	main()
	// Output:
	// Hello world
}

func TestGreet(t *testing.T) {
	type scenario struct {
		lang             language
		expectedGreeting string
	}

	var tests = map[string]scenario{
		"English": {
			lang:             "en",
			expectedGreeting: "Hello world",
		},
		"French": {
			lang:             "fr",
			expectedGreeting: "Bonjour le monde",
		},
		"Akkadian, not supported": {
			lang:             "akk",
			expectedGreeting: `unsupported language: "akk"`,
		},
		"Greek": {
			lang:             "el",
			expectedGreeting: "Χαίρετε Κόσμε",
		},
		"Hebrew": {
			lang:             "he",
			expectedGreeting: "שלום עולם",
		},
		"Urdu": {
			lang:             "ur",
			expectedGreeting: "ہیلو دنیا",
		},
		"Vietnamese": {
			lang:             "vi",
			expectedGreeting: "Xin chào Thế Giới",
		},
		"Spanish": {
			lang:             "es",
			expectedGreeting: "Hola",
		},
		"Empty": {
			lang:             "",
			expectedGreeting: `unsupported language: ""`,
		},
	}

	// range over all scenarios
	for scenarioName, tc := range tests {
		t.Run(scenarioName, func(t *testing.T) {
			greeting := greet(tc.lang)

			if greeting != tc.expectedGreeting {
				t.Errorf("expected: %q, got: %q", tc.expectedGreeting, greeting)
			}
		})
	}
}
