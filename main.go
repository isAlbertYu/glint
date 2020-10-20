package main

import (
	"go-glint/glint"
	"net/http"
)

func main() {
	myEngine := glint.GetSingleEngine()
	myEngine.GET("/index", func(c *glint.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})

	v1 := glint.NewGroup("/v1")
	{
		v1.GET("/", func(c *glint.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Glint</h1>")
		})
		//v1.GET("/hello", func(c *glint.Context) {
		//	c.String(http.StatusOK, "hello %s, this is %s\n", c.Query("name"), c.Path)
		//})
	}
	v2 := glint.NewGroup("/v2")
	{
		v2.GET("/hello/:name", func(c *glint.Context) {
			c.String(http.StatusOK, "hello %s, this is %s\n", c.GetRouteParam("name"), c.Path)
		})
		v2.POST("/login", func(c *glint.Context) {
			c.JSON(http.StatusOK, glint.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})
	}
	// curl localhost:9999/v2/login -X POST -d "username=world&password=123"
	myEngine.Run(":9999")
}
