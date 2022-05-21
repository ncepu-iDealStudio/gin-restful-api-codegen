windows:
	SET CGO_ENABLED=0&SET GOOS=windows&SET GOARCH=amd64&go build -ldflags "-s -w" -o deploy/LRYCodeGen-win.exe LRYGoCodeGen
mac:
	CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o deploy/LRYCodeGen-mac LRYGoCodeGen
linux:
	SET CGO_ENABLED=0&SET GOOS=linux&SET GOARCH=amd64&go build -o ./deploy/LRYCodeGen-linux LRYGoCodeGen