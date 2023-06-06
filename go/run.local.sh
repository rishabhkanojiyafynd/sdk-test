rm -rf go.mod
go mod init go_lang_test
go get -u "github.com/pixelbin-dev/pixelbin-go/v2@<branch-name>"
go run index.go
