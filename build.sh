#! /bin/bash

echo "Creating bin"
mkdir -p bin

echo "Building for linux"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/vulphix-linux-amd64

echo "Building for Windows"
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/vulphix-windows-amd64.exe

echo "Building for Mac"
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o ./bin/vulphix-darwin-arm64
