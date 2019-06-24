package main

import (
	"io"
	"log"
	"net/http"
)

/**
 * @Author: pyh
 * @Date: 2019/6/24 15:14
 * @Function:	测试go自带的HTTP包功能	(我的是访问127.0.0.1:9090)
 * 				运行main函数后  访问本机ip:9090  则打印Hello, this is Go Server!
 */

func HelloGoServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, this is Go Server!")
}

func main() {
	http.HandleFunc("/", HelloGoServer)
	err := http.ListenAndServe(":9090", nil) //监听并服务9090端口
	if err != nil {
		log.Fatal("ListenAndServe ", err)
	}
}
