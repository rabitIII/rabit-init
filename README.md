# React + Golang

> 技术栈： gorm + gin



## 目录配置

```powershell
|- grs-server
    |- api          # 请求
    |- config       # 初始化的样式表
    |- core         # 初始化
    |- flag
    |- global       # 全局配置
    |- middleware
    |- models
    |- routers      # api路由
    |- service
    |- utils
    go.mod          # go配置
    main.go         # 程序入口

```

### 配置 gin web框架
```powershell
go get 
```

### 配置 yaml 文件
首先将一些端口（后端api的端口、Mysql数据的）写进`settings.yaml`样式文件中，格式如下
```yaml
system:         # gin框架的别名
  ip: ""        # 运行的ip地址
  port: 8080    # 运行时的端口号
  env: "dev"    # 框架运行的模式
```

### 配置数据库

第三方库下载
```powershell
go get gorm.io/driver/mysql
go get gorm.io/gorm
```


### 配置 日志
```powershell
go get github.com/sirupsen/logrus 
```