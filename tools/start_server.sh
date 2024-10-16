#!/bin/bash

if [ "$OS" = "Windows_NT" ]; then
    echo -e "taskkill /F /IM python.exe\ntaskkill /F /IM pythonw.exe" > temp_kill.bat
    cmd.exe /C temp_kill.bat
    rm temp_kill.bat
else 
    pkill -f gdspx_web_server.py
fi

PROJ_PATH=$1
echo "gdgo server running at http://localhost:8005"
python ./tools/gdspx_web_server.py -r ../$PROJ_PATH/.builds/web -p 8005  &
