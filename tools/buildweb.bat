cd ..
call scons platform=web target=editor
cd test
copy  /Y ..\bin\godot.web.editor.wasm32.zip .
powershell -command "Expand-Archive -Path .\godot.web.editor.wasm32.zip -DestinationPath . -Force"

run.bat %1