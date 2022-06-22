package main

import (
	"Day3-trie/rob"
	"net/http"
)

func main() {
	r := rob.New()
	r.Get("/", func(c *rob.Context) {
		c.HTML(http.StatusOK, "<h1>Hello RobGin</h1>")
	})

	r.Get("/hello", func(c *rob.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.Get("/hello/:name", func(c *rob.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.Get("/assets/*filepath", func(c *rob.Context) {
		c.JSON(http.StatusOK, rob.H{"filepath": c.Param("filepath")})
	})

	r.Run(":9999")
}
