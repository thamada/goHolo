
BIT = 64
TARGET_BASE = test


build:
	gcc -c -fpic ${TARGET_BASE}.c
	gcc -c -fpic ${TARGET_BASE}_wrap.c
	gcc -shared ${TARGET_BASE}.o ${TARGET_BASE}_wrap.o -o ${TARGET_BASE}.so
	go tool 8g ${TARGET_BASE}.go
	go tool 8c -I/usr/local/Cellar/go/1.7.4_2/libexec/src/runtime ${TARGET_BASE}_gc.c
	go tool pack grc ${TARGET_BASE}.a ${TARGET_BASE}.8 ${TARGET_BASE}_gc.8

swig:
	swig -go -intgosize ${BIT} ${TARGET_BASE}.i


clean:
	rm -rf ${TARGET_BASE}.o ${TARGET_BASE}_{wrap,gc}.{c,o} ${TARGET_BASE}.go

c: clean


