# Hello, World in Go

A beginner-friendly Go project to help you write your first Go program!

## What You'll Learn

- How to structure a basic Go program
- Package declarations and imports
- Functions and return types
- Variables and string formatting
- Command-line arguments
- Writing and running tests
- Table-driven testing (a Go idiom)

## Prerequisites

- Go 1.20 or later installed ([Download Go](https://go.dev/dl/))
- Basic command-line knowledge

## Common Commands

### Build the program
```bash
make hello
```

### Run the program
```bash
# Run with default greeting
./bin/hello

# Run with a custom name
./bin/hello YourName
```

Or use the Makefile:
```bash
make run
```

### Run tests
```bash
make test
```

### Clean up built binaries
```bash
make clean
```

### See all available commands
```bash
make help
```

## Expected Output

Running `./bin/hello` will output:
```
Hello, World
```

Running `./bin/hello Gopher` will output:
```
Hello, Gopher
```

## Project Structure

```
.
├── hello.go       # Main program with comments explaining Go concepts
├── hello_test.go  # Tests demonstrating table-driven testing
├── Makefile       # Build automation
├── go.mod         # Go module definition
└── bin/           # Built binaries (created after building)
```

## Troubleshooting

**"go: command not found"**
- Make sure Go is installed and in your PATH
- Run `go version` to verify installation

**"permission denied" when running ./bin/hello**
- Run `chmod +x bin/hello` to make it executable

**Tests failing**
- Make sure you've saved all changes to the files
- Try running `go mod tidy` to ensure dependencies are correct

## Next Steps

Once you're comfortable with this project, try:

1. **Add more features:**
   - Accept multiple names and greet them all
   - Add a flag for different greeting styles (formal, casual, etc.)
   - Read names from a file

2. **Learn more Go concepts:**
   - Structs and methods
   - Error handling with multiple return values
   - Reading user input with `bufio.Scanner`
   - Working with JSON

3. **Explore the standard library:**
   - `net/http` for building web servers
   - `encoding/json` for working with JSON
   - `io` for file operations

## Resources

- [A Tour of Go](https://go.dev/tour/) - Interactive introduction to Go
- [Go by Example](https://gobyexample.com/) - Hands-on examples
- [Effective Go](https://go.dev/doc/effective_go) - Best practices guide
