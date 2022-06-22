/*
	将路由(router)独立出来，方便之后增强。
	设计上下文(Context)，封装 Request 和 Response ，提供对 JSON、HTML 等返回类型的支持。
*/

package main

import (
	"fmt"
	"net/http"

	"Day2-context/rob"
)

func main() {
	r := rob.New()
	r.Get("/", func(c *rob.Context) {
		c.HTML(http.StatusOK, "<h1>Hello RobGin</h1>")
	})
	r.Get("/hello", func(c *rob.Context) {
		// expect /hello?name=geektutu
		name := c.Query("name")
		c.String(http.StatusOK, "hello %s, you're at %s\n", name, c.Path)
	})

	r.POST("/login", func(c *rob.Context) {
		c.JSON(http.StatusOK, rob.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	fmt.Println("正在监听9999端口！！！")
	r.Run(":9999")
}
