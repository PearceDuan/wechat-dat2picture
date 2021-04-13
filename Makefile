BIN_NAME:=wechat-dat2picture

.PHONY: build linux windows clean

prepare:
	go get -u github.com/go-bindata/go-bindata/...
	go-bindata -o=asset/asset.go -pkg=asset sample/...
build: linux
linux: prepare
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ${BIN_NAME} main.go
windows:prepare
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o ${BIN_NAME}.exe main.go
clean:
	rm -rf ${BIN_NAME}
	rm -rf ${BIN_NAME}.exe
	rm -rf asset/