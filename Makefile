TARGET=wording

all: tests build

build: prebuild
	@go build -o bin/$(TARGET) main.go

release: tests prebuild
	@CGO_ENABLED=1 go build -ldflags="-s -w" -a -installsuffix cgo -o bin/$(TARGET) main.go
	@cp -r public bin/public

windows: tests prebuild
	CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -o bin/$(TARGET).exe main.go
	@cp -r public bin/public

prebuild:
	@tailwindcss -o ./public/css/styles.css -c ./tailwind.config.js
	@templ generate view

nrun: build run

run:
	@bin/$(TARGET)

tests: test

test:
	@go test ./tests/...

testV:
	go test -v ./tests/...

dev:
	@air -c ".tdd.air.toml"

webdev:
	@air -c ".air.toml"

clean:
	@rm -rf bin/public
	@rm bin/*

.PHONY: build prebuild release nrun run test tests dev clean
