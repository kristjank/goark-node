REM linux
set GOOS=linux
set GOARCH=amd64
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
set GOARCH=amd64
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
set GOARCH=amd64
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
cd ..

REM linux/arm
set GOOS=linux
set GOARCH=arm
go build
if not exist "dist" mkdir dist
cd dist
if not exist "linuxarm" mkdir linuxarm
cd linuxarm
move ..\..\goark-node .
mkdir cfg
copy ..\..\cfg\*.* cfg
mkdir logs
mkdir db
cd ..
cd .. 

REM linux/arm
set GOOS=linux
set GOARCH=arm
set GOARM=5
go build
if not exist "dist" mkdir dist
cd dist
if not exist "linuxarm5" mkdir linuxarm5
cd linuxarm
move ..\..\goark-node .
mkdir cfg
copy ..\..\cfg\*.* cfg
mkdir logs
mkdir db
cd ..
cd .. 

REM linux/arm
set GOOS=linux
set GOARCH=arm6
go build
if not exist "dist" mkdir dist
cd dist
if not exist "linuxarm6" mkdir linuxarm6
cd linuxarm
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
move linuxarm.zip GOARKNodeLinuxRelease_ARM.zip
move linuxarm5.zip GOARKNodeLinuxRelease_ARM5.zip
move linuxarm6.zip GOARKNodeLinuxRelease_ARM6.zip

