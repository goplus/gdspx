#!/bin/bash

pip install scons
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
scons target=editor arch=x86_64
cd ..

echo "init and run demo project."
./scripts/run.sh
