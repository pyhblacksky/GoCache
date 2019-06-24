# GoCache
Use go to realize simple cache


## 实现三个缓存服务接口

**PUT**

PUT /cache/<key>
content
<value>



**GET**

GET /cache/<key>
content
<value>



**DEL**

DELETE /cache/<key>



------

## 使用Golang的HTTP包自带的包用作常用处理器

​	Golang中处理 HTTP 请求主要跟两个东西相关：**ServeMux** 和 **Handler**。

​	ServrMux 本质上是一个 HTTP 请求路由器（或者叫多路复用器，Multiplexor）。它把收到的请求与一组预先定义的 URL 路径列表做对比，然后在匹配到路径的时候调用关联的处理器（Handler）。

​	处理器（Handler）负责输出HTTP响应的头和正文。任何满足了http.Handler接口的对象都可作为一个处理器。通俗的说，对象只要有个如下签名的ServeHTTP方法即可：

```go
ServeHTTP(http.ResponseWriter, *http.Request)
```

​	Golang的 HTTP 包自带了几个函数用作常用处理器，比如NotFoundHandler 和 RedirectHandler。
​	NotFoundHandler返回一个简单的请求处理器，该处理器会对每个请求都回复"404 page not found"。

​	RedirectHandler返回一个请求处理器，该处理器会对每个请求都使用状态码code重定向到网址url。



------

## 测试结果

![PUT](https://github.com/pyhblacksky/GoCache/blob/master/doc/put.PNG)



![GET](https://github.com/pyhblacksky/GoCache/blob/master/doc/get.PNG)
