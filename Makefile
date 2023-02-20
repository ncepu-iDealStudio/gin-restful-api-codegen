windows:
	SET CGO_ENABLED=0&SET GOOS=windows&SET GOARCH=amd64&go build -ldflags "-s -w" -o ./public/LRYCodeGen-win64.exe LRYGoCodeGen
mac:
	SET CGO_ENABLED=0&SET GOOS=darwin&SET GOARCH=amd64&go build -ldflags "-s -w" -o ./public/LRYCodeGen-mac LRYGoCodeGen
linux:
	SET CGO_ENABLED=0&SET GOOS=linux&SET GOARCH=amd64&go build -ldflags "-s -w" -o ./public/LRYCodeGen-linux LRYGoCodeGen