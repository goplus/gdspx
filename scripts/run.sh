#!/bin/bash

extra_init=""

while getopts ":t:i:e" opt; do
  case ${opt} in
    t )
      cmdname=$OPTARG
      ;;
    e )
      extra_init="true"
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

if [ "$extra_init" == "true" ]; then
    make initload 
    make run
else
    make $cmdname
fi
cd ../../