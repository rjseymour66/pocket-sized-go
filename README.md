# Pocket-sized Go

## Testing

You can name files `hello_world_internal_test.go`. The `internal` signifies that you are testing unexposed/private methods. This is a convention, not a standard.

The compliler ignores files that end in `_test.go`.

When you test variations of the same function, you can use `TestFunction_Scenario1`, and `TestFunction_Scenario2` to distinguish. (Likely easier to use table tests.)

### Testing steps

A test calls the function and checks its returned value, or the state of some variable, against an expected value or state. Every test has 4 steps:
1. **Preparation**. We set up everything that we need to test: input values, expected outputs, env vars, global variables, network connections, etc.
2. **Execution**. Call the tested function. This step is usually a single line.
3. **Decision**. Check the output to the expected output with comparisons, evaluations.
4. **Teardown**. Clean up resources to get back to the beginning state. The `defer` keyword is useful here.

### Table tests

#### Testing scenarios

A scenario is a structure that aggregates data types into a meaningful entity, such as the requirements for a test:
```go
type scenario struct {
    lang             language
    expectedGreeting string
}
```
After you create the scenario struct, you can create a map of scenarios, where the key is the name of the test, and the value is the scenario structure:
```go
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
    ...
```

When you run the tests with `range`, the scenario name is the key, and the test cases are the values:
```go
// range over all scenarios
for scenarioName, tc := range tests {
    t.Run(scenarioName, func(t *testing.T) {
        greeting := greet(tc.lang)

        if greeting != tc.expectedGreeting {
            t.Errorf("expected: %q, got: %q", tc.expectedGreeting, greeting)
        }
    })
}
```


### Example syntax to test STDOUT of a function

If you test the STDOUT of a function, you can prepend it with `ExampleFuncName()`. Within that function, you can use Go's `Example` syntax. This syntax allows you to write commented lines to test the output. The first contains `Output:`, and the second contains the expected output.

For example, the following `Example` syntax tests a hello world program:

```go
func ExampleMain() {
	main()
	// Output:
	// Hello world
}
```

### Creating new types out of basic types

Using the proper type is important. Creating a custom type out of a basic type tells others what type is required, especially in parameters. For example:
```go
type language string
```

## Maps

A map is a hash table--a set of pairs of distinct keys and values. Defining a map is similar to defining a struct literal:
```go
var phrasebook = map[language]string{
	"el": "Χαίρετε Κόσμε",
	"en": "Hello world",
	"fr": "Bonjour le monde",
	"he": "שלום עולם",
	"ur": "ہیلو دنیا",
	"vi": "Xin chào Thế Giới",
}
```

When you try to access a value in a map, you can return a value and a `bool`. Use bracket syntax to provide the value that you are looking for. It is idiomatic to use `ok` for the `bool`. For example:
```go
greeting, ok := phrasebook[l]
if !ok {
    return fmt.Sprintf("unsupported language: %q", l)
}
return greeting
```

## Flags

The following is a simple example of a flag:
```go
var lang string
flag.StringVar(&lang, "lang", "en", "The required language, e.g. en, ur...")
flag.Parse()
```

The `flag` package provides similar functions for assigning flags. For example, `.StringVar` and `.String`. Use `.StringVar` when you create a variable and provide it as a pointer as the first argument. Use `.String` when you want to return a pointer:
```go
var lang string
flag.StringVar(&lang, "lang", "en", "The required language, e.g. en, ur...")

lang := flag.String("lang", "en", "The required...")
```

## Logger

Logging is when you keep track of the current state or events with readable messages. These messages are assigned an importance level: trace, bug, error, fatal, etc.

The most common are _debug_, _info_, and _error_. Declare them as an enumeration of constants:
```go
const (
	// LevelDebug represents the lowest level of log, mostly used for debugging purposes.
	LevelDebug Level = iota
	// LevelInfo represents a logging level that contains information deemed valuable.
	LevelInfo
	// LevelError represents the highest logging level, only to be used to trace errors.
	LevelError
)
```
Use `iota` to signify an enumeration.

When it makes sense, mimic library functions. For example, the following logger method mimic `fmt.Printf`'s signature:

```go
// Logger is used to log information
type Logger struct {
}

// Debugf formats and prints a message if the log level is debug or higher
func (l *Logger) Debugf(format string, args ...any) {

}
```

### iota

`= iota` lets the compiler know that we are starting an enumeration. It creates a sequence of numbers that increment on each line.

> You can use iota on any type that is based on an integer.

### Variadic functions

You might need to pass a variable number of parameters to a function: zero, one, or many. Use the variadiac function syntax--the last argument of a function is of the type `...{some type}`. It is common to see `...any`, because `any` is (sort of?) an alias for the empty inferface (`interface{}`). For example: 

```go
// Debugf formats and prints a message if the log level is debug or higher
func (l *Logger) Debugf(format string, args ...any) {

}
```

## Misc

### Multiple return values

You will handle multiple value assignment in the following common cases:
- Checking if a key is present in a map
- With the `range` keyword when iterating through a data structure
- Reading from a channel with the `<-` operator. This returns a value and wheher the channel is closed.
- Functions that return multiple values. 

### Quotes and strings

Go provides 3 types of quotes:

| Type         | Example | Description |
|--------------|:--------|:------------|
| double quote | `" "`   | Creates literal strings. |
| backtick     | \` \`    | Creates raw literal strings (cannot use escape sequence, such as `\n`). Use these so you do not have to escape double quotes.  |
| single quote | ' '     | Creates runes, a single unicode code point. |

### Pointers

Go functions pass argmuments by value. If you want to alter a value, pass the function a pointer to the value.

| Operator | Name                 | Description |
|----------|:---------------------|:------------|
| `&`      | Address operator     | Retrieves the address of a variable. |
| `*`      | Indirection operator | Access the value that the pointer points to (_follow_ the pointer) |

### go.mod

### Creating objects

If you do not use a constructor, you can create a zero-valued object:

```go
var log pocketlog.Logger
log := pocketlog.Logger{}
```

It is better to create a `NewXxx()` function that acts as a constructor. If the name of the object is obvious from the package name, just name it `New()`:

```go
// Logger is used to log information
type Logger struct {
	threshold Level
}

// New returns a logger, ready to log at the required threshold
func New(threshold Level) *Logger {
	return &Logger{
		threshold: threshold,
	}
}
```
Notice that the Logger struct has no exposed fields, and `New()` returns a pointer to a Logger. Returning a pointer is a best practice to conserve memory.