windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o deploy/LRYCodeGen.exe LRYGoCodeGen
mac:
	CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o deploy/LRYCodeGen LRYGoCodeGen
linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o deploy/LRYCodeGen LRYGoCodeGen