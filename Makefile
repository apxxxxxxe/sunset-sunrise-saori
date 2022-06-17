build: sunset-saori.go go.mod go.sum
	GOOS=windows GOARCH=386 go build -o sunset.exe
