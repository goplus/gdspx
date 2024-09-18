#!/bin/bash
# build godot web version

# make sure emsdk's version is same as godot build system
em++ --version
godot_version="4.2.2.stable"

# build godot
echo "================ build_godot  ====================="
cd godot
# gdspx disable gdextension
#scons platform=web target=template_debug threads=no dlink_enabled=yes 
#mv bin/godot.web.template_debug.wasm32.dlink.zip bin/web_dlink_debug.zip

# build web export templates
scons platform=web target=template_debug threads=no
mv bin/godot.web.template_debug.wasm32.zip bin/web_dlink_debug.zip
rm ~/.local/share/godot/export_templates/$godot_version/web_*.zip
cp bin/web_dlink_debug.zip ~/.local/share/godot/export_templates/$godot_version/web_dlink_debug.zip
cp bin/web_dlink_debug.zip ~/.local/share/godot/export_templates/$godot_version/web_dlink_release.zip
cp bin/web_dlink_debug.zip ~/.local/share/godot/export_templates/$godot_version/web_debug.zip
cp bin/web_dlink_debug.zip ~/.local/share/godot/export_templates/$godot_version/web_release.zip

# build editor
scons target=editor arch=x86_64 dev_build=yes


cd ..