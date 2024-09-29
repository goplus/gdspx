# godot for spx


## Getting Started

### start and run 
```
# install env
    git clone git@github.com:realdream-ai/gdspx.git
    cd gdspx/cmd/gdspx/
    go install .
    cd ../../

# run demo
    gdspx run tutorial/01_aircraft  
```

### how to use
Usage:

    gdspx <command> [path]      

The commands are:

    - help            # Show help info
    - init            # Create a gdspx project in the current directory
    - run             # Run the current project
    - editor          # Open the current project in editor mode
    - build           # Build the dynamic library
    - export          # Export the PC package (macOS, Windows, Linux) (TODO)
    - buildweb        # Build for WebAssembly (WASM)
    - exportweb       # Export the web package

 eg:

    gdspx init                      # create a project in current path
    gdspx init ./test/demo01        # create a project at path ./test/demo01 

