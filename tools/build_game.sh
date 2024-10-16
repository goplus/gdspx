#!/bin/bash

# Get current directory name
CURRENT_PATH=("$PWD")

# Default values
EXPORT_WEB=false
PROJECT_PATH="tutorial/01_aircraft"
if [ "$OS" = "Windows_NT" ]; then
    GODOT="./godot/bin/godot.windows.editor.dev.x86_64"
elif [ "$OS" = "Linux" ]; then
    GODOT="./godot/bin/godot.linuxbsd.editor.dev.x86_64"
else
    GODOT="./godot/bin/godot.darwin.editor.dev.x86_64"
fi


PLATFORM=""
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

GD_PROJ_DIR="$PROJECT_PATH" 
GD_LIB_DIR=$GD_PROJ_DIR/.godot
LIB_DIR=$GD_PROJ_DIR/lib

# delete gdextension 
rm -rf $LIB_DIR
rm -f $GD_LIB_DIR/extension_list.cfg
rm -f $GD_PROJ_DIR/gdspx.gdextension 

mkdir -p $GD_LIB_DIR
mkdir -p $LIB_DIR

TEMPLATE_DIR=$CURRENT_PATH/cmd/gdspx/pkg/impl/template/
cp $TEMPLATE_DIR/export_presets.cfg $GD_PROJ_DIR/export_presets.cfg 

if [ "$PLATFORM" = "web" ]; then
    # update extension list
    mkdir -p $PROJECT_PATH/.builds/web

    # Build wasm
    cd "$PROJECT_PATH" || { echo "Project path not found"; exit 1; }
    echo "Building Go wasm"
    GOOS=js GOARCH=wasm go build -o ".builds/web/gdspx.wasm"
    cd $CURRENT_PATH

    # Export to web if enabled
    if [ "$EXPORT_WEB" = true ]; then
        echo "================ EXPORT_WEB ====================="
        $GODOT --headless --quit --path $GD_PROJ_DIR --export-debug Web 
    fi
elif  [ "$PLATFORM" = "" ]; then #pc
    # update extension list
    cp $TEMPLATE_DIR/extension_list.cfg $GD_LIB_DIR/extension_list.cfg 
    cp $TEMPLATE_DIR/gdspx.gdextension $GD_PROJ_DIR/gdspx.gdextension 

    cd $PROJECT_PATH 
	CGO_ENABLED=1 go build -buildmode=c-shared -o ./lib/$LIB_NAME main.go
    cd $CURRENT_PATH
    echo $GODOT --debug  --path $GD_PROJ_DIR --editor 
    if [ "$EDITOR" = true ]; then
        $GODOT --debug  --path $GD_PROJ_DIR --editor 
    else 
        $GODOT --debug  --path $GD_PROJ_DIR
    fi 
fi 