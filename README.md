# 项目结构

- CMD 通常用于存储项目的入口
  - api 存放服务的入口文件
  - migrate 存放数据库迁移文件
- db 存放数据库文件
- service 存放业务逻辑
- config 存放配置文件

godotenv.Load()

# 数据存储

- config.go 用于存储配置结构体

  - var Env = initDBConfig()
  - dbConfig struct: 数据库连接所需的参数
  - initDBConfgi: 初始化配置
  - getEnv: 从环境中获取变量

- db.go 用于定义数据库

  - NewMySQLStoratge: 通过外部参数初始化数据库

- model.go 用于定义实体

  - xxxStore interface: 实体需要实现的操作方法
  - xxx struct: 对应数据库的实体
  - xxxPayload struct: 对应请求或响应的实体

- service/xxx: 实现实体的业务逻辑
  - router.go: 实现处理逻辑
  - store.go: 实现实体的存储逻辑
