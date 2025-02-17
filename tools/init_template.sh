#!/bin/bash

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJ_DIR=$SCRIPT_DIR/..
cd $PROJ_DIR

GOPATH=$(go env GOPATH)
VERSION=$(cat ./cmd/gdspx/template/version)
echo "version=$VERSION GOPATH=$GOPATH"

pip install scons==4.7.0

# Install gdspx
cd cmd/gdspx/ || exit
go install .
cd ../../ || exit


PLATFORM=""
while getopts "p:" opt; do
    case "$opt" in
        p) PLATFORM="$OPTARG" ;;
        *) echo "Usage: $0 [-p platform]"; exit 1 ;;
    esac
done

cd godot || exit
echo "Init engine done."

GODOT_VERSION=4.2.2.stable
DST_DIR=""

if [[ "$(uname)" == "Linux" ]]; then
    DST_DIR="$HOME/.local/share/godot/export_templates/$GODOT_VERSION"
    if [ -z "$PLATFORM" ]; then
        PLATFORM="linux"
    fi
elif [[ "$(uname)" == "Darwin" ]]; then  # macOS
    DST_DIR="$HOME/Library/Application Support/Godot/export_templates/$GODOT_VERSION"
    if [ -z "$PLATFORM" ]; then
        PLATFORM="macos"
    fi
elif [[ "$(uname -o 2>/dev/null)" == "Msys" ]] || [[ "$(uname -o 2>/dev/null)" == "Cygwin" ]]; then
    # Windows (Git Bash, Cygwin, MSYS)
    DST_DIR="$APPDATA/Godot/export_templates/$GODOT_VERSION"
    if [ -z "$PLATFORM" ]; then
        PLATFORM="windows"
    fi
else
    echo "Unsupported OS"
    exit 1
fi

ARCH=x86_64
if [[ "$(uname -m)" == "aarch64" || "$(uname -m)" == "arm64" ]]; then
    ARCH=arm64
fi

echo "Destination binary path: $DST_DIR"
if [ "$PLATFORM" = "linux" ]; then

    scons target=template_debug dev_build=yes
    scons target=template_release dev_build=yes
    echo "Destination binary path: $DST_DIR"
    cp bin/godot.linuxbsd.template_* "$DST_DIR/"
    mv "$DST_DIR/godot.linuxbsd.template_debug"*  "$DST_DIR/linux_debug.$ARCH"
    mv "$DST_DIR/godot.linuxbsd.template_release"*  "$DST_DIR/linux_release.$ARCH"

elif [ "$PLATFORM" = "android" ]; then
    scons platform=android target=template_debug arch=arm32
    scons platform=android target=template_debug arch=arm64
    cd platform/android/java || exit
    # On Linux and macOS
    ./gradlew generateGodotTemplates

    cd $PROJ_DIR/godot || exit
    cp bin/android* "$DST_DIR/"
else
    echo "Unknown platform"
fi

cd $PROJ_DIR