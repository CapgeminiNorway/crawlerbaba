#!/usr/bin/env sh

echo "`go version`"

echo "making builds... linux"
GOOS=linux GOARCH=amd64 go build crawler.go
mv crawler crawler_linux

echo "making builds... mac"
GOOS=darwin GOARCH=amd64 go build crawler.go
mv crawler crawler_mac

echo "making builds... windows"
GOOS=windows GOARCH=amd64 go build crawler.go
mv crawler.exe crawler_win.exe

echo "making builds... -> default"
go build crawler.go

