# GolangDemo

借鉴https://github.com/EDDYCJY/go-gin-example的样例

使用Gin框架，规范了配置文件(conf文件夹)、日志处理(pkg/logging文件夹)、中间件(middleware文件夹)等，将路由与业务分离(url路由在routers中，业务逻辑在service中)。

主要有5点修改：
1. 加入了eureka注册；
2. 完善了redis更新删除逻辑。articles更新、删除时会先更新/删除数据库再删除redis，后续再考虑队列的方式；
3. log使用printf的形式，而不是println，Fatalf会退出；
4. 使用go module管理包；
5. 加入prometheus监控。


## 目录分工

pkg中定义公共内容，pkg/e/code.go中添加http状态码，在pkg/e/msg.go中添加对应的报错内容

middleware中定义中间件，比如jwt，metrics等

models目录中定义数据类型(与数据库对应)，同时定义响应表的操作方法(增删改查)

runtime中放应用运行过程中产生的内容

routers目录中定义url路由

service中定义业务逻辑，service/cache_service定义增删改查的redis操作，可以再新建目录编写业务逻辑

运行swagger init会产生docs文件夹，swagger的编写见routers/api/v1/article.go中

## 配置项
在conf/app.ini中进行配置：

app：
- PageSize表示多少个查询结果进行分页
- JwtSecret是jwt的参数，建议进行修改，建议进行修改
- RuntimeRootPath 运行产生文件的路径
- ImagePrefixUrl 页面上查看图片的前缀，url为服务器IP:PORT
- ImageSavePath 图片存储的路径
- ImageMaxSize上传图片的限制
- ImageAllowExts文件的格式
- LogSavePath 日志的存储路径
- LogSaveName日志名称
- LogFileExt 日志格式
- TimeFormat日志名称中的时间格式

server：
- RunMode 运行模式，有debug、release等选择，生产用release
- HttpPort 服务器的端口
- ReadTimeout 读取超时时间
- WriteTimeout 写超时时间

database：
- Type 数据库类型
- User 数据库用户名
- Password 数据库密码
- Host 数据库地址
- Name 数据库名，即database
- TablePrefix 数据表前缀，在本例中所有的数据表都会以blog_开头
- MaxIdleConn 数据库的最大空闲连接
- MaxOpenConn 数据库最大连接

redis:
- Host redis的地址
- Password redis密码
- MaxIdle 最大空闲连接数
- MaxActive 最大活动连接
- IdleTimeout 空闲连接超时时间

eureka:
- AppName 应用名称，即要注册到eureka的应用名
- EurekaServerUrl eureka地址
- StatusUrl 实例绝对状态页的URL路径，为其他服务提供信息时来找到这个实例的状态的路径
- HealthUrl 实例的绝对健康检查URL路径
- DataCenterInfo 实例被部署在的数据中心
- SecurePort 实例应该接收通信的安全端口，默认为443




## eureka注册
可以参考：https://github.com/Netflix/eureka/wiki/Eureka-REST-operations

即需要把如下格式的内容 POST到/eureka/apps/appID	
```
{
 "instance":{
 "hostName":"127.0.0.1",
 "app":"golangDemo",
 "ipAddr":"127.0.0.1",
 "vipAddress":"vendor",
 "status":"UP",
 "port":8000,
 "securePort":"443",
 "homePageUrl":"http://127.0.0.1:8000/",
 "statusPageUrl":"http://127.0.0.1:8000/info",
 "healthCheckUrl":"http://127.0.0.1:8000/metrics",
 "dataCenterInfo":{"name":"MyOwn"},
 "metadata":{"instanceId":"vendor:shalkdha"}
 }
}
```

## redis更新删除
articles更新、删除时会先更新/删除数据库再删除redis，对于redis删除失败，目前未处理，后续考虑将删除失败的任务放到队列中，起协程处理队列中的内容。

## go module管理包
go module使用是需要打开，windows中使用命令:`set GO111MODULE=on`，linux中使用`export GO111MODULE=on`

在运行go run/build/install main.go 会自动下载依赖包

## prometheus监控
prometheus监控在middleware/metrics/metrics.go中，目前监控了每支交易的响应时间和请求数





