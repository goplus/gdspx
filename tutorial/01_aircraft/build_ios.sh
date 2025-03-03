#!/bin/bash

set -e

# Configuration variables
GO_SRC_DIR="./go"
LIB_DIR="./lib"
FRAMEWORK_NAME="gdspx"
XCFRAMEWORK_PATH="$LIB_DIR/lib$FRAMEWORK_NAME.ios.xcframework"
BUILD_DIR=".godot/tmp/gobuild"  # 使用.godot/tmp/gobuild作为中间结果目录
SIMULATOR_DIR="$BUILD_DIR/simulator"
DEVICE_DIR="$BUILD_DIR/device"
HEADERS_DIR="$BUILD_DIR/headers"

# Create directories
rm -rf "$BUILD_DIR" "$XCFRAMEWORK_PATH"
mkdir -p "$SIMULATOR_DIR" "$DEVICE_DIR" "$LIB_DIR" "$HEADERS_DIR"

echo "📦 Building Go libraries for iOS..."

# Create a dummy header file with the required exports
cat > "$HEADERS_DIR/libgdspx.h" << EOF
#ifndef LIBGDSPX_H
#define LIBGDSPX_H

#include <stdlib.h>

// GDExtension initialization function
void GDExtensionInit(void *p_interface, const void *p_library, void *r_initialization);

#endif // LIBGDSPX_H
EOF

# Copy our C headers to the headers directory
cp "$GO_SRC_DIR"/*.h "$HEADERS_DIR/"

cd "$GO_SRC_DIR"

# Get SDK paths
SIMULATOR_SDK_PATH=$(xcrun --sdk iphonesimulator --show-sdk-path)
DEVICE_SDK_PATH=$(xcrun --sdk iphoneos --show-sdk-path)

# Disable signal handling in Go for iOS
export GODEBUG=cgocheck=0,asyncpreemptoff=1,panicnil=1

# Build for iOS Simulator (x86_64)
echo "🔨 Building for iOS Simulator (x86_64)..."
CGO_ENABLED=1 \
GOOS=darwin \
GOARCH=amd64 \
CGO_CFLAGS="-isysroot $SIMULATOR_SDK_PATH -mios-simulator-version-min=12.0 -arch x86_64 -fembed-bitcode" \
CGO_LDFLAGS="-isysroot $SIMULATOR_SDK_PATH -mios-simulator-version-min=12.0 -arch x86_64" \
go build -tags=ios -buildmode=c-archive -trimpath -ldflags="-w -s" -o "../$SIMULATOR_DIR/libgdspx-x86_64.a" .

# Build for iOS Simulator (arm64)
echo "🔨 Building for iOS Simulator (arm64)..."
CGO_ENABLED=1 \
GOOS=darwin \
GOARCH=arm64 \
CGO_CFLAGS="-isysroot $SIMULATOR_SDK_PATH -mios-simulator-version-min=12.0 -arch arm64 -fembed-bitcode" \
CGO_LDFLAGS="-isysroot $SIMULATOR_SDK_PATH -mios-simulator-version-min=12.0 -arch arm64" \
go build -tags=ios -buildmode=c-archive -trimpath -ldflags="-w -s" -o "../$SIMULATOR_DIR/libgdspx-arm64-sim.a" .

# Build for iOS Device (arm64)
echo "🔨 Building for iOS Device (arm64)..."
CGO_ENABLED=1 \
GOOS=darwin \
GOARCH=arm64 \
CGO_CFLAGS="-isysroot $DEVICE_SDK_PATH -mios-version-min=12.0 -arch arm64 -fembed-bitcode" \
CGO_LDFLAGS="-isysroot $DEVICE_SDK_PATH -mios-version-min=12.0 -arch arm64" \
go build -tags=ios -buildmode=c-archive -trimpath -ldflags="-w -s" -o "../$DEVICE_DIR/libgdspx-arm64.a" .

cd ..

# Create a fat binary for simulator (combines arm64 and x86_64)
echo "🔗 Creating fat binary for simulator..."
lipo -create -output "$SIMULATOR_DIR/libgdspx.a" \
  "$SIMULATOR_DIR/libgdspx-x86_64.a" \
  "$SIMULATOR_DIR/libgdspx-arm64-sim.a"

# Create XCFramework
echo "🎁 Creating XCFramework..."
xcrun xcodebuild -create-xcframework \
  -library "$SIMULATOR_DIR/libgdspx.a" -headers "$HEADERS_DIR" \
  -library "$DEVICE_DIR/libgdspx-arm64.a" -headers "$HEADERS_DIR" \
  -output "$XCFRAMEWORK_PATH"

# 删除中间结果目录
echo "🧹 Cleaning up temporary build files..."
rm -rf "$BUILD_DIR"

echo "✅ Successfully built libgdspx.ios.xcframework!"
echo "📍 Location: $XCFRAMEWORK_PATH"
