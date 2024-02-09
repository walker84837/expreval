operatingsystem := $(shell go env GOOS)

.PHONY: all clean release debug

all:
	@echo "Building for '$(operatingsystem)'..."
	go build -ldflags="-s -w" -gcflags="-N -l" ./...

release: all
	@echo "Building release version..."
	CGO_ENABLED=0 go build -ldflags="-s -w" -gcflags="-trimpath=$(PWD)" ./...

debug:
	@echo "Building debug version..."
	go build -ldflags="-s -w" -gcflags="" ./...

clean:
	@rm -rf ./*.a *.o core *.log
