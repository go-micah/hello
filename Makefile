# Build the hello program
hello: 	
	go build -o bin/hello hello.go

# Run tests with verbose output
test:
	go test -v ./...

# Remove built binaries and clean up
clean:
	rm -rf bin/

# Run the program
run: hello
	./bin/hello

# Show help information
help:
	@echo "Available targets:"
	@echo "  make hello  - Build the program"
	@echo "  make test   - Run tests"
	@echo "  make run    - Build and run the program"
	@echo "  make clean  - Remove built binaries"
	@echo "  make help   - Show this help message"