# gowiki
# Installation process

```go
mkdir gowiki && cd gowiki
go mod init gowiki
go mod tidy

go get github.com/labstack/echo/v4
vi server.go
go mod tidy
go run server.go

go install github.com/a-h/templ/cmd/templ@latest
go get github.com/a-h/templ
```

