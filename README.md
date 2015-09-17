# MySQL Monitor Script

## 需要注意的点
1. 在编译时将base/cfg.go文件里面硬编码值修改成合适的值.
1. mysql socket文件命名必须符合规范,类似这样,mysql的前缀+底划线+使用此数据库的应用名称`mysql_appname.sock`
1. 默认部署这个脚本的机器上需要同时部署falcon-agent,且开的1988的http端口
