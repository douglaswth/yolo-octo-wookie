ifeq ($(OS),Windows_NT)
EXEEXT:=.exe
endif

PROGRAM:=yolo-octo-wookie$(EXEEXT)

.PHONY: all test clean

all: $(PROGRAM)

$(PROGRAM): $(filter-out %_test.go,$(wildcard *.go))
	go build -v ./...

depend:
	go get -v github.com/onsi/ginkgo/ginkgo
	go get -v github.com/onsi/gomega

test: $(PROGRAM)
	go test -v ./...

clean:
	go clean
