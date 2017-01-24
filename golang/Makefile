#Time-stamp: <2017-01-24 12:31:03 hamada>

target = ./run.x

GOPATH := $(realpath .)
export GOPATH

default: build

build: vet
	go build -o ${target} main.go 
	${target}

vet:
	go vet ./src/...

# https://golang.org/cmd/go/#hdr-Run_gofmt_on_package_sources
fmt:
	go fmt -x -n ./src/...

gorun:
	go run main.go 

# https://github.com/golang/lint
# go get github.com/golang/lint/golint
lint:
	golint ./src

doc:
	godoc -http=:8080 -index



clean:
	rm -rf *~ .*~ golang ${target} src/*~ 


c: clean

