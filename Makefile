BIN_NAME:=wechat-dat2picture

.PHONY: build linux windows clean

build: linux
linux:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ${BIN_NAME} main.go
windows:
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o ${BIN_NAME}.exe main.go
clean:
	rm -rf ${BIN_NAME}
	rm -rf ${BIN_NAME}.exe