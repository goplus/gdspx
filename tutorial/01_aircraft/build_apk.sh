#!/bin/bash
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd $SCRIPT_DIR

# Default project directory
DEFAULT_PROJ_DIR=.

# Parse command line arguments
while getopts "p:" opt; do
    case $opt in
        p) PROJ_DIR="$OPTARG";;
        \?) echo "Invalid option: -$OPTARG" >&2; exit 1;;
    esac
done

# Use default if not specified
PROJ_DIR=${PROJ_DIR:-$DEFAULT_PROJ_DIR}

# Convert PROJ_DIR to absolute path
PROJ_DIR="$(cd "$PROJ_DIR" && pwd)"

# Check if ANDROID_NDK_ROOT is set
if [ -z "$ANDROID_NDK_ROOT" ]; then
    echo "Error: ANDROID_NDK_ROOT environment variable is not set"
    echo "Please set it to your Android NDK installation path"
    exit 1
fi

LIB_DIR=$PROJ_DIR/lib
GO_DIR=$PROJ_DIR/go
cd $GO_DIR

# Set common variables
# Detect system architecture and OS
OS_NAME="$(uname -s)"
ARCH="$(uname -m)"

if [[ "$OS_NAME" =~ "MINGW"|"MSYS"|"CYGWIN" ]]; then
    # Windows environment (MSYS2/MinGW/Cygwin)
    case "$ARCH" in
        x86_64|amd64) HOST_TAG="windows-x86_64";;
        *)            echo "Unsupported Windows architecture: $ARCH"; exit 1;;
    esac
elif [[ "$OS_NAME" == "Linux" ]]; then
    case "$ARCH" in
        x86_64)  HOST_TAG="linux-x86_64";;
        aarch64) HOST_TAG="linux-aarch64";;
        *)       echo "Unsupported Linux architecture: $ARCH"; exit 1;;
    esac
elif [[ "$OS_NAME" == "Darwin" ]]; then
    HOST_TAG="darwin-x86_64"
else
    echo "Unsupported operating system: $OS_NAME"
    exit 1
fi

NDK_TOOLCHAIN="$ANDROID_NDK_ROOT/toolchains/llvm/prebuilt/$HOST_TAG/bin"
MIN_SDK=21
echo "Building for arm64-v8a..."
CGO_ENABLED=1 \
GOOS=android \
GOARCH=arm64 \
CC="$NDK_TOOLCHAIN/aarch64-linux-android$MIN_SDK-clang" \
go build -buildmode=c-shared -o $LIB_DIR/libgdspx-android-arm64.so .

echo "Building for armeabi-v7a..."
CGO_ENABLED=1 \
GOOS=android \
GOARCH=arm \
CC="$NDK_TOOLCHAIN/armv7a-linux-androideabi$MIN_SDK-clang" \
go build -buildmode=c-shared -o $LIB_DIR/libgdspx-android-arm32.so .

echo "Build android so completed successfully!"

# Paths
ENGINE_BINARY="/Users/tjp/go/bin/gdspx2.0.1_darwin"
PROJECT_PATH="$PROJ_DIR/project.godot"
APK_PATH="$PROJ_DIR/builds/test.apk"
BUILD_DIR=$(dirname "$APK_PATH")

# Check if ENGINE_BINARY environment variable is set
if [ -z "$ENGINE_BINARY" ]; then
    echo "Error: ENGINE_BINARY environment variable is not set"
    echo "Please set it to your Godot binary path, e.g.:"
    echo "export ENGINE_BINARY=/path/to/godot"
    exit 1
fi



# Create builds directory if it doesn't exist
mkdir -p "$BUILD_DIR"

# Check if Godot binary exists
if [ ! -f "$ENGINE_BINARY" ]; then
    echo "Error: Godot binary not found at $ENGINE_BINARY"
    exit 1
fi

# Check if project file exists
if [ ! -f "$PROJECT_PATH" ]; then
    echo "Error: Godot project file not found at $PROJECT_PATH"
    exit 1
fi

# Import project to ensure resources are up to date
echo "Importing project resources..."
"$ENGINE_BINARY" --headless --path "$(dirname "$PROJECT_PATH")" --editor --quit

# Export the project to APK
echo "Exporting Godot project to APK..."
"$ENGINE_BINARY" --headless --path "$(dirname "$PROJECT_PATH")" --export-debug "Android" "$APK_PATH"

if [ ! -f "$APK_PATH" ]; then
    echo "Error: APK export failed"
    exit 1
fi

# Check if adb is available
if ! command -v adb &> /dev/null; then
    echo "Error: adb command not found. Please ensure Android SDK platform tools are installed and in your PATH"
    exit 1
fi

# Check if any Android device is connected
if ! adb devices | grep -q "device$"; then
    echo "Error: No Android device connected. Please connect a device and enable USB debugging"
    exit 1
fi

echo "Installing APK..."
adb install -r "$APK_PATH"

if [ $? -eq 0 ]; then
    echo "APK installation successful!"
else
    echo "Error: APK installation failed"
    exit 1
fi

cd -
