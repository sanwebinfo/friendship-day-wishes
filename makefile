BUILD_DIR=./build

clean:
	rm -rf ${BUILD_DIR}

build:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64       go build -o build/wish-windows-amd64.exe wish.go
	CGO_ENABLED=0 GOOS=windows GOARCH=arm64       go build -o build/wish-windows-arm64.exe wish.go
	CGO_ENABLED=0 GOOS=windows GOARCH=386         go build -o build/wish-windows-386.exe   wish.go
	CGO_ENABLED=0 GOOS=linux   GOARCH=amd64       go build -o build/wish-linux-amd64       wish.go
	CGO_ENABLED=0 GOOS=linux   GOARCH=386         go build -o build/wish-linux-386         wish.go
	CGO_ENABLED=0 GOOS=linux   GOARCH=arm64       go build -o build/wish-linux-arm64       wish.go
	CGO_ENABLED=0 GOOS=linux   GOARCH=arm GOARM=7 go build -o build/wish-linux-arm-7       wish.go
	CGO_ENABLED=0 GOOS=darwin  GOARCH=amd64       go build -o build/wish-darwin-amd64      wish.go
	CGO_ENABLED=0 GOOS=darwin  GOARCH=arm64       go build -o build/wish-darwin-arm64      wish.go
	CGO_ENABLED=0 GOOS=android GOARCH=arm64       go build -o build/wish-android-arm64      wish.go