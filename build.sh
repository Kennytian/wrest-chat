#!/bin/sh
#

set -e
set -o noglob

###########################################

export CGO_ENABLED=0
export GO111MODULE=on

export GOOS=windows
export GOARCH=386

echo building for $GOOS/$GOARCH

target=build/wrest.exe
go build -ldflags="-s -w" -o $target main.go

####################################################################

cp README.md build/
cp wcf-sdk/libs/sdk.dll build/
cp wcf-sdk/libs/spy.dll build/

mv build wechat-rest
zip -r wechat-rest.zip wechat-rest/