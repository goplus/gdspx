.DEFAULT_GOAL := build

GOOS?=$(shell go env GOOS)
GOARCH?=$(shell go env GOARCH)
CLANG_FORMAT?=$(shell which clang-format | which clang-format-10 | which clang-format-11 | which clang-format-12)
CURRENT_PATH=$(shell pwd)
PROJECT_PATH?="./tutorial/02_mario"  # Use ?= for default value
.PHONY: engine init run fmt gen server wasm web pc editor upload

# Check if an argument is passed to 'run' target, only then override PROJECT_PATH
override PROJECT_PATH := $(if $(word 2,$(MAKECMDGOALS)),$(word 2,$(MAKECMDGOALS)),$(PROJECT_PATH))

ifeq ($(GOOS),windows)
	LIB_PATH=gdspx-windows-$(GOARCH).dll
	GODOT="./godot/bin/godot.windows.editor.dev.x86_64.exe"
else ifeq ($(GOOS),darwin)
	LIB_PATH=gdspx-macos-$(GOARCH).framework
	GODOT="./godot/bin/godot.macos.editor.dev.x86_64"
else ifeq ($(GOOS),linux)
	LIB_PATH=gdspx-linux-$(GOARCH).so
	GODOT="./godot/bin/godot.linuxbsd.editor.dev.x86_64"
else
	LIB_PATH=gdspx-$(GOOS)-$(GOARCH).so
	GODOT="./godot/bin/godot.linuxbsd.editor.dev.x86_64"
endif



fmt:
	go fmt ./... 

init:
	./tools/init.sh

server:
	./tools/start_server.sh 

engine: 
	./tools/build_engine.sh 

build: 
	./tools/build_game.sh --gd $(GODOT) --lib $(LIB_PATH) --path $(PROJECT_PATH)

run:
	$(MAKE) fmt 
	./tools/build_game.sh --gd $(GODOT) --lib $(LIB_PATH) --path $(PROJECT_PATH)
web:
	$(MAKE) fmt 
	./tools/build_game.sh --gd $(GODOT) --lib $(LIB_PATH) --path $(PROJECT_PATH) --export  --platform web

editor:
	./tools/build_game.sh --gd $(GODOT) --lib $(LIB_PATH) --path $(PROJECT_PATH) --editor

wasm:
	./tools/build_game.sh --gd $(GODOT) --lib $(LIB_PATH) --path $(PROJECT_PATH) --platform web


gen:
	cd ./cmd/codegen && go run . && cd $(CURRENT_PATH) && \
	$(MAKE) fmt && $(MAKE) fmt 
	
upload:
	./webserver/upload.sh 

initload:
	LOG_LEVEL=info \
	GOTRACEBACK=1 \
	GODEBUG=sbrk=1,asyncpreemptoff=1,cgocheck=0,invalidptr=1,clobberfree=1,tracebackancestors=0 \
	$(GODOT) --verbose --debug  --path  $(PROJECT_PATH)/project/ --editor --quit