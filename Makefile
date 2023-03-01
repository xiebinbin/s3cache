all: deps clear_bin build_mac_arm build_linux_amd64 cp_config test

deps: FORCE
	CGO_CFLAGS_ALLOW=-Xpreprocessor go get ./...
clear_bin: FORCE
	rm -rf ./bin && mkdir -p bin
build_mac_arm: FORCE
	CGO_ENABLED=1 GOOS=darwin GOARCH=arm64 CGO_CFLAGS_ALLOW=-Xpreprocessor go build -o ./bin/mac_arm64 ./server.go
build_linux_amd64: FORCE
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 CC=x86_64-linux-musl-gcc  CXX=x86_64-linux-musl-g++ CGO_CFLAGS_ALLOW=-Xpreprocessor go build -o ./bin/linux_amd64 ./server.go
cp_config: FORCE
	cp config.example.yml ./bin/config.yml && cp readme.md ./bin/readme.md && cp image-pro.conf ./bin/image-pro.conf && cp nginx.conf ./bin/nginx.conf
test: FORCE
	CGO_CFLAGS_ALLOW=-Xpreprocessor go test -v ./...
FORCE: