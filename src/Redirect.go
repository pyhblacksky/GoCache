package main

import (
	"log"
	"net/http"
)

/**
 * @Author: pyh
 * @Date: 2019/6/24 15:18
 * @Function:	重定向测试
				输入127.0.0.1:3000/foo 后跳转为百度
*/

func main() {
	mux := http.NewServeMux() //创建一个空的 ServeMux

	rh := http.RedirectHandler("http://www.baidu.com", 307) //创建了一个重定向处理器
	mux.Handle("/foo", rh)

	log.Println("Listening..........")
	http.ListenAndServe(":3000", mux)
}
