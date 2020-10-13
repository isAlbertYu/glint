package main

import (
	"go-glint/glint"
	"net/http"
)

func main() {
	r := glint.New()
	r.GET("/", func(c *glint.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	r.GET("/hello", func(c *glint.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *glint.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.GetRouteParam("name"), c.Path)
	})

	//r.GET("/assets/*filepath", func(c *glint.Context) {
	//	c.JSON(http.StatusOK, glint.H{"filepath": c.GetRouteParam("filepath")})
	//})

	r.GET("/assets/:name/*filepath", func(c *glint.Context) {
		c.JSON(http.StatusOK, glint.H{"name": c.GetRouteParam("name"), "filepath": c.GetRouteParam("filepath")})
	})

	r.Run(":9999")
}
