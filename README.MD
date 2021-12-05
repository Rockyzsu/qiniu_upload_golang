# 简约的七牛文件上传页面

简约且简单，只有一个首页，用于上传图片。

----

* 主要因为手机app端七牛没有上传的客户端，并且自己的电脑多操作系统，最方便的办法是通过浏览器上传文件，随时随地只要能联网就可以上传文件，并返回图片或者文件在服务器的URL。





![20211204160053.png](https://s2.loli.net/2021/12/04/rsiKQg49je8yEM2.png)



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
### 运行

```go
go run main.go
```

默认使用80端口访问。打开浏览器，输入网址：localhost即可。



欢迎关注交流。



