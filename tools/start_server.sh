#!/bin/bash
pkill -f http.server
pkill -f godot_web_server

echo "gdgo server running at http://localhost:8005"
python ./tools/godot_web_server.py -r ../builds/web_go -p 8005  &
