数据超时回调API
====
timeout_api 是一个超时回调机制，对发过来的一个数据指定超时时间，在这个时间过期后，触发一个访问API调用。

安装
-----
```bash
go get github.com/logmecn/timeout_api
```

使用方法
---
运行编译后的可执行文件，默认监听0.0.0.0:8001。这时如果发送一个数据如下：
```bash
curl -d '{"data":"aaaa=bbb","url":"http://127.0.0.1:8001","tout":5}' http://127.0.0.1
```
程序将在指定的 {tout} 这里是5秒之后，会去访问指定的链接: http://127.0.0.1:8001?aaaa=bbb

参数说明
-----------
| key  | value  | 说明                                         | 示例                        |
|------|--------|--------------------------------------------------|--------------------------------|
| data | string | url地址需要带的参数                                  | key1=value1&key2=value2        |
| url  | string | 当指定时间超时时，访问此链接+data拼接后的url  | http://pages\.com/?mkey=mvalue |
| tout | int    | 超时时间，单位为秒                           | 5                              |

拼接后的将访问的url地址是url+data，即： http://pages.com/?mkey=mvalue&key1=value1&key2=value2

