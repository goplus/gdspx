.DEFAULT_GOAL := build

GOOS?=$(shell go env GOOS)
GOARCH?=$(shell go env GOARCH)
CLANG_FORMAT?=$(shell which clang-format | which clang-format-10 | which clang-format-11 | which clang-format-12)
CWD=$(shell pwd)

CURRENT_DIR_NAME = $(notdir $(CURDIR))

OUTPUT_PATH=project/lib
TEST_MAIN=main.go

ifeq ($(GOOS),windows)
	TEST_BINARY_PATH=$(OUTPUT_PATH)/gdspx-windows-$(GOARCH).dll
	GODOT="../../godot/bin/godot.windows.editor.x86_64.exe"
else ifeq ($(GOOS),darwin)
	TEST_BINARY_PATH=$(OUTPUT_PATH)/gdspx-macos-$(GOARCH).framework
	GODOT="../../godot/bin/godot.macos.editor.x86_64"
else ifeq ($(GOOS),linux)
	TEST_BINARY_PATH=$(OUTPUT_PATH)/gdspx-linux-$(GOARCH).so
	GODOT="../../godot/bin/godot.linuxbsd.editor.x86_64"
else
	TEST_BINARY_PATH=$(OUTPUT_PATH)/gdspx-$(GOOS)-$(GOARCH).so
	GODOT="../../godot/bin/godot.linuxbsd.editor.x86_64"
endif

.PHONY: goenv build clean run

goenv:
	go env

build: 
	CGO_ENABLED=1 \
	go build -buildmode=c-shared -o "$(TEST_BINARY_PATH)" $(TEST_MAIN)

buildfast: goenv
	go build .

clean:
	rm -f project/lib/libgodotgo-*

ci_gen_project_files:
	CI=1 \
	LOG_LEVEL=info \
	GOTRACEBACK=1 \
	GODEBUG=sbrk=1,gctrace=1,asyncpreemptoff=1,cgocheck=0,invalidptr=1,clobberfree=1,tracebackancestors=5 \
	$(GODOT) --headless --verbose --path project/ --editor --quit

run:
	$(MAKE) build 
	LOG_LEVEL=info \
	GOTRACEBACK=1 \
	GODEBUG=sbrk=1,gctrace=1,asyncpreemptoff=1,cgocheck=0,invalidptr=1,clobberfree=1,tracebackancestors=0 \
	$(GODOT) --debug  --path project/

editor:
	LOG_LEVEL=info \
	GOTRACEBACK=1 \
	GODEBUG=sbrk=1,asyncpreemptoff=1,cgocheck=0,invalidptr=1,clobberfree=1,tracebackancestors=0 \
	$(GODOT) --verbose --debug  --path project/ --editor

initload:
	LOG_LEVEL=info \
	GOTRACEBACK=1 \
	GODEBUG=sbrk=1,asyncpreemptoff=1,cgocheck=0,invalidptr=1,clobberfree=1,tracebackancestors=0 \
	$(GODOT) --verbose --debug  --path project/ --editor --quit


gen:
	cd ../../cmd && go run . && cd ../tutorial/$(CURRENT_DIR_NAME) && \
	$(MAKE) build
	
fmt:
	cd ../.. && go fmt ./... && cd tutorial/$(CURRENT_DIR_NAME) 