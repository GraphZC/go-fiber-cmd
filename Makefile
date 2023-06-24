clean:
	rm -rf build

build: clean
	echo "Compiling for every OS and Platform"
	GOOS=freebsd GOARCH=386 go build -o bin/fiber-cmd-freebsd-386 main.go
	GOOS=linux GOARCH=386 go build -o bin/fiber-cmd-linux-386 main.go
	GOOS=windows GOARCH=386 go build -o bin/fiber-cmd-windows-386 main.go

run:
	go run .