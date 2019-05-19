GO ?= GO111MODULE=on go

all: build

build:
	$(GO) build -o vmctl main.go

clean:
	rm vmctl
