# cannot use relative path in GOROOT, otherwise 6g not found. For example,
#   export GOROOT=../go  (=> 6g not found)
# it is also not allowed to use relative path in GOPATH
export GOROOT=$(realpath ../go)
export GOPATH=$(realpath .)
export PATH := $(GOROOT)/bin:$(GOPATH)/bin:$(PATH)

PKG="github.com/siongui/godom"
GO_VERSION=1.9.2
DEV_DIR=../


devserver: fmt local js
	@echo "\033[92mDevelopment Server Running ...\033[0m"
	@go run devserver/server.go -dir=example

js:
	@echo "\033[92mGenerating JavaScript ...\033[0m"
	@gopherjs build example/app.go -o example/app.js

fmt:
	@echo "\033[92mGo fmt source code...\033[0m"
	@go fmt *.go
	@go fmt devserver/*.go
	@go fmt example/*.go
	@go fmt wasm/*.go

godoc: local
	@echo "\033[92mgodoc server...\033[0m"
	#@godoc -http=:8000
	@godoc cmd/github.com/siongui/godom
	#@godoc -url "pkg/github.com/siongui/godom"

local:
	@[ -d src/${PKG}/ ] || mkdir -p src/${PKG}/
	@cp *.go src/${PKG}/

update_ubuntu:
	@echo "\033[92mUpdating Ubuntu ...\033[0m"
	@sudo apt-get update && sudo apt-get upgrade && sudo apt-get dist-upgrade

download_go:
	@echo "\033[92mDownloading and Installing Go ...\033[0m"
	@cd $(DEV_DIR) ; wget https://storage.googleapis.com/golang/go$(GO_VERSION).linux-amd64.tar.gz
	@cd $(DEV_DIR) ; tar xvzf go$(GO_VERSION).linux-amd64.tar.gz

install:
	@echo "\033[92mInstalling GopherJS ...\033[0m"
	go get -u github.com/gopherjs/gopherjs

clean:
	rm -rf bin/ pkg/ src/
