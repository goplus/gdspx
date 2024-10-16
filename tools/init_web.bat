
@echo off
setlocal

REM Clone the emsdk repository
git clone https://github.com/emscripten-core/emsdk.git
cd emsdk

call emsdk install 3.1.39
call emsdk activate 3.1.39

REM Set up the environment
call emsdk_env.bat
setx PATH "%cd%\upstream\emscripten;%PATH%"

endlocal
