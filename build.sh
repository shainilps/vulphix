#! /bin/bash

echo "Creating bin"
mkdir -p bin 

echo "Building for linux"
CGO=0  GOOS=linux GOARCH=amd64 go build -o ./bin/vulphix-linux-amd64

echo "Building for Windows"
CGO=0  GOOS=windows GOARCH=amd64 go build -o ./bin/vulphix-windows-amd64.exe

echo "Building for Mac"
CGO=0  GOOS=darwin GOARCH=arm64 go build -o ./bin/vulphix-darwin-arm64
