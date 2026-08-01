# Default task (run tests)
.PHONY: all test dev coverage build publish clean

# Run all tests
test:
	go test ./...

# Run tests with auto-reload (watch mode)
# Note: Install air/gotestexec if needed, or use go test in loop
dev:
	go test -v ./...

# Generate test coverage report and open HTML in browser
coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated at coverage.html"

# Run local development server
serve:
	go run demo/server.go

# Build task: runs tests and builds the project
build: test
	go build -v ./...
	go run build.go

# Tag and prepare release for Go module
publish: build
	@echo "Publishing new version..."
	git tag v1.0.0
	git push origin v1.0.0

# Cleanup generated build and coverage files
clean:
	rm -f coverage.out coverage.html
