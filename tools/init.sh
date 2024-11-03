#!/bin/bash
pip install scons==4.7.0
pip install ninja 

if [ ! -d "godot" ]; then
    echo "Godot directory not found. Creating and initializing..."
    mkdir godot
    cd godot
    git init 
    git remote add origin git@github.com:JiepengTan/godot.git
    git fetch --depth 1 origin spx-4.2.2
    git checkout spx-4.2.2
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
gopath=${GOPATH:-$(go env GOPATH)}
dstBinPath="$gopath/bin/gd4spx4.2.2"
echo "Destination binary path: $dstBinPath"
if [ "$OS" = "Windows_NT" ]; then
    cp godot/bin/godot.windows.editor.dev.x86_64 $dstBinPath
elif [ "$OS" = "Linux" ]; then
    cp godot/bin/godot.linuxbsd.editor.dev.x86_64 $dstBinPath
else
    cp godot/bin/godot.macos.editor.dev.x86_64 $dstBinPath
fi