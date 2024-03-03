#!/usr/bin/bash
archs=(amd64 arm64)
mkdir -p ./bin
echo "Started to compile."
for arch in ${archs[@]}
do
	env CGO_ENABLED=0 GOOS=linux GOARCH=${arch} go build -o ./bin_cgoless/CryptoKeyServer_${arch}_linux
    echo "Finished compilation for ${arch}_linux"
    env CGO_ENABLED=0 GOOS=darwin GOARCH=${arch} go build -o ./bin_cgoless/CryptoKeyServer_${arch}_darwin
    echo "Finished compilation for ${arch}_darwin"
    env CGO_ENABLED=0 GOOS=windows GOARCH=${arch} go build -o ./bin_cgoless/CryptoKeyServer_${arch}_windows.exe
    echo "Finished compilation for ${arch}_windows"
done