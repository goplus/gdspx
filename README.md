# godot for spx


## Quick start
### Prerequisite Environment Setup
- Install go (version >= 1.22.3)
- Install python (version >= 3.8)
- Install make
- Add the Go environment's bin directory to the system PATH.
```
    export PATH=$PATH:$GOPATH/bin
```


### quick start 
```
    git clone git@github.com:realdream-ai/gdspx.git
    cd gdspx
    make pc
    gdspx run tutorial/01_aircraft
```

### how to use
Usage:

    gdspx <command> [path]      

The commands are:

    - help            # Show help info
    - init            # Create a gdspx project in the current directory
    - editor          # Open the current project in editor mode
    - build           # Build the dynamic library
    - run             # Run the current project
    - export          # Export the PC package (macOS, Windows, Linux) (TODO)
    - buildweb        # Build for WebAssembly (WASM)
    - runweb          # Run the current project in browser
    - exportweb       # Export the web package

 eg:

    gdspx init                      # create a project in current path
    gdspx init ./test/demo01        # create a project at path ./test/demo01 


### how to develop (Need run on the root dir)

Run the following commands in the **root** directory:

    make <command>

The commands are:

    - pc             # Rebuild engine from source for PC (Windows, macOS, Linux)
    - web            # Rebuild engine from source for WEB (Editor mode)
    - webpack        # Rebuild engine from source for WEB (Game mode)
    - fmt            # Format current project's code
    - gen            # Generate the engine wrap files (use this when the engine is updated)
  
 eg:

    make gen       # Update engine warp codes
    make pc        # Rebuild the engine