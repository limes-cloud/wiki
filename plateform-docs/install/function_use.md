##### 使用说明 
limes-cloud/kratos 为了简化框架的使用，基于go-kratos进行了二开，封装了一些常见的中间件，使用context上下文进行传递可以进行直接使用，可以通过配置就可以快速开始。

##### 配置接入
limes-cloud/kratos 配置接入和kratos一直，在kratos的基础上新增了配置中心，可以使用以下方式快速接入
``` 
// 参数传入初始化
configure.New(host,name,token)

// 从环境变量中初始化
kratos.Config(configure.NewFromEnv())

// 环境配置变量
APP_NAME=Manager
CONF_HOST=localhost:6082
CONF_TOKEN=8A62C35740D5817A3F8C6958C4BE6B2C
```
从配置中心接入需要先运行配置中心，如果你嫌比较麻烦，你可以直接连接我线上的配置中心
``` 
HOST=101.34.229.39:6082
具体的app_name 和 token可以在配置中心中查看
```

##### 服务配置
``` 
# 服务配置
server:
  http: # http服务配置
    addr: 0.0.0.0:7003  # 端口地址
    timeout: 1s # 超时时间
    formatResponse: true  # 是否返回格式化
    marshal: # 序列化配置
      emitUnpopulated: true # 默认值不忽略
      useProtoNames: true # 使用proto name返回http字段
    cors: # 跨域配置
      allowCredentials: true
      allowOrigins: [ "*" ]
      allowMethods: [ "GET","POST","PUT","DELETE","OPTIONS" ]
      AllowHeaders: ["Content-Type", "Content-Length", "Authorization"]
      ExposeHeaders: ["Content-Length", "Access-Control-Allow-Headers"]
  grpc: # grpc服务配置
    addr: 0.0.0.0:8003 # 端口地址
    timeout: 1s # 超时时间
```

##### 日志配置
``` 
log:
  level: 0 #日志输出等级
  output:
    - stdout # stdout:控制台输出，k8s日志收集
    - file # file:输出到文件
  file: #output存在file时此配置才可生效
    name: ./tmp/runtime/output.log #日志存放地址
    maxSize: 1 #日志文件最大容量,单位m
    maxBackup: 5 #日志文件最多保存个数
    maxAge: 1 #保留就文件的最大天数,单位天
    compress: false #是否进行压缩归档
```
日志使用
``` 
ctx.Logger().Info("error","服务出现问题")
```

##### 数据库配置
```
database:
  # 数据库实例名称,如有多个数据库可新增
  system:
    enable: false #是否启用数据库
    drive: mysql #数据库类型
    dsn: root:root@tcp(127.0.0.1:3306)/operation?charset=utf8mb4&parseTime=True&loc=Local #连接dsn
    maxLifetime: 2h #最大生存时间
    maxOpenConn: 10 #最大连接数量
    maxIdleConn: 10 #最大空闲数量
    logLevel: 1 #日志等级 1:error 2:warn 3:info 4:debug
    slowThreshold: 2s #慢sql阈值
```
数据库使用
``` 
// 获取第一个数据库实例
ctx.DB()
// 获取指定数据库实例
ctx.DB('实例名称')
```

##### redis配置
```
redis:
  catch: #redis实例名称,如有多个redis可新增
    enable: false #是否启用redis
    host: 127.0.0.1:6379 #redis地址
    username:  #连接账号
    password:  #连接密码
```
redis 使用
``` 
// 获取第一个redis实例
ctx.Redis()
// 获取指定redis实例
ctx.Redis('实例名称')
```

##### 文件加载器
在我们开发过程中，有很多需要去读取各种文件的操作，而loader组件把这些操作进行简单话，可以通过以下快速获取，返回文件[]byte。

```
loader:
  login: static/cert/login.key
```
loader使用
``` 
ctx.Loader('实例名称')
ctx.Loader("login")
```

##### 并发池配置
``` 
pool: #并非池配置
  size: 100000 #最大协程数量
  expiryDuration: 30s #过期时间
  preAlloc: true #是否预分配
  maxBlockingTasks: 1000 #最大的并发任务
  nonblocking: true #设置为true时maxBlockingTasks讲失效，不限制并发任务
```
并发池使用
``` 
type Task struct {
	Number int
}

func (t *Task) Run() {
	fmt.Println(t.Number)
}


func main(){
    ctx := kratos.MustContext(context.Background())
    task := &Task{}
	_ = ctx.Go(task)
}

```

##### 邮箱配置
我们的业务中常常需要发送邮件需要进行通知，使用邮箱插件可以快速使用
``` 
email: # 邮件发送相关配置
  template: #邮件模板
    captcha:  #邮件模板名称
      subject: 验证码发送通知 #邮件模板主题
      path: static/template/email/default.html #邮件模板路径
      type: "text/html" #文本内容格式
  user: xxx@qq.com #发送者
  host: smtp.qq.com #发送host
  port: 25  #发送端口
  name: 青柠校园 #发送者名称
  password: xxx # 发送者授权码
```
邮箱使用
``` 
template := ctx.Email().Template('邮件模板名称')
template.Send('发送者邮箱','发送者名称','模板变量')
```

##### 验证码配置
验证码配置支持邮箱验证码和图片验证码，可以快速支持在业务场景中的验证场景
``` 
captcha:
  login:  #验证码名称
    type: image #验证码类型 目前支持image/email
    length: 6 #验证码长度
    expire: 180s #过期时间
    redis: cache #redis名称
    height: 801 #图片高度,仅image有效
    width: 240 #图片宽度,仅image有效
    skew: 0.7 #验证数字倾斜程度,仅image有效
    dotCount: 80 #干扰像素点数量,仅image有效
  changePassword:  #验证码名称
    type: email #验证码类型
    length: 6 #验证码长度
    expire: 180s #过期时间
    redis: cache #redis名称
    template: captcha #邮箱模板名称,仅email有效
```
captcha 使用
``` 
ctx.Captcha()

// captcha接口顶底
type Captcha interface {
    // 发送邮箱验证码
	Email(tp string, ip string, to string) (Response, error)
	// 生成图片验证码
	Image(tp string, ip string) (Response, error)
	// 验证邮箱
	VerifyEmail(tp, ip, id, answer string) error
	// 验证图片验证码
	VerifyImage(tp, ip, id, answer string) error
}
```

##### JWT 访问鉴权配置
主要是用来判断用户是否有访问权限，只要配置了就会自动开启，可以使用ctx.JWT()进行生成token
``` 
jwt:
  redis: cache #redis名称
  secret: "limes-cloud" #密钥
  expire: 2h #过期时间
  renewal: 1s #续期时间
  whitelist: #忽略token校验以及鉴权的白名单
    POST:/v1/login/captcha: true #http 方法名:path  grpc GRPC:operation
    POST:/v1/login: true
    POST:/v1/logout: true
    POST:/v1/token/refresh: true
```
jwt插件使用
``` 
ctx.JWT()
// jwt接口定义
type Jwt interface {
    // 生成Token
	NewToken(m map[string]any) (string, error)
	// 解析token
	Parse(ctx context.Context, dst any) error
	// 解析token成mao
	ParseMapClaims(ctx context.Context) (map[string]any, error)
	// 判断访问路径是否为白名单
	IsWhitelist(path string) bool
	// 判断token是否为黑名单
	IsBlacklist(token string) bool
	// 添加token为黑名单
	AddBlacklist(token string)
	// 从ctx中获取token
	GetToken(ctx context.Context) string
	// 进行token续期，生成新token
	Renewal(ctx context.Context) (string, error)
}
```

##### authentication 资源鉴权配置
主要是用来判断用户是否对接口是否具有权限，主要是基于casbin进行验证。
``` 
authentication:
  db: system #db名称
  redis: cache #redis名称
  roleKey: role_keyword #jwt中存储角色的关键字的下标
  skipRole: ["super_admin"] #跳过检查的角色
  whitelist:
    method:path: true
```
authentication 使用
``` 
ctx.Authentication()

// 鉴权接口
type Authentication interface {
    // 添加到白名单
	AddWhitelist(path string, method string)
	// 移除拍名单
	RemoveWhitelist(path, method string)
	// 判断是否为白名单
	IsWhitelist(path string, method string) bool
	// 是否具有权限
	Auth(role, path, method string) bool
	// 获取角色名称
	GetRole(ctx context.Context) (string, error)
	// 获取casbin实例
	Enforce() *casbin.Enforcer
	// 是否跳过角色
	IsSkipRole(role string) bool
}
```

##### 限流器配置
```
ratelimit: true #是否开自适应限流
```

##### IP获取
快速获取ip地址
```
ctx.IP()
```

##### logging 请求日志
``` 
logging:    
    whitelist:
        method:path: true
```