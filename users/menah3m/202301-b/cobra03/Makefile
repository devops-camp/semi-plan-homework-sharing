# GOOS 操作系统类型
# GOARCH 操作系统架构
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)


build:
	go build -o out/aliyun-$(GOOS)-$(GOARCH) .

buildx:
	#CGO_ENABLED 是否使用 C语言 版本的 Go 编译器，默认为1，如果代码中使用了 C的库，就设置为1
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 make build
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 make build
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 make build
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 make build
	CGO_ENABLED=0 GOOS=windows GOARCH=arm64 make build
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 make build


clean:
	rm -rf out/