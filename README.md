# godot for spx

## prepare env
1. install "**make**" for build
2. install python
3. run script 
    ```
    ./scripts/init.sh
    ```


## run demo 
```
./scripts/run.sh -i tutorial/01_aircraft
```

Use the "WSAD" keys to move the player

<p align="center"><img src="docs\pics\01_aircraft.png?raw=true" width="128"></p> 



## make cmd
eg:  make pc 

- server        : launch a local web server (http://127.0.0.1:8005/)
- pc            : run on pc (win, mac, linux)
- web           : build and export web package to dir (./build/games )
- fmt           : format the code
- init          : init the environment
- wasm          : build web wasm 
- editor        : run in editor mode



```
	go install realdream-ai/gdspx/cmd/gdspx@main
```