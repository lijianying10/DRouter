# DRouter 

## 静态文件服务与反向代理路由工具

## 解决问题
1. 网站跨域
2. 在自己网站上请求其他网站的数据（网站整合）
3. 前端开发脚手架，无需更新js地址即可请求到测试服务器。

## 文件说明
config：配置文件样例
main.go ： 程序主要入口 记录服务开始时间，绑定动态路由，开启http服务
logD： 简单的log模块，记录所有接口走向调用次数以及报错。
Router： 读取conf模块中的配置数据，绑定路由。
conf： 读取配置文件并保存配置到内存中

## 配置文件请注意：
1. 文档使用的是json
2. 路由中代理的distnation最后不要有斜杠
3. 支持其他端口
4. 如果接口重复了。则后出现的为准

## 使用说明：
1. 配置文件位置首先检查 /etc/drouter.conf
2. 然后检查 ./drouter.conf
3. 最后如果命令中出现了文件参数则直接读取文件

## 下载地址：
[DRouter_darwin_386](http://picstorage.qiniudn.com/DRouter_darwin_386)


[DRouter_darwin_amd64](http://picstorage.qiniudn.com/DRouter_darwin_amd64)


[DRouter_freebsd_386](http://picstorage.qiniudn.com/DRouter_freebsd_386)


[DRouter_freebsd_amd64](http://picstorage.qiniudn.com/DRouter_freebsd_amd64)


[DRouter_freebsd_arm](http://picstorage.qiniudn.com/DRouter_freebsd_arm)


[DRouter_linux_386](http://picstorage.qiniudn.com/DRouter_linux_386)


[DRouter_linux_amd64](http://picstorage.qiniudn.com/DRouter_linux_amd64)


[DRouter_linux_arm](http://picstorage.qiniudn.com/DRouter_linux_arm)


[DRouter_netbsd_386](http://picstorage.qiniudn.com/DRouter_netbsd_386)


[DRouter_netbsd_amd64](http://picstorage.qiniudn.com/DRouter_netbsd_amd64)


[DRouter_netbsd_arm](http://picstorage.qiniudn.com/DRouter_netbsd_arm)


[DRouter_openbsd_386](http://picstorage.qiniudn.com/DRouter_openbsd_386)


[DRouter_openbsd_amd64](http://picstorage.qiniudn.com/DRouter_openbsd_amd64)


[DRouter_plan9_386](http://picstorage.qiniudn.com/DRouter_plan9_386)


[DRouter_windows_386.exe](http://picstorage.qiniudn.com/DRouter_windows_386.exe)


[DRouter_windows_amd64.exe](http://picstorage.qiniudn.com/DRouter_windows_amd64.exe)



## TODOList：
1. https 支持
2. 换个更好的配置文件