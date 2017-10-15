go get golang.org/x/sys/unix
set GOARCH=amd64

REM linux
set GOOS=linux
go build 
if not exist "dist" mkdir dist
cd dist
mkdir linux
cd linux
move ..\..\goark-node .
mkdir cfg
copy ..\..\cfg\*.* cfg
mkdir logs
mkdir db
cd ..
cd ..


REM windows
set GOOS=windows
go build 
if not exist "dist" mkdir dist
cd dist
if not exist "windows" mkdir windows
cd windows
move ..\..\goark-node.exe .
mkdir cfg
copy ..\..\cfg\*.* cfg
mkdir logs
mkdir db
cd ..
cd ..

REM darwin
set GOOS=darwin
go build 
if not exist "dist" mkdir dist
cd dist
if not exist "darwin" mkdir darwin
cd darwin
move ..\..\goark-node .
mkdir cfg
copy ..\..\cfg\*.* cfg
mkdir logs
mkdir db
cd ..

REM create archive
for /d %%X in (*) do "c:\Program Files\7-Zip\7z.exe" a -mx "%%X.zip" "%%X\*"

move linux.zip GOARKNodeLinuxRelease_x64.zip
move windows.zip GOARKNodeWindowsRelease_x64.zip
move darwin.zip GOARKNodeDarwinRelease_x64.zip


