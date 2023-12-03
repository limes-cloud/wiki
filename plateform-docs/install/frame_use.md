
##### 框架简介
前端使用字节开源框架 [Arco Design](https://arco.design/)，arco-design主要服务于字节跳动旗下中后台产品的体验设计和技术实现，主要由UED设计和开发同学共同构建及维护。

后端使用B站开源框架 [Go Kratos](https://go-kratos.dev/docs/)，go-kratos是一款卓越的后端框架，提供了丰富的可插拔接口，使得开发变得高效灵活。在项目开发中，我们通常会封装一些常用的中间件，如MySQL、Redis、Config等。然而，由于基于kratos进行开发时，一旦中间件代码发生变更，就需要在所有相关文件中进行相应修改，这可能会增加开发的复杂性。

因此，我在go-kratos的基础上进行了一些创新，开发了一些常见的插件，以加速整个系统的开发进度。这些插件不仅使得中间件的使用更为便捷，还减轻了对代码的频繁修改，提高了开发效率。这样，我们可以更专注于业务逻辑的实现，而不必过多关注底层框架的变更。

##### 插件安装

###### 安装go-kratos
``` 
go install github.com/limes-cloud/kratos/cmd/kratos@latest 
```


###### 安装protoc
```
进入一下连接，安装自己合适的版本即可
https://github.com/protocolbuffers/protobuf/releases/
```


###### 安装扩展包
```
go install github.com/limes-cloud/kratos/cmd/protoc-gen-go-http@latest &&\
go install github.com/limes-cloud/kratos/cmd/protoc-gen-go-errors@latest &&\
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest &&\
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest &&\
go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest &&\
go install github.com/envoyproxy/protoc-gen-validate@latest
```


###### 进行自动生成grpc代码配置
- 1）打开Goland设置
- 2）找到Tool 》 File Watcher
- 3）添加custom配置,文件类型选择proto buffer
- 4）配置一下命令,program 选择你下载的protoc,args配置为以下代码

```
--proto_path=$FileDir$
--proto_path=$ProjectFileDir$/third_party
--go_out=:.
--go-http_out=.
--go-grpc_out=.
--go-errors_out=.
--validate_out=lang=go:.
$FilePath$
```
这样你只要保存proto文件就会自动生成对应的grpc go的代码了。

##### 新建服务
``` 
kratos new [服务名]
```

##### 目录说明

```
internal 我们的业务实现主要在这里
    |- service 主要实现grpc的接口,新加一个proto方法，就需要在这里新增一个对应的方法调用
    |- logic 主要是具体函数的业务逻辑实现
    |- model 主要是具体函数的数据层的操作
```