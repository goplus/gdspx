# godot for spx


## Getting Started (WIP, Linux only for now)

### install cmd
```
	go install realdream-ai/gdspx/cmd/gdspx@main
```

### how to use
    
    - gdspx init            # Create a gdspx project in the current directory
    - gdspx run             # Run the current project
    - gdspx build           # Build the dynamic library
    - gdspx export          # Export the PC package (supports macOS, Windows, Linux)
    - gdspx buildweb        # Build for WebAssembly (WASM)
    - gdspx exportweb       # Export the web package




## develop 
### prepare env
1. install "**make**" for build
2. install python
3. run script 
    ```
    ./scripts/init.sh
    ```


### run demo 
```
./scripts/run.sh -i tutorial/01_aircraft
```

Use the "WSAD" keys to move the player

<p align="center"><img src="docs\pics\01_aircraft.png?raw=true" width="128"></p> 



### make cmd
eg:  make pc 

    - make server        # Launch a local web server (http://127.0.0.1:8005/)
    - make build         # Build DLL for PC (Windows, macOS, Linux)
    - make run           # Run on PC (Windows, macOS, Linux)
    - make web           # Build and export the web package to the directory (./build/games)
    - make fmt           # Format the code
    - make init          # Initialize the environment
    - make wasm          # Build for WebAssembly (WASM)
    - make editor        # Run in editor mode


