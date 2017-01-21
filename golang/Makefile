#Time-stamp: <2017-01-21 13:09:24 hamada>

build:
	go build -o run.x main.go 
	./run.x

gorun:
	go run main.go 

f:
	gofmt main.go |tee tmp.go
	@diff main.go tmp.go

