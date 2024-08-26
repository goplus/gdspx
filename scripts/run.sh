#!/bin/bash
cd tutorial/01_aircraft
pwd
if [ "$1" == "--init" ]; then
    make initload 
    make run
else
    make run
fi
cd ../../