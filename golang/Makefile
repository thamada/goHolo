#Time-stamp: <2017-01-31 00:19:08 hamada>

target = ./bin/golang

GOPATH := $(realpath .)
GOBIN := $(realpath .)/bin
export GOPATH
export GOBIN

default: build

run:
	@${target}

build: vet fmt
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

docweb:
	godoc -http=:8080 -index

doc:
	go doc sync WaitGroup

lib:
	go build -buildmode=c-archive -o libgo.a main.go

clean:
	rm -rf *~ .*~ golang ${target} src/*~ .DS_Store
	rm -rf ./pkg/*

c: clean

r: run

