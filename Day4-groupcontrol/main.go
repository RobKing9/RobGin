package main

import (
	"net/http"

	"Day4-groupcontrol/rob"
)

func main() {
	r := rob.New()
	r.GET("/index", func(c *rob.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *rob.Context) {
			c.HTML(http.StatusOK, "<h1>Hello RobGin</h1>")
		})

		v1.GET("/hello", func(c *rob.Context) {
			// expect /hello?name=geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *rob.Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *rob.Context) {
			c.JSON(http.StatusOK, rob.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}

	r.Run(":9999")
}
