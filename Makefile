all: deps clear_bin build_server_darwin_arm build_server_linux_amd64 build_m3u8_server_darwin_arm build_m3u8_server_linux_amd64 cp_config test

deps: FORCE
	CGO_CFLAGS_ALLOW=-Xpreprocessor go get ./...
clear_bin: FORCE
	rm -rf ./bin && mkdir -p bin
build_server_darwin_arm: FORCE
	CGO_ENABLED=1 GOOS=darwin GOARCH=arm64 CGO_CFLAGS_ALLOW=-Xpreprocessor go build -o ./bin/server_darwin_arm64 ./cmd/server/main.go
build_server_linux_amd64: FORCE
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 CC=x86_64-linux-musl-gcc  CXX=x86_64-linux-musl-g++ CGO_CFLAGS_ALLOW=-Xpreprocessor go build -o ./bin/server_linux_amd64 ./cmd/server/main.go
build_m3u8_server_darwin_arm: FORCE
	CGO_ENABLED=1 GOOS=darwin GOARCH=arm64 CGO_CFLAGS_ALLOW=-Xpreprocessor go build -o ./bin/m3u8_server_darwin_arm64 ./cmd/m3u8-server/main.go
build_m3u8_server_linux_amd64: FORCE
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 CC=x86_64-linux-musl-gcc  CXX=x86_64-linux-musl-g++ CGO_CFLAGS_ALLOW=-Xpreprocessor go build -o ./bin/m3u8_server_linux_amd64 ./cmd/m3u8-server/main.go
cp_config: FORCE
	cp config.example.yml ./bin/config.yml && cp readme.md ./bin/readme.md && cp image-pro.conf ./bin/image-pro.conf && cp nginx.conf ./bin/nginx.conf
test: FORCE
	CGO_CFLAGS_ALLOW=-Xpreprocessor go test -v ./...
FORCE: