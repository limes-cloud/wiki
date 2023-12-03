##### 系统预览
- 系统地址：[点击预览](http://admin.qlime.cn/)
- 登录账号：128029101@qq.com
- 登录密码：12345678

##### 环境要求
- nodejs
- golang(1.19+)
- mysql
- redis

##### 安装前须知


##### 管理平台前端安装

```
# 拉取代码
git clone https://github.com/limes-cloud/admin-platform-web

# 执行安装
pnpm install

# 执行运行
pnpm run dev
```

##### configure 配置中心安装
```
# 拉取代码
git clone https://github.com/limes-cloud/configure

# 运行服务
go run cmd/configure/main.go

# 进入配置中心独立管理面板
http://localhost:6080
```

##### gateway 统一网关安装
```
# 拉取代码
git clone https://github.com/limes-cloud/gateway

# 运行服务
go run cmd/gateway/main.go
```

##### manager 管理中心安装
```
# 拉取代码
git clone https://github.com/limes-cloud/manager

# 运行服务
go run cmd/manager/main.go
```

##### resource 资源中心安装
```
# 拉取代码
git clone https://github.com/limes-cloud/resource

# 运行服务
go run cmd/resource/main.go
```


