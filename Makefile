BINARY := zanao
BUILD_FLAGS := -ldflags="-s -w"

build:
	go build $(BUILD_FLAGS) -o build/$(BINARY) .

install:
	go install $(BUILD_FLAGS) .

build-all:
	GOOS=linux   GOARCH=amd64  go build $(BUILD_FLAGS) -o build/$(BINARY)-linux-amd64 .
	GOOS=linux   GOARCH=arm64  go build $(BUILD_FLAGS) -o build/$(BINARY)-linux-arm64 .
	GOOS=darwin  GOARCH=amd64  go build $(BUILD_FLAGS) -o build/$(BINARY)-darwin-amd64 .
	GOOS=darwin  GOARCH=arm64  go build $(BUILD_FLAGS) -o build/$(BINARY)-darwin-arm64 .
	GOOS=windows GOARCH=amd64  go build $(BUILD_FLAGS) -o build/$(BINARY)-windows-amd64.exe .

clean:
	rm -rf build/

.PHONY: build install build-all clean
