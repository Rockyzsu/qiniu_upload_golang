# 简约的七牛文件上传页面

简约且简单，用于上传图片与文本，管理历史图片。

你是否使用七牛上传了不少图片，却发现找回以前上传的历史图片非常困难？ 所以写了个web用来管理七牛的图片系统


----

* 主要因为手机app端七牛没有上传的客户端，并且自己的电脑多操作系统，最方便的办法是通过浏览器上传文件，随时随地只要能联网就可以上传文件，并返回图片或者文件在服务器的URL。

上传图片首页
![20211204160053.png](https://s2.loli.net/2021/12/04/rsiKQg49je8yEM2.png)

历史图片浏览与管理 删除曾经上传图片
[![b4JqKA.png](https://s1.ax1x.com/2022/03/10/b4JqKA.png)](https://imgtu.com/i/b4JqKA)

文本复制器

[![b4YCvj.png](https://s1.ax1x.com/2022/03/10/b4YCvj.png)](https://imgtu.com/i/b4YCvj)

[![bIXEGj.png](https://s1.ax1x.com/2022/03/11/bIXEGj.png)](https://imgtu.com/i/bIXEGj)

这个主要为了便于手机和电脑同步一些碎片化的信息，如果在没有网页微信的情况下，手机和电脑共享粘贴板。



## 使用方法：

首先配置七牛存储的AK，SK值，为了安全起见，存入到本地或者服务器的环境变量

Linux

```
export qiniu_access_key='xxxxxxxxxxx'
export qiniu_secret_key='xxxxxxxxxxx'
export qiniu_bucket='bucket名字'
export qiniu_space='你的域名空间'
```

Windows:

百度如何添加环境变量即可

----
* 修改机房设置：
```bigquery
this.cfg.Zone = &storage.ZoneHuanan // 默认用的华南地区，其他地区的需要在这里改
```

* 数据库配置文件
```js
service/config.json
{
  "username": "root",
  "password": "",
  "host": "127.0.0.1",
  "port": 2306,
  "db": "db"
}
```

### 运行

```go
go run main.go
```

默认使用80端口访问。打开浏览器，输入网址：localhost即可。

* build编译
```bigquery
go build -o "qiniu"
运行：
./qiniu
```
欢迎关注交流。



