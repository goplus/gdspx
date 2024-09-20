#!/bin/bash
pkill -f gdspx_web_server

echo "gdgo server running at http://localhost:8005"
python ./tools/gdspx_web_server.py -r ../builds/games -p 8005  &
