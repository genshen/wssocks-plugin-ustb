#!/bin/sh
# build app using fyne-cross in docker

# How to run: go to root dir of this repo, run `./client-ui/fyne-cross-build.sh`

appBuildNumber=5
appID="wssocks-ustb-client-ui.genshen.github.com"
appVersion=0.8.0-beta # use appVersion in v0.8.0 release
appIcon=./client-ui/app-512.png
buildEnv="GOPROXY=https://goproxy.cn"
buildFlags="--app-build ${appBuildNumber} --app-id ${appID} --icon ${appIcon} --env ${buildEnv}"

fyne-cross linux ${buildFlags} --arch amd64 ./client-ui
fyne-cross windows ${buildFlags} --arch amd64 ./client-ui
fyne-cross darwin ${buildFlags} --arch amd64 ./client-ui
