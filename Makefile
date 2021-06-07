hello:
	echo "Hello"


build:
	go build -o bin/server server.go


run:
	go run server.go


compile:
	echo "Compiling for every OS and Platform"
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 server.go
	GOOS=linux GOARCH=386 go build -o bin/main-linux-386 server.go
	GOOS=windows GOARCH=386 go build -o bin/main-windows-386 server.go