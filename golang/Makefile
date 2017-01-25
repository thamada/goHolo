#Time-stamp: <2017-01-25 14:24:37 hamada>

target = ./bin/golang

GOPATH := $(realpath .)
GOBIN := $(realpath .)/bin
export GOPATH
export GOBIN

default: build

run:
	${target}

build: vet
	go get
	go build -o ${target} main.go 
	${target}

vet:
	go vet ./src/...

# https://golang.org/cmd/go/#hdr-Run_gofmt_on_package_sources
fmt:
	go fmt . ./src/...
#	go fmt -x -n ./src/...

gorun:
	go run main.go 

# https://github.com/golang/lint
# go get github.com/golang/lint/golint
lint:
	golint ./src

doc:
	godoc -http=:8080 -index

clean:
	rm -rf *~ .*~ golang ${target} src/*~ .DS_Store


c: clean

r: run

