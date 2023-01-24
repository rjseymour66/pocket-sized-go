# Pocket-sized Go

## Testing

You can name files `hello_world_internal_test.go`. The `internal` signifies that you are testing unexposed/private methods. This is a convention, not a standard.

The compliler ignores files that end in `_test.go`.

When you test variations of the same function, append an underscore and a description of the expected output or test scenario to the test function name. For example, you can use `TestFunction_Scenario1`, and `TestFunction_Scenario2` to distinguish different scenarios. (Likely easier to use table tests.)

### Package names

Append the `_test` to the package name of a test file. For example, a file that tests the `example` package can use the `example_test` package, and the Go compiler will not complain if you include it in the `example` package.

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


If you test the STDOUT of a function, you can prepend it with `ExampleFuncName()`. Within that function, you can use Go's `Example` syntax. This syntax allows you to write commented lines to test the output. The first contains `Output:`, and the second contains the expected output. This kind of testing is called _open-box_ testing.

For example, the following `Example` syntax tests a hello world program:

```go
func ExampleMain() {
	main()
	// Output:
	// Hello world
}
```
> Notice that the `Example` syntax does not accept as an argument the `*testing` object.

_Closed-box_ testing is when you test a system from the outside. This type of testing is useful with the `Example` syntax. 


### Creating new types out of basic types

Using the proper type is important. Creating a custom type out of a basic type tells others what type is required, especially in parameters. For example:
```go
type language string
```

### Mocking interfaces

If you have to test a function that uss an interface that is not standard, you can create a type that holds the same type of data that your function uses, and then implement the interface on that type.

For example, you need to mock the `io.Writer` interface because you are testing a logger that writes a `string` to an `io.Writer`. The `io.Writer` interface writes a string of bytes and returns the number of bytes written and an error value:
```go
Write(p []byte) (n int, err error)
```

To mock this interface, create a type that holds the `string` type:
```go
type testWriter struct {
	contents string
}
```

Next, implement the `io.Writer` interface on the `testWriter` type:
```go
// Write implements the io.Writer interface
func (tw *testWriter) Write(p []byte) (n int, err error) {
	tw.contents = tw.contents + string(p)
	return len(p), nil
}
```
Now, you can execute tests using your function, and pass the `testWriter` to the code that you are testing:
```go
testedLogger := pocketlog.New(tc.level, pocketlog.WithOutput(tw))

testedLogger.Debugf(debugMessage)
testedLogger.Infof(infoMessage)
testedLogger.Errorf(errorMessage)

if tw.contents != tc.expected {
    t.Errorf("invalid contents, expected %q, got %q", tc.expected, tw.contents)
}
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

## Logging

Logging is when you keep track of the current state or events with readable messages. These messages are assigned an importance level: trace, bug, error, fatal, etc.
> Logs are not a debugging tool.

Log storage can be expensive--you should understand what you must log, and how format them. Here are some tips:
- Write clear messages.
- Be concise. Do not exceed 1000 characters or bytes. You might want to set a limit in the logging function that truncates messages that exceed the limit.
- Create structured messages (i.e. JSON). Often, programs--not humans--read logs and to display them in dashboards.



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

#### Returning Printf vals

If you are writing a function that doesn't return anything but logs to the console, use two blank identifiers for the values that `fmt.Printf` returns (int, error):
```g0
_, _ = fmt.Printf(format+"\n", args...)
```
The blank identifier means that we know that the function returns values, but we do not want to use them.

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

`go mod download`
: Downloads the dependencies for a project.

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

## Documentation

Must read: [Godoc: documenting Go code](https://go.dev/blog/godoc)

Use `go doc <path>` to return the documentation for a package or symbol (ex: function) within the `<path>` directory and subdirectories. For example, the following commands are executed in the `pocketlog` directory:

```shell
pocketlog
├── level.go
├── logger.go
└── logger_test.go

$ go doc .
package pocketlog // import "log/pocketlog"

type Level byte
    const LevelDebug Level = iota ...
type Logger struct{ ... }
    func New(threshold Level) *Logger

$ go doc New
package pocketlog // import "."

func New(threshold Level) *Logger
    New returns a logger, ready to log at the required threshold.
```
## Interfaces

Interfaces let you perform actions to a variety of inputs and outputs if they conform to the correct interface contract. Go provides standard interfaces that address many common use cases.

### io.Writer

An object of type `io.Writer` can write to any destination that implements its interface.

```go
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

### io.Reader

An object of type `io.Reader` can read from any source that implements its interface:

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

## Strategies 

### Functional Options pattern

The Functional Options pattern creates flexible constructors for Go types. It simplifies configuration, and lets the user set specific behaviors [without altering the API](https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis).

1. Create a new file called `options.go`.
2. Define a new exported function type (a new type of function) called `Options`. The function should accept as an argument a pointer of the type you are implementing this pattern on, so the function can change the type.
   For example, the following type is for a logger of type `Logger`:
   ```go
   type Option func(*Logger)
   ```
3. For each constructor configuration option that you want to add, create a function that returns a function of type `Option`:
   - Name the function `WithXxx`, where `Xxx` is the name of the configuration that the function sets. The function takes as a parameter the option that you are setting.
   - The return function of type `Option` sets the option:
   ```go
   // WithOutput returns a configuration function that sets the output of logs.
   func WithOutput(output io.Writer) Option {
	   return func(lgr *Logger) {
		   lgr.output = output
	   }
   }
   ``` 
4. Update the constructor type to accept variable number of `Options`:
   ```go
   // New returns a logger, ready to log at the required threshold.
   func New(threshold Level, opts ...Option) *Logger {
        // implementation
   }
   ```
5. In the constructor, set any default values, then loop through `opt` and execute the provided `Option` functions on the object that the constructor returns:
   ```go
   // New returns a logger, ready to log at the required threshold.
   func New(threshold Level, opts ...Option) *Logger {
	   // set defaults
	   lgr := &Logger{threshold: threshold, output: os.Stdout}

	   // add config options
	   for _, configFunc := range opts {
		   configFunc(lgr)
	   }

	   return lgr
   }
   ```

## bufio

TODO Reader vs Scanner

## Character encoding

All Go source files are encoded in the Unicode Transformation Format (UTF), a character encoding format that is able to encode all possible character code points in Unicode. UTF-8 means that the character is encoded using 8 bits (a byte).

### History

[Read this](https://deliciousbrains.com/how-unicode-works/).

#### ASCII

First, the most-commonly used encoding was the American Standard Code for Information Interchange (ASCII). This encoding covered the latin alphabet. In ASCII, each character gets one byte, which means that you can only represent 256 characters.

| Character | ASCII code | Byte representation |
|-----------|------------|------|
| A | 065 | 01000001 |
| a | 097 | 01100001 |
| B | 066 | 01000010 |
...



#### Unicode

Unicode can represent every symbol and every language. Unicode assigns a _code point_ to each character. Each code point is a a number that is given meaning by the Unicode standard.

| Character | Code point |
|-----------|------------|
| A | U+0041 |
| a | U+0061 |

#### UTF-8

UTF-8 is an encoding system for Unicode. It represents characters in a set of one to four bytes. The first characters in the Unicode library use one byte--the first 128 use the ASCII binary representation--but later characters use up to four bytes. 

| Character | Code point | UTF-8 binary encoding |
|-----------|------------|-----------------------|
| A | U+0041 | 01000001 |
| a | U+0061 | 01100001 |
| 0 | U+0030 | 00110000 |
| 9 | U+0039 | 00111001 |
...

## runes

runes support [Unicode](https://en.wikipedia.org/wiki/Unicode), which means that a rune is a Unicode codepoint.

For example, the following code prints out the length in UTF-8, and then the length in runes (Unicode code points). Notice that `Hello, 世界` is 9 characters:

```go
fmt.Println(len("Hello, 世界"))         // 13 UTF-8 bytes
fmt.Println(len([]rune("Hello, 世界"))) //  9 Unicode code points
```
When you iterate over a string, you can access its byte representation or its rune representation:

```go
fmt.Println([]byte(str))    // [82 121 97 110 32 83 101 121 109 111 117 114]
fmt.Println([]rune(str))    // [82 121 97 110 32 83 101 121 109 111 117 114]
```

You cannot convert a byte slice to a rune slice directly--it returns an error. First, you must convert the bytes to a `string`, then to a `[]rune`:

```go
bSlice := []byte{'s', 'l', 'i', 'c', 'e'}
fmt.Println("byte slice\t", bSlice)                     // [115 108 105 99 101]
fmt.Println("byte -> string\t", string(bSlice))         // slice
fmt.Println("string -> rune\t", []rune(string(bSlice))) // [115 108 105 99 101]
```