.DEFAULT_GOAL := pc

CURRENT_PATH=$(shell pwd)
.PHONY: engine init initweb fmt gen upload updatemod templates cmd

# Format code
fmt:
	go fmt ./... 

# Update mod
updatemod:
	gdspx updatemod
# Build templates
templates:
	./tools/build_engine.sh -a
# Build current platform's engine
pc:
	./tools/build_engine.sh -e
# Build current platform's engine template
pcpack: 
	./tools/build_engine.sh
# Build web engine
web: 
	./tools/build_engine.sh -p web -e
# Build web engine template
webpack: 
	./tools/build_engine.sh -p web

# Build android engine
android:
	./tools/build_engine.sh -p android

# Build ios engine
ios:
	./tools/build_engine.sh -p ios 
# Generate code
gen:
	cd ./cmd/codegen && go run . && cd $(CURRENT_PATH) 

# Install gdspx command
cmd:
	cd cmd/gdspx/ && go install . && cd ../../ 

# Upload to web server
upload:
	./webserver/upload.sh 

%:
	@:
