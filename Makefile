.PHONY: all deps build test vet run docker clean

# Default target to build and test
all: deps vet build test

# Get the dependencies
deps:
	go mod tidy

# Build the binary
build:
	go build -o medevac main.go

# Run tests
test:
	go test -v ./...

# Run go vet
vet:
	go vet ./...

# Run the application
run:
	./medevac

# Build Docker image
docker:
	docker build -t your-docker-repo/medevac:latest .

# Clean build files
clean:
	rm -f medevac
