.SILENT: build install clean

build:
	go build -o bin/boucles-v2 main.go

install: build
	cp bin/boucles-v2 /usr/local/bin/boucles-v2

clean:
	rm -rf ./bin