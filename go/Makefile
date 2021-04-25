all: macos ios

macos:
	CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -buildmode=c-shared -o ../macos/server-macos.dylib main.go
ios-simulator:
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=1 SDK=iphonesimulator CC=`go env GOROOT`/misc/ios/clangwrap.sh go build -tags ios -buildmode=c-shared -o ../ios/server-ios.dylib main.go
ios:
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=1 SDK=iphoneos CC=`go env GOROOT`/misc/ios/clangwrap.sh CGO_CFLAGS="-fembed-bitcode" go build -tags ios -buildmode=c-shared -o ../ios/server-ios.dylib main.go

