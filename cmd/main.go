package main

import (
	//	"fmt"
	//"log"
	//"net/http"
	//"os"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/labstack/echo/v4"
)

func main(){
	time.Now()
	// Initialize a connection pool and assign it to the pool global
	// variable.
	pool = &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}

	e := echo.New()

	e.POST("/add", addUrl)

	e.Logger.Fatal(e.Start(":8080"))
}
