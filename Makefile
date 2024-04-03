VERSION := $(shell git describe --tags --abbrev=0 2>/dev/null || echo "unknown")

format:
	gofmt -s -w ./

lint:
	golint

test:
	go test -v

build:
	@echo "Building for $(TARGET)"
	CGO_ENABLE=0 GOOS=$(shell uname | tr '[:upper:]' '[:lower:]') GOARCH=$(TARGET) go build -v -o kbot-$(TARGET) -ldflags "-X 'github.com/wefgg/kbot/cmd.appVersion=${VERSION}'"

clean:
	rm -rf kbot-*
