# CSA考核 IM即时通讯



### 代码总体架构

```
├── README.md           					// 说明文档
	├── config
		└── config.ymal						// 配置文件
    ├── api									// api接口
    ├── service								// 服务逻辑
    ├── router								// 路由
    ├── chat								// 聊天服务
    ├── dao									// 数据库层
    ├── models								// 模型层
    ├── utils								// 工具
    ├── settings							// 配置设置
    ├── logger								// 日志初始化
    ├── log									// 日志
    ├── go.mod
    └── main.go
```



### 功能列表：

- **登录注册**
- **添加删除好友**
- **单人聊天**
- 修改头像
- 群聊天
- 群好友列表
- 添加群组
- 文本消息
- 图片消息
- 文件发送



### 技术和框架

- web框架Gin
- 长连接WebSocket
- 日志框架zap
- 配置管理viper
- ORM框架gorm
- 数据库MySQL
- 消息队列NSQ



#### 用户模块

- 用户的注册与登录
- 获取用户信息
- 更新头像

#### 好友模块

- 好友添加删除

#### 聊天功能

- Websocet 保持用户长连接
- 群组群聊功能
