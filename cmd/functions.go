package main

import (
	// "errors"
	"fmt"
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
	data := echo.Map{}
    if err := c.Bind(&data); err != nil {
        return err
    } else {
        name := fmt.Sprintf("%v", data["name"])
        url := fmt.Sprintf("%v", data["url"])
		short_url	:= generateShortUrl(name, url)
		return c.String(http.StatusOK, "name:" + name + ", url:" + url + ", short:" + short_url)
    }

}

func generateShortUrl(name string, url string)(string) {
	conn := pool.Get()
	defer conn.Close()

	conn.Do("SET", name, url)
	fmt.Println("Added ", name, url)
	return name
}

// func getItem(name string){
// 	conn := pool.Get()
// 	defer conn.Close()
// 	conn.Do("GET", "name")
// 	return "ok"
// }
