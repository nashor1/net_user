# net_user——bypass some hook
这是一个用goland写的netUser

参照微软官方文档

<https://learn.microsoft.com/zh-cn/windows/win32/api/lmaccess/nf-lmaccess-netuseradd>

调用netapi.dl下的win32函数

#使用方法

编译main.go文件后在windows终端下执行如下命令
此文件只能用于创建user用户不能创建管理员账户

```
main.exe name=xxx passwd=xxx
//name代表你的用户名
//passwd代表用户密码
```
