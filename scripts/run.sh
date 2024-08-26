#!/bin/bash

while getopts ":t:i:" opt; do
  case ${opt} in
    t )
      cmdname=$OPTARG
      ;;
    i )
      dstdir=$OPTARG
      ;;
    \? )
      echo "Invalid option: $OPTARG" 1>&2
      ;;
    : )
      echo "Option $OPTARG requires an argument" 1>&2
      ;;
  esac
done

shift $((OPTIND -1))

if [ -z "$dstdir" ]; then
    dstdir=tutorial/01_aircraft
fi

if [ -z "$cmdname" ]; then
    cmdname="run"
fi

cp Makefile $dstdir
cd $dstdir

if [ ! -d ".godot" ]; then
    make initload 
fi
make $cmdname
cd ../../