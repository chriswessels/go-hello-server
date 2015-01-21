echo "Building Linux binary..."
GOOS=linux GOARCH=amd64 go build -o main_linux main.go
echo "Building OSX binary..."
go build -o main_osx main.go
