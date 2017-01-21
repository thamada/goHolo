#Time-stamp: <2017-01-21 13:08:37 hamada>

build:
	go build -o run.x main.go 
	./run.x

f:
	gofmt main.go |tee tmp.go
	@diff main.go tmp.go

