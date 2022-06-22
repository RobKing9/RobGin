/*
	Gee框架的原型
	实现了路由映射表
	提供了用户注册静态路由的方法
	包装了启动服务的函数
*/
package main

import (
	"Gee/rob"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := rob.New()
	r.Get("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.path = %q\n", req.URL.Path)
	})

	r.Get("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})
	fmt.Println("正在监听9999端口！！！")
	log.Fatal(r.Run(":9999"))
}
