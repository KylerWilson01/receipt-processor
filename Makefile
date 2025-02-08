.PHONY: clean test watch build 

APP_NAME = receipt-processor
BUILD_DIR = $(PWD)/build
GOOS = linux
GOARCH = amd64
CGO_ENABLED = 0

clean:
	rm -rf ./build

test: clean 
	go test -timeout 30s -coverprofile=cover.out -cover ./...
	go tool cover -func=cover.out

watch: clean 
	air main.go

build: test
	GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=$(CGO_ENABLED) go build -tags netgo -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) main.go

run: clean
	go run .
