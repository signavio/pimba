CURRENT_TAG := $(shell git describe --tags | sed -e 's/^v//')
ARCH = amd64

release:
	GOOS=linux GOARCH=$(ARCH) go build -o releases/pimba-$(CURRENT_TAG)-linux-amd64 main.go
	GOOS=freebsd GOARCH=$(ARCH) go build -o releases/pimba-$(CURRENT_TAG)-freebsd-amd64 main.go
	GOOS=darwin GOARCH=$(ARCH) go build -o releases/pimba-$(CURRENT_TAG)-darwin-amd64 main.go
	GOOS=windows GOARCH=$(ARCH) go build -o releases/pimba-$(CURRENT_TAG)-windows-amd64.exe main.go

image:
	docker build . -f tools/docker/Dockerfile -t signavio/pimba:$(CURRENT_TAG)
	docker build . -f tools/docker/Dockerfile -t signavio/pimba:latest

version:
	@echo $(CURRENT_TAG)

default: release
