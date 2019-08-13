# ginode
一个和nodejs很像的框架

## 本地开发
beego框架自带了本地调试工具，在修改代码后可以自动重启，幸运的是，`bee`工具同样可以在`gin`项目中使用
> bee run

但是`bee`不能使用在非`GOPATH`目录下

另外一种使用了`dogo`，本项目包含了dogo.json配置，可根据自己项目路径修改
> dogo