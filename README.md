# HelloWorld

>一个用Go语言开发的，基于兴趣爱好的IM通讯服务框架。

## 写在前面

很高兴你关注到了这个项目。这个是一个用 Go 语言开发的基于 Socket/WebSocket 的通讯服务。可能略有繁琐，但相信其并不复杂。

目前支持基于 Socket 的通讯和基于 WebSocket ，并且基于 Socket 的通讯是包含客户端的，基于 WebSocket 的没有包含客户端，后期将逐步完善。

如果你有好的意见或建议请通过 `issues` 提出，如果你想加入这个项目可以直接 `Fork` 并将改动 `new pull request`

   
### 项目开发资料
    https://github.com/Yanlong-LI/HelloWorld-GO

#### 客户端源码地址：
    https://github.com/Yanlong-LI/HelloWorldClient

#### 服务端源码地址：
    https://github.com/Yanlong-LI/HelloWorldServer
    
    
### 开发设想
中心化管理：实现有点问题

    一个中心服务器
        负责用户的管理和服务器管理
        
    多个第三方服务器
        为用户提供多个频道
        
    客户端
        用户端
        
    问题在于如果中心服务器崩掉，所有服务器都将无法使用
探讨去中心化服务的可行性：
    
    1、服务发现
    2、服务同步、互联
    3、隐私保护、安全保护
    4、