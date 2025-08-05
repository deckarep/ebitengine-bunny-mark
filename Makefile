# Adapted from: https://github.com/programatta/space-invaders
# Thank You :)

VERSION := 0.0.1 # Version of game.

DIST_DIR := bin
WEB_DIR := ${DIST_DIR}/web
WEB_WASM := $(WEB_DIR)/ebiten-bunny-mark.wasm
MODULE := .

.PHONY: build build-doc build-web build-mac-arm run run-web clean

build:
	go build -ldflags "-X main.Version=$(VERSION)" -o ${DIST_DIR}/ebiten-bunny-mark *.go

build-win:
	env GOOS=windows GOARCH=amd64 go build -ldflags "-X main.Version=$(VERSION)" -o ${DIST_DIR}/ebiten-bunny-mark.exe *.go

# Requiere de un OSX para realizar compilación nativa con bindings de C
# build-mac:
# 	env GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.Version=$(VERSION)" -o ${DIST_DIR}/ebiten-bunny-mark-mac *.go

build-mac-arm:
	env GOOS=darwin GOARCH=arm64 go build -ldflags "-X main.Version=$(VERSION)" -o ${DIST_DIR}/ebiten-bunny-mark-macarm *.go

build-web:
	mkdir -p ${WEB_DIR}
	env GOOS=js GOARCH=wasm go build -ldflags "-X main.Version=$(VERSION)" -buildvcs=false -o ${WEB_WASM} ${MODULE}
	cp $$(go env GOROOT)/lib/wasm/wasm_exec.js ${WEB_DIR}
	printf '%s\n' \
	'<!DOCTYPE html>' \
	'  <head>' \
	'    <meta charset="UTF-8">' \
	'    <title>Ebitengine Bunny Mark</title>' \
	'    <link rel="shortcut icon" type="image/x-icon" href="favicon.ico">' \
	'  </head>' \
	'<style>' \
    '	p {' \
    '		position: absolute;' \
    '		left: 25%;' \
    '		width: 50%;' \
    '		height: 50%;' \
    '		top: 25%;' \
    '		margin: auto;' \
    '		z-index: -100;' \
    '		background: url(gopher-running.svg) center center no-repeat;' \
    '	}' \
    '</style>' \
	'  <body>' \
	'	 <p id="loading_text"></p>' \
	'    <script src="wasm_exec.js"></script>' \
	'    <script>' \
	'    	async function loadGame() {' \
    '       	const go = new Go();' \
    '       	const response = await fetch("ebiten-bunny-mark.wasm");' \
    '       	const buffer = await response.arrayBuffer();' \
    '       	const obj = await WebAssembly.instantiate(buffer, go.importObject);' \
    '       	await new Promise(r => setTimeout(r, 500));' \
    '       	go.run(obj.instance);' \
    '       	document.getElementById("loading_text").style.display = "none";' \
    '       }' \
    '       loadGame();' \
	'    </script>' \
	'  </body>' \
	> ${WEB_DIR}/index.html
	# Finally, change to web directory and create the archive.
	( cd ${WEB_DIR} && zip archive.zip index.html ebiten-bunny-mark.wasm wasm_exec.js )

build-docs: build-web
	cp ${WEB_DIR}/wasm_exec.js docs/
	cp ${WEB_DIR}/index.html docs/
	cp ${WEB_DIR}/ebiten-bunny-mark.wasm docs/

build-all: build build-mac-arm build-win build-web

run:
	go run *.go

run-web:
	go run github.com/hajimehoshi/wasmserve@latest .

clean:
	rm -rf ${DIST_DIR}