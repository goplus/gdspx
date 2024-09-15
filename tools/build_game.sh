#!/bin/bash

# Get current directory name
CURRENT_PATH=("$PWD")

# Default values
EXPORT_WEB=false
PROJECT_PATH="tutorial/01_aircraft"
GODOT="./godot/bin/godot.linuxbsd.editor.dev.x86_64"
PLATFORM="pc"
EDITOR=false
LIB_NAME=false
# Parse command line arguments
while [[ "$#" -gt 0 ]]; do
    case $1 in
        --export) EXPORT_WEB=true ;;
        --path) PROJECT_PATH="$2"; shift ;;
        --gd) GODOT="$2"; shift ;;
        --platform) PLATFORM="$2"; shift ;;
        --editor) EDITOR=true ;;
        --lib) LIB_NAME="$2"; shift ;;
        *) echo "Unknown parameter passed: $1"; exit 1 ;;
    esac
    shift
done

GD_PROJ_DIR="$PROJECT_PATH/project" 
GD_LIB_DIR=$GD_PROJ_DIR/.godot
LIB_DIR=$GD_PROJ_DIR/lib

# delete gdextension 
rm -rf $LIB_DIR
rm -f $GD_LIB_DIR/extension_list.cfg
rm -f $GD_PROJ_DIR/gdspx.gdextension 


mkdir -p $GD_LIB_DIR
mkdir -p $LIB_DIR

if [ "$PLATFORM" = "web" ]; then
    # update extension list
    mkdir -p ./builds/web_go

    # Build wasm
    cd "$PROJECT_PATH" || { echo "Project path not found"; exit 1; }
    echo "Building Go wasm"
    GOOS=js GOARCH=wasm go build -tags platform_web -o "$CURRENT_PATH/builds/web_go/gdspx.wasm"
    cd $CURRENT_PATH

    # Export to web if enabled
    if [ "$EXPORT_WEB" = true ]; then
        echo "================ EXPORT_WEB ====================="
        $GODOT --headless --quit --path $GD_PROJ_DIR --export-debug "Web" "$CURRENT_PATH/builds/web_go/index.html"
    fi
elif  [ "$PLATFORM" = "pc" ]; then
    # update extension list
    echo "res://gdspx.gdextension" > $GD_LIB_DIR/extension_list.cfg
    TEMPLATE_DIR=$CURRENT_PATH/cmd/template/project/gdspx.gdextension
    cp $TEMPLATE_DIR $GD_PROJ_DIR/gdspx.gdextension 

    cd $PROJECT_PATH 
	CGO_ENABLED=1 go build -tags platform_pc -buildmode=c-shared -o ./project/lib/$LIB_NAME main.go
    cd $CURRENT_PATH
    
    if [ "$EDITOR" = true ]; then
        $GODOT --debug  --path $GD_PROJ_DIR --editor 
    else 
        $GODOT --debug  --path $GD_PROJ_DIR
    fi 
fi 