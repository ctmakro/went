set GOARCH=amd64
set GOOS=windows
go build -o went_win64.exe

set GOOS=darwin
go build -o went_darwin

set GOOS=linux
go build -o went_linux64
