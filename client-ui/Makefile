# makefile for crossing building on macOS
# supported os: windows, macOS

PHONY: clean all package

PACKAGE=github.com/genshen/wssocks-plugin-ustb/client-ui
FLAGS=-gcflags="all=-trimpath=${GOPATH}" -asmflags "all=-trimpath=${GOPATH}"
appID=wssocks-ustb-client-ui.genshen.github.com
appName="wssocks-ustb client"
APP_VERSION=0.5.1
PACKAGE_FLAG=-icon app-512.png -appID ${appID} --appVersion ${APP_VERSION} -release -appBuild 4

all: wssocks-ustb-ui-macOS-amd64 wssocks-ustb-ui-windows-amd64.exe # wssocks-ustb-ui-linux-amd64

wssocks-ustb-ui-macOS-amd64:
	CC=clang CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build ${FLAGS} -o wssocks-ustb-ui-macOS-amd64 ${PACKAGE}

wssocks-ustb-ui-windows-amd64.exe:
	CC=x86_64-w64-mingw32-gcc CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build ${FLAGS} -ldflags -H=windowsgui -o wssocks-ustb-ui-windows-amd64.exe ${PACKAGE}

# wssocks-ustb-ui-linux-amd64:
# 	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o wssocks-ustb-ui-linux-amd64 ${PACKAGE}

# packaging
package: client-ui-macOS-amd64.app client-ui-windows-amd64.exe # client-ui-linux-amd64

client-ui-macOS-amd64.app: wssocks-ustb-ui-macOS-amd64
	cp -r wssocks-ustb-ui-macOS-amd64 wssocks-ustb-client
	CC=clang CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 fyne package -os darwin ${PACKAGE_FLAG} -executable wssocks-ustb-client -name ${appName}
	rm -rf client-ui-macOS-amd64.app; mv ${appName}.app client-ui-macOS-amd64.app

client-ui-windows-amd64.exe: wssocks-ustb-ui-windows-amd64.exe
	CC=x86_64-w64-mingw32-gcc CGO_ENABLED=1 GOOS=windows GOARCH=amd64 fyne package -os windows ${PACKAGE_FLAG} -executable wssocks-ustb-ui-windows-amd64.exe -name ${appName}.exe
	mv client-ui.exe client-ui-windows-amd64.exe

# export CGO_CFLAGS='-D _POSIX'

# client-ui-linux-amd64:

clean:
	rm -rf wssocks-ustb-ui-macOS-amd64 wssocks-ustb-ui-windows-amd64.exe client-ui-macOS-amd64.app client-ui-windows-amd64.exe
