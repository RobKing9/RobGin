package main

import (
	"Day7-errorRecovery/rob"
	"net/http"
)

func main() {
	r := rob.Default()
	r.GET("/", func(c *rob.Context) {
		c.String(http.StatusOK, "Hello Geektutu\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *rob.Context) {
		names := []string{"geektutu"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":9999")
}
