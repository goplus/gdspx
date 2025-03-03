#!/bin/bash

# install gdspx
cd cmd/gdspx/
go install .
cd ../../

GOPATH=$(go env GOPATH)
VERSION=$(cat ./cmd/gdspx/template/version)
echo "version="$VERSION " GOPATH=" $GOPATH

ARCH=x86_64
if [[ "$(uname -m)" == "aarch64" || "$(uname -m)" == "arm64" ]]; then
    ARCH=arm64
fi


pip3 install scons==4.7.0
if [ ! -d "godot" ]; then
    echo "Godot directory not found. Creating and initializing..."
    mkdir godot
    cd godot
    git init 
    git remote add origin git@github.com:realdream-ai/godot.git
    git fetch --depth 1 origin spx4.2.2
    git checkout spx4.2.2
    echo "Godot repository setup complete."
else
    cd godot
    echo "Godot directory already exists."
fi
#  dev_build=yes

if [ "$OS" = "Windows_NT" ]; then
    scons target=editor arch=$ARCH vsproj=yes dev_build=yes
else
    scons target=editor arch=$ARCH dev_build=yes
fi
cd ..

echo "init engine done."
dstBinPath="$GOPATH/bin/gdspx$VERSION"
echo "Destination binary path: $dstBinPath"
if [ "$OS" = "Windows_NT" ]; then
    cp godot/bin/godot.windows.editor.dev.$ARCH $dstBinPath"_win.exe"
elif [[ "$(uname)" == "Linux" ]]; then
    cp godot/bin/godot.linuxbsd.editor.dev.$ARCH $dstBinPath"_linux"
else
    cp godot/bin/godot.macos.editor.dev.$ARCH $dstBinPath"_darwin"
fi


if [ "$1" != "-a" ]; then
    exit
fi

# build release version
cd godot
if [ "$OS" = "Windows_NT" ]; then
    scons target=editor arch=$ARCH vsproj=yes optimize=size
else
    scons target=editor arch=$ARCH optimize=size
fi
cd ..

dstBinPath="$GOPATH/bin/gdspx$VERSION"
echo "Destination binary path: $dstBinPath"
if [ "$OS" = "Windows_NT" ]; then
    cp godot/bin/godot.windows.editor.$ARCH $dstBinPath"_win_prod.exe"
elif [[ "$(uname)" == "Linux" ]]; then
    cp godot/bin/godot.linuxbsd.editor.$ARCH $dstBinPath"_linux_prod"
else
    cp godot/bin/godot.macos.editor.$ARCH $dstBinPath"_darwin_prod"
fi