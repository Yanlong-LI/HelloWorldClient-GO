操作码
===
> 服务操作码

|CODE|NAME|ACTION|DESCRIPTION|
|---|---|----|---
|0|Config|Receive|接收配置参数|
|1|Disconnected|Receive Send|断开连接|
|6001|Login|Receive|用户登陆|
|6002|LoginSuccess|Send|登陆成功，发送令牌|
|6003|Ticket|Send|发送登陆验证|
|6004|AuthTicket|Receive|收到验证单|
|6005|LoginFail|Send|登陆失败|