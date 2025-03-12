# godot for spx


## Quick start
### 1. Prerequisite Environment Setup
- Install go (version >= 1.22.3)
- Install python (version >= 3.8)
- Install make
- Add the Go environment's bin directory to the system PATH.
```
    export PATH=$PATH:$GOPATH/bin
```


### 2. Quick start 
```
    git clone git@github.com:realdream-ai/gdspx.git
    cd gdspx
    make pc
    gdspx run tutorial/01_aircraft
```

### 3. How to use
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
    - exportapk       # Export the Android package
    - exportios       # Export the iOS package

 eg:

    gdspx init                      # create a project in current path
    gdspx init ./test/demo01        # create a project at path ./test/demo01 


### 4. How to develop (Need run on the root dir)

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



### 5. Setup build environment for android
- refer to [build engine](https://docs.godotengine.org/en/4.2/contributing/development/compiling/compiling_for_android.html) , [Deploying](https://docs.godotengine.org/en/4.2/tutorials/export/exporting_for_android.html)

1. Installing OpenJDK 17
2. install android sdk (AndoridStudio or command-line tools)
3. install ndk: 23.2.8568313
4. export ANDROID_NDK_ROOT, ANDROID_HOME
    eg:
    ```
    export ANDROID_SDK_ROOT=$HOME/Library/Android/sdk
    export ANDROID_NDK_ROOT=$HOME/Library/Android/sdk/ndk/23.2.8568313
    export PATH="${ANDROID_SDK_ROOT}/platform-tools:${ANDROID_NDK_ROOT}:${PATH}"
    ```
5. Creating a debug.keystore
6. Configuring the location of the Android SDK and debug.keystore in Godot


### 6. Setup build environment for ios
- refer to [build engine](https://docs.godotengine.org/en/4.2/contributing/development/compiling/compiling_for_ios.html), [Deploying](https://docs.godotengine.org/en/4.2/tutorials/export/exporting_for_ios.html#doc-exporting-for-ios)


