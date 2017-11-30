GOARCH=amd64 GOOS=windows go build -o went_win64.exe
GOARCH=amd64 GOOS=darwin go build -o went_darwin
GOARCH=amd64 GOOS=linux go build -o went_linux64
tar -czf went.tar.gz went_linux64 went_darwin went_win64.exe 
