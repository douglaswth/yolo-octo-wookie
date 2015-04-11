ifeq ($(OS),Windows_NT)
EXEEXT:=.exe
export CC:=x86_64-w64-mingw32-gcc
export CXX:=x86_64-w64-mingw32-g++
endif

PROGRAM:=yolo-octo-wookie$(EXEEXT)

.PHONY: all test clean

all: $(PROGRAM)

$(PROGRAM): $(filter-out %_test.go,$(wildcard *.go))
	go build -v ./...

test: $(PROGRAM)
	go test -v ./...

clean:
	go clean
