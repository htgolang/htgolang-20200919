cookie 存储位置: 浏览器上
服务器如何告诉浏览器进行存储: Response Set-Cookie
浏览器如何携带信息： Request Cookie

http 无状态

1. 我的浏览器 第一次请求 删除用户操作
2. 我的浏览器 第二次请求 添加用户操作

客户端浏览器请求时检查Cookie头 sid
若无sid 服务器端生成一个唯一表示并且告诉浏览器你的ID是sid
浏览器以后在请求时都携带sid信息


1. 认证
给客户端分配唯一标识SID

login user/password 认证成功 服务器端 SID  状态=1

每个请求 需要登陆才能访问 客户端COOKIE SID -> 服务器查询 状态 1 => 允许操作
                                                      非1 => 跳转到登陆页面

session 服务器
session和客户端的关联: sid
浏览器sid如何告诉给服务器端: Cookie
浏览器sid是如何分配的: 服务器端检查客户端无SID生成的一个唯一标识, Set-Cookie

Session管理 => http 未提供基本功能

jwt
token