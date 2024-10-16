#!/bin/bash
# build godot web version

# make sure emsdk's version is same as godot build system
godot_version="4.2.2.stable"

# Check if the 'emsdk' directory exists
if [ ! -d "emsdk" ]; then
    echo "emsdk not found, install emsdk first"
    git clone https://github.com/emscripten-core/emsdk.git
    cd emsdk
    ./emsdk install 3.1.39
    ./emsdk activate 3.1.39
    # Check the operating system type
    if [ "$OS" = "Windows_NT" ]; then
        # call emsdk_env.bat
        source ./emsdk_env.sh
        # setx PATH "%cd%\upstream\emscripten;%PATH%"
        export PATH="$(pwd)/upstream/emscripten:$PATH"
    else
        source ./emsdk_env.sh
    fi
    cd ..
fi

em++ --version
# build godot
cd godot

# gdspx disable gdextension
#scons platform=web target=template_debug threads=no dlink_enabled=yes 
#mv bin/godot.web.template_debug.wasm32.dlink.zip bin/web_dlink_debug.zip
# Get the path to the AppData\Roaming directory
if [ "$OS" = "Windows_NT" ]; then
    appdata_path="c:$(echo $APPDATA | sed 's/\\/\//g' | sed 's/C://g')"
    gd_template_path="${appdata_path}/Godot/export_templates/${godot_version}"
else
    gd_template_path=~/.local/share/godot/export_templates/$godot_version
fi

# build web export templates
scons platform=web target=template_debug threads=no
mv bin/godot.web.template_debug.wasm32.zip bin/web_dlink_debug.zip
rm $gd_template_path/web_*.zip
cp bin/web_dlink_debug.zip $gd_template_path/web_dlink_debug.zip
cp bin/web_dlink_debug.zip $gd_template_path/web_dlink_release.zip
cp bin/web_dlink_debug.zip $gd_template_path/web_debug.zip
cp bin/web_dlink_debug.zip $gd_template_path/web_release.zip


cd ..