# godot for spx

## prepare env
1. install "make" in your system 
2. clone repos 
    ```
    # clone godot
    git clone git@github.com:JiepengTan/godot.git
    cd godot
    git checkout -t origin/spx-4.2.2
    scons target=editor
    cd ..

    # clone gdspx
    git clone git@github.com:realdream-ai/gdspx.git
    cd gdspx
    ```
3. make cmd "godot" work
    - windows
        1. rename ../bin/godot.windows.editor.x86_64.exe to ../bin/godot.exe  
        2. add ../bin/godot.exe to system path
    - linux | macos
       ```
        sudo mv ../bin/godot*.editor /usr/local/bin/godot
       ```

## run demo
Use the "WSAD" keys to move the player

<p align="center"><img src="docs\pics\01_aircraft.png?raw=true" width="128"></p> 

### windows
```
.\scripts\run.bat
```

### linux | macos
```
./scripts/run.sh
```