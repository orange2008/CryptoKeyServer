#!/usr/bin/bash
archs=(amd64 arm64)
for arch in ${archs[@]}
do
	env GOOS=linux GOARCH=${arch} go build -o CryptoKeyServer_${arch}_linux
    env GOOS=darwin GOARCH=${arch} go build -o CryptoKeyServer_${arch}_darwin
    env GOOS=windows GOARCH=${arch} go build -o CryptoKeyServer_${arch}_windows.exe
    env GOOD=openbsd GOARCH=${arch} go build -o CryptoKeyServer_${arch}_openbsd
done