# log
该日志在zap日志的基础上扩展了链路追踪的能力。
在当前项目下的etc目录新建一个log.yaml，配置信息如下
```yaml
log:
  type: console
  fileName: app.log
  filePath: logs
  level: debug
```
关于参数含义见config.go中注释
如需开启trace功能需要在main.go加上
```go
log.SetEnableTrace(true)
```

