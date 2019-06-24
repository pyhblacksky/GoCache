package main

import (
	"./cache"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

/**
 * @Author: pyh
 * @Date: 2019/6/24 15:25
 * @Function:	实现简单缓存功能
 */

type cacheHandler struct {
	cache.Cache
}

//CacheHandler
func (h *cacheHandler) CacheHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("url:", r.URL, " Method : ", r.Method)

	//split
	key := strings.Split(r.URL.EscapedPath(), "/")[2]

	if len(key) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	m := r.Method

	if m == http.MethodPut {
		h.HandlePut(key, w, r)
		return
	} else if m == http.MethodGet {
		h.HandleGet(key, w, r)
		return
	} else if m == http.MethodDelete {
		h.HandleDelete(key, w, r)
		return
	}

	w.WriteHeader(http.StatusMethodNotAllowed)
}

//HandlePut
func (h *cacheHandler) HandlePut(k string, w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadAll(r.Body)

	//内容不为空
	if len(b) != 0 {
		e := h.Set(k, b)
		if e != nil {
			//此时说明出错
			log.Println(e)
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.Write([]byte("successful by pyh"))
		}
	}
}

//HandleGet
func (h *cacheHandler) HandleGet(k string, w http.ResponseWriter, r *http.Request) {
	b, e := h.Get(k)

	if e != nil {
		//出错
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(b) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Write(b)
}

//HandleDelete
func (h *cacheHandler) HandleDelete(k string, w http.ResponseWriter, r *http.Request) {
	e := h.Del(k)

	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Write([]byte("successful  by pyh"))
	}
}

//主程序,监听网址 127.0.0.1:25000/cache/
//使用postman测试
func main() {
	c := cache.New("inmemory")
	h := cacheHandler{c}
	http.HandleFunc("/cache/", h.CacheHandler)
	http.ListenAndServe(":25000", nil)
}
