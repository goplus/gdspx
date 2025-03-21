#!/usr/bin/env bash

basedir=$(cd $(dirname "$0"); pwd)
basedir=$basedir/docker
cd $basedir
podman=`command -v podman`

if [ -z "$podman" ]; then
  echo "podman needs to be in PATH for this script to work."
  exit 1
fi

if [ -z "$1" ]; then
  echo "Usage: $0 <proxy_url>"
  echo
  echo "Example: $0 http://192.168.31.147:7890"
  echo
  echo "proxy_url:"
  echo "        Proxy URL for network access (e.g., http://192.168.31.147:7890)"
  echo
  echo "The resulting image version will be 4.x-f36."
  exit 1
fi

PROXY_URL=$1
godot_branch="4.x"
base_distro="f36"
img_version=$godot_branch-$base_distro
files_root="$basedir/files"

if [ ! -z "$PS1" ]; then
  # Confirm settings
  echo "Docker image tag: ${img_version}"
  echo
  while true; do
    read -p "Is this correct? [y/n] " yn
    case $yn in
      [Yy]* ) break;;
      [Nn]* ) exit 1;;
      * ) echo "Please answer yes or no.";;
    esac
  done
fi

mkdir -p logs

# Add proxy settings
export HTTP_PROXY="$PROXY_URL"
export HTTPS_PROXY="$PROXY_URL"

"$podman" build \
  --build-arg HTTP_PROXY="$PROXY_URL" \
  --build-arg HTTPS_PROXY="$PROXY_URL" \
  -t godot-fedora:${img_version} -f Dockerfile.base . 2>&1 | tee logs/base.log

podman_build() {
  # You can add --no-cache as an option to podman_build below to rebuild all containers from scratch.
  "$podman" build \
    --build-arg HTTP_PROXY="$PROXY_URL" \
    --build-arg HTTPS_PROXY="$PROXY_URL" \
    --build-arg img_version=${img_version} \
    -v "${files_root}":/root/files:z \
    -t godot-"$1:${img_version}" \
    -f Dockerfile."$1" . \
    2>&1 | tee logs/"$1".log
}

podman_build linux
exit 0
podman_build windows
podman_build web
podman_build android

XCODE_SDK=15
OSX_SDK=14.0
IOS_SDK=17.0
if [ ! -e "${files_root}"/MacOSX${OSX_SDK}.sdk.tar.xz ] || [ ! -e "${files_root}"/iPhoneOS${IOS_SDK}.sdk.tar.xz ] || [ ! -e "${files_root}"/iPhoneSimulator${IOS_SDK}.sdk.tar.xz ]; then
  if [ ! -e "${files_root}"/Xcode_${XCODE_SDK}.xip ]; then
    echo "files/Xcode_${XCODE_SDK}.xip is required. It can be downloaded from https://developer.apple.com/download/more/ with a valid apple ID."
    exit 1
  fi

  echo "Building OSX and iOS SDK packages. This will take a while"
  podman_build xcode
  "$podman" run -it --rm \
    -e HTTP_PROXY="$PROXY_URL" \
    -e HTTPS_PROXY="$PROXY_URL" \
    -v "${files_root}":/root/files:z \
    -e XCODE_SDKV="${XCODE_SDK}" \
    -e OSX_SDKV="${OSX_SDK}" \
    -e IOS_SDKV="${IOS_SDK}" \
    godot-xcode:${img_version} \
    2>&1 | tee logs/xcode_packer.log
fi

podman_build osx
podman_build ios
