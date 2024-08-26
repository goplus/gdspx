#!/bin/bash
dstdir=tutorial/01_aircraft
cp Makefile $dstdir
cd $dstdir
pwd
if [ "$1" == "--init" ]; then
    make initload 
    make run
else
    make run
fi
cd ../../