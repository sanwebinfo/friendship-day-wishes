BUILD_DIR=./build

clean:
	rm -rf ${BUILD_DIR}

build:
	CGO_ENABLED=0 GOOS=linux   GOARCH=amd64       go build -o build/wish-linux-amd64       wish.go
	CGO_ENABLED=0 GOOS=linux   GOARCH=arm64       go build -o build/wish-linux-arm64       wish.go