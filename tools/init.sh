#!/bin/bash

# install gdspx
cd cmd/gdspx/
go install .
cd ../../

GOPATH=$(go env GOPATH)
VERSION=$(cat ./cmd/gdspx/pkg/impl/template/version)
echo "version="$VERSION " GOPATH=" $GOPATH


pip install scons==4.7.0
if [ ! -d "godot" ]; then
    echo "Godot directory not found. Creating and initializing..."
    mkdir godot
    cd godot
    git init 
    git remote add origin https://github.com/realdream-ai/godot.git
    git fetch --depth 1 origin spx4.2.2
    git checkout spx4.2.2
    echo "Godot repository setup complete."
else
    cd godot
    echo "Godot directory already exists."
fi
#  dev_build=yes

if [ "$OS" = "Windows_NT" ]; then
    scons target=editor arch=x86_64 vsproj=yes dev_build=yes
    echo '"godot.windows.editor.dev.x86_64.exe" %*' > bin/godot.bat
else
    scons target=editor arch=x86_64 dev_build=yes
fi
cd ..

echo "init engine done."
dstBinPath="$GOPATH/bin/gdspx$VERSION"
echo "Destination binary path: $dstBinPath"
if [ "$OS" = "Windows_NT" ]; then
    cp godot/bin/godot.windows.editor.dev.x86_64 $dstBinPath"_win.exe"
elif [ "$OS" = "Linux" ]; then
    cp godot/bin/godot.linuxbsd.editor.dev.x86_64 $dstBinPath"_linux"
else
    cp godot/bin/godot.macos.editor.dev.x86_64 $dstBinPath"_darwin"
fi