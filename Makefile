GOOS := darwin
GOARCH := amd64

.PHONY: all clean deps build install

all: clean deps build

deps:
	go get -u github.com/aws/aws-sdk-go
	go get -u github.com/jessevdk/go-flags

build:
	mkdir -p bin
	$(eval targetfiles := $(shell ls -1 | grep -v "base.go" | grep -E ".*\.go$$"))
	$(eval binnames := $(basename $(targetfiles)))
	for binname in $(binnames); do \
		mkdir -p build/$$binname; \
		cp -p base.go build/$$binname/; \
		cp -p $$binname.go build/$$binname/; \
		cd build/$$binname; \
		GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o ../../bin/$$binname; \
		cd -; \
	done

install:
	install -m 0755 bin/* /usr/local/bin/

clean:
	rm -fr ./build
	rm -fr ./bin
