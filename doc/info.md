# 端启动方式

端的启动方式有两种：

- 创建房间
  - 创建房间成功后会生成一串字符串，房主可分享字符串给好友
- 加入房间
  - 通过好友分享的字符串可申请加入对应房间

端启动后，会使用tcp协议监听端口，接收其他端的消息



# 端对端消息传递

消息体

```go
type Message struct {
	// 消息类型
	Typ int `json:"typ"`
	// 消息内容
	Content interface{} `json:"content"`
	// 发送者个人信息
	SendUser User `json:"send_user"`
}
```

消息类型

```go
const (
	// TypChatMsg 聊天消息
	TypChatMsg = 1
	// TypApplicationMsg 申请加入room消息
	TypApplicationMsg = 2
	// TypApplicationResultMsg 处理申请结果消息
	TypApplicationResultMsg = 3
	// TypNewPeerAddedMsg 新人加入消息
	TypNewPeerAddedMsg = 4
)
```



# 消息处理方式

- 为方便描述，将所有端分成3类：
  - 房主端：成功创建房间的端
  - 申请端：尝试通过分享字符串加入已有房间的端
  - 现有成员端：已经通过加入房间的端

## 申请加入room消息

申请端根据房主分享的字符串可做加入room申请，房主端则会收到申请端的申请消息，房主端的处理逻辑如下：

- 决定是否批准加入：
  - 不同意：
    - 给申请端发送不同意的`处理申请结果消息`
  - 同意
    - 给申请端发送同意的`处理申请结果消息`
    - 给房间现有成员端发送`新人加入消息`，更新内存内room信息，主要是更新room成员消息



## 处理申请结果消息

申请端发起申请后会等待房主的审批结果，处理逻辑如下：

- 不同意：
  - 结束进程
- 不同意：
  - 更新内存内的room消息



## 新人加入消息

房主端同意申请端的加入房间申请后，应该通知所有的现有成员端，现有成员端收到消息后，处理逻辑如下：

- 更新内存内的room消息