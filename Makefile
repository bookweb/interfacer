.PHONY: install-deps
install-deps:
	go install github.com/bookweb/interfacer/cmd/interfacer@latest

.PHONY: build

build:
	go build -o interfacer cmd/interfacer/main.go

generate:
	go generate

test-generate:
	./interfacer generate --type MyStruct --receiver myStruct --output my_struct.gen.go .