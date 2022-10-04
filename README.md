# swag-crash-demo-001

## How to recur swag crash problem?

```bash
$ go install github.com/swaggo/swag/cmd/swag@latest
$ swag --version                                        
swag version v1.8.6
$ go version 
go version go1.18.5 linux/amd64
$ swag init -g main.go --parseInternal --parseDependency
2022/10/04 18:56:10 Generate swagger docs....
2022/10/04 18:56:10 Generate general API Info, search dir:./
2022/10/04 18:56:12 ParseComment error in file /data/go/pkg/mod/github.com/astaxie/beego@v1.12.3/parser.go :form is not supported paramType
```
