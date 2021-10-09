package main

import (
	// "errors"
	"net/http"

	"github.com/gomodule/redigo/redis"
	"github.com/labstack/echo/v4"
)

var pool *redis.Pool

type UrlObject struct {
	Name string `redis:"name"`
	UrlLong string `redis:"url_long"`
	URLShort string `redis:"url_short"`
	Description string `redis:"description"`
}

// e.POST("/new", new)
func addUrl(c echo.Context)  error {


	// Get name and email
	name := c.FormValue("name")
	source_url  := c.FormValue("url")
	short_url	:= generateShortUrl(name, source_url)
	//	target_url	:= # C
	return c.String(http.StatusOK, "name:" + name + ", url:" + source_url + ", short:" + short_url)
}

func generateShortUrl(name string, url string)(string) {
	conn := pool.Get()
	defer conn.Close()

	conn.Do("SET", "foo", "bar")

	return "ok"
}
