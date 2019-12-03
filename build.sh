rm -rf bin
mkdir bin
GOOS=windows GOARCH=amd64 go build -o bin/HttpFileServer-amd64.exe
GOOS=linux GOARCH=amd64 go build -o bin/HttpFileServer-linux-arm
GOOS=linux GOARCH=arm go build -o bin/HttpFileServer-linux-amd64
GOOS=darwin GOARCH=amd64 go build -o bin/HttpFileServer-osx-amd64
