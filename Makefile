build:
	dep ensure
	env GOOS=linux go build -ldflags="-s -w" -o bin/echo echo/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/ascii ascii/main.go