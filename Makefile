GOROOT=$(shell go env GOROOT)

build:
	cp "$(GOROOT)/misc/wasm/wasm_exec.js" bin/wasm_exec.js
	rm -f bin/generate.wasm && GOOS=js GOARCH=wasm go build -o ./bin/generate.wasm

check:
	shasum bin/* | tee /tmp/shasum
	make build
	shasum -c /tmp/shasum bin/*
