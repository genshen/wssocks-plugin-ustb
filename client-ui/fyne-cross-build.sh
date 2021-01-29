#!/bin/sh
# build app using fyne-cross in docker

# How to run: go to root dir of tthis repo, run `./client-ui/fyne-cross-build.sh`

appBuildNumber=2
appID="wssocks-ustb-client-ui.genshen.github.com"
appVersion=0.6.0
appIcon=./client-ui/app-512.png
buildEnv="GOPROXY=https://goproxy.cn"
buildFlags="--app-build ${appBuildNumber} --app-id ${appID} --app-version ${appVersion} --icon ${appIcon} --env ${buildEnv}"

fyne-cross linux ${buildFlags} --arch amd64 ./client-ui
fyne-cross windows ${buildFlags} --arch amd64 ./client-ui
fyne-cross darwin ${buildFlags} --arch amd64 ./client-ui
