# makefile for crossing building on macOS
# supported os: windows, macOS

# make sure fyne is installed before packaging (see: https://developer.fyne.io/started/packaging).

PHONY: clean all package

PACKAGE=github.com/genshen/wssocks-plugin-ustb/client-ui
FLAGS=--trimpath
appID=wssocks-ustb-client-ui.genshen.github.com
appName="wssocks-ustb client"
APP_VERSION=0.8.0-beta  # use appVersion in v0.8.0 release
PACKAGE_FLAG=-icon app-512.png -appID ${appID} -release -appBuild 5

all: wssocks-ustb-ui-macOS-amd64 wssocks-ustb-ui-macOS-arm64 wssocks-ustb-ui-windows-amd64.exe # wssocks-ustb-ui-linux-amd64

wssocks-ustb-ui-macOS-amd64:
	CC=clang CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build ${FLAGS} -o wssocks-ustb-ui-macOS-amd64 ${PACKAGE}

wssocks-ustb-ui-macOS-arm64:
	CC=clang CGO_ENABLED=1 GOOS=darwin GOARCH=arm64 go build ${FLAGS} -o wssocks-ustb-ui-macOS-arm64 ${PACKAGE}

wssocks-ustb-ui-windows-amd64.exe:
	CC=x86_64-w64-mingw32-gcc CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build ${FLAGS} -ldflags -H=windowsgui -o wssocks-ustb-ui-windows-amd64.exe ${PACKAGE}

wssocks-ustb-ui-macOS-universal: wssocks-ustb-ui-macOS-amd64 wssocks-ustb-ui-macOS-arm64
	lipo wssocks-ustb-ui-macOS-amd64 wssocks-ustb-ui-macOS-arm64 -create -output wssocks-ustb-ui-macOS-universal

# wssocks-ustb-ui-linux-amd64:
# 	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o wssocks-ustb-ui-linux-amd64 ${PACKAGE}

# packaging
package: client-ui-macOS-universal.app client-ui-windows-amd64.exe # client-ui-linux-amd64

client-ui-macOS-universal.app: wssocks-ustb-ui-macOS-universal
	CC=clang CGO_ENABLED=1 GOOS=darwin fyne package -os darwin ${PACKAGE_FLAG} -executable wssocks-ustb-ui-macOS-universal -name ${appName}
	rm -rf client-ui-macOS-universal.app; mv ${appName}.app client-ui-macOS-universal.app

client-ui-windows-amd64.exe: wssocks-ustb-ui-windows-amd64.exe
	CC=x86_64-w64-mingw32-gcc CGO_ENABLED=1 GOOS=windows GOARCH=amd64 fyne package -os windows ${PACKAGE_FLAG} -executable wssocks-ustb-ui-windows-amd64.exe -name ${appName}.exe
	mv ${appName}.exe client-ui-windows-amd64.exe

# export CGO_CFLAGS='-D _POSIX'

# client-ui-linux-amd64:

clean:
	rm -rf wssocks-ustb-ui-macOS-amd64 wssocks-ustb-ui-macOS-arm64 wssocks-ustb-ui-macOS-universal wssocks-ustb-ui-windows-amd64.exe client-ui-macOS-universal.app client-ui-windows-amd64.exe
	rm -rf wssocks-ustb-client fyne.syso
