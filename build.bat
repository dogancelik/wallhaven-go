@echo off

set "build=%~dp0build"

mkdir %build%
cd /D %build%
del /Q *

go build ../cli
ren cli.exe wallhaven.exe
for /F "tokens=2 delims= " %%i in ('wallhaven.exe -v') do set "version=%%i"

copy ..\cli\cli.json wallhaven.json
7za a -tzip -i!* %version%.zip
cd ..