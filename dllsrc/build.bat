set GOROOT=G:\Go\go1.13
set GOPATH=E:\gopath
SET PATH=%PATH%;%GOROOT%\bin\;G:\Program Files (x86)\Dev-Cpp\MinGW64\bin
SET CGO_ENABLED=1
SET GOOS=windows
SET GOARCH=386
SET CC=x86_64-w64-mingw32-gcc.exe

SET GO111MODULE=on
set GOPROXY=https://goproxy.cn,direct
%GOROOT%\bin\go.exe build -work -buildmode=c-shared -ldflags "-s -w -extldflags '-static'" -o libtextractor.dll libtextractor.go

pause


