* `go mod`的常用命令

| 命令 | 依赖 |
| --------------- | ---------------------------------------------- |
| go mod download | 下载依赖包到本地（默认为 GOPATH/pkg/mod 目录） |
| go mod edit     | 编辑 go.mod 文件                               |
| go mod graph    | 打印模块依赖图                                 |
| go mod init     | 初始化当前文件夹，创建 go.mod 文件             |
| go mod tidy     | 增加缺少的包，删除无用的包                     |
| go mod vendor   | 将依赖复制到 vendor 目录下                     |
| go mod verify   | 校验依赖                                       |
| go mod why      | 解释为什么需要依赖                             |

* 服务部署

`nohup` 

* 打包

1. 打开cmd
```shell
set GOARCH=amd64
set GOOS=linux
go build
```

2. 把打包后的文件上传服务器
3. 改变服务器文件权限

```shell
chmod 777 ./blog
```

4. 启动服务

```shell
nohup ./blog
```
