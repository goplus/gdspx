.DEFAULT_GOAL := build

GOOS?=$(shell go env GOOS)
GOARCH?=$(shell go env GOARCH)
CLANG_FORMAT?=$(shell which clang-format | which clang-format-10 | which clang-format-11 | which clang-format-12)
CURRENT_PATH=$(shell pwd)
PROJECT_PATH?="./tutorial/01_aircraft"  # Use ?= for default value
.PHONY: engine init run fmt gen server wasm web pc editor

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
	./tools/build_engine.sh && \
	./tools/start_server.sh 

server:
	./tools/start_server.sh 

engine: 
	./tools/build_engine.sh 
	$(MAKE) web 

build: 
	./tools/build_game.sh --platform pc --gd $(GODOT) --lib $(LIB_PATH) --path $(PROJECT_PATH)

pc:
	$(MAKE) fmt 
	./tools/build_game.sh --platform pc --gd $(GODOT) --lib $(LIB_PATH) --path $(PROJECT_PATH)
web:
	$(MAKE) fmt 
	./tools/build_game.sh --platform web --gd $(GODOT) --lib $(LIB_PATH) --export --path $(PROJECT_PATH)

editor:
	./tools/build_game.sh --platform pc --gd $(GODOT) --lib $(LIB_PATH) --path $(PROJECT_PATH) --editor

wasm:
	./tools/build_game.sh --platform web --gd $(GODOT) --lib $(LIB_PATH) --path $(PROJECT_PATH)


gen:
	cd ./cmd && go run . && cd $(CURRENT_PATH) && \
	$(MAKE) wasm
	$(MAKE) fmt 
