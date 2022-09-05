# learn-go-swagger
learn-go-swagger

## Dependencies

生成服务端和客户端代码
```shell
go get github.com/go-openapi/errors
go get github.com/go-openapi/loads
go get github.com/go-openapi/runtime
go get github.com/go-openapi/spec
go get github.com/go-openapi/strfmt
go get github.com/go-openapi/swag
go get github.com/go-openapi/validate
go get github.com/jessevdk/go-flags
go get golang.org/x/net/context
```

生成规范
```shell
go get github.com/go-openapi/errors
go get github.com/go-openapi/strfmt
go get github.com/go-openapi/swag
go get github.com/go-openapi/validate
```

## 构建CLI
生成CLI
```shell
# swagger generate cli -f [http-url|filepath] --cli-app-name [app-name]
swagger generate cli -f examples/cli/swagger.yml --cli-app-name demo
```