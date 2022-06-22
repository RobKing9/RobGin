package main

import (
	"log"
	"net/http"
	"time"

	"Day5-middleware/rob"
)

func onlyForV2() rob.HandlerFunc {
	return func(c *rob.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := rob.New()
	r.Use(rob.Logger()) // global midlleware
	r.GET("/", func(c *rob.Context) {
		c.HTML(http.StatusOK, "<h1>Hello RobGin</h1>")
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2()) // v2 group middleware
	{
		v2.GET("/hello/:name", func(c *rob.Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	r.Run(":9999")
}
