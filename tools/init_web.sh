#!/bin/bash

git clone https://github.com/emscripten-core/emsdk.git
cd emsdk

./emsdk install 3.1.39
./emsdk activate 3.1.39

source ./emsdk_env.sh