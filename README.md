# gin-example
my go gin-example

### mac 编译linux 64位
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a ./

### Swagger文档
下载：
go get -u github.com/swaggo/swag/cmd/swag

生成：
swag init

访问链接：
http://地址/swagger/index.html
