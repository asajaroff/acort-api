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
	// Initialize a connection pool and assign it to the pool global
	// variable.
	time.Now()
	pool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:49153")
		},
		MaxIdle:         10,
		MaxActive:       0,
		IdleTimeout:     240 * time.Second,
		Wait:            false,
		MaxConnLifetime: 0,
	}

	e := echo.New()
	e.POST("/add", addUrl)
	e.Logger.Fatal(e.Start(":8080"))
}
