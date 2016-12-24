#Time-stamp: <2017-01-16 04:01:09 hamada>

build:
	go build main.go 
	mv main run.x
	./run.x

f:
	gofmt main.go |tee tmp.go
	@diff main.go tmp.go

