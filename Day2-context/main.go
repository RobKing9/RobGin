/*
	将路由(router)独立出来，方便之后增强。
	设计上下文(Context)，封装 Request 和 Response ，提供对 JSON、HTML 等返回类型的支持。
*/

package main

import (
	"fmt"
	"net/http"

	"Day2-context/gee"
)

func main() {
	r := gee.New()
	r.Get("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	r.Get("/hello", func(c *gee.Context) {
		// expect /hello?name=geektutu
		name := c.Query("name")
		c.String(http.StatusOK, "hello %s, you're at %s\n", name, c.Path)
	})

	r.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	fmt.Println("正在监听9999端口！！！")
	r.Run(":9999")
}
