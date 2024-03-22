package main

import (
	"log"
	"net/http"
	"time"
	
	"github.com/888go/gin"
	"golang.org/x/sync/errgroup"
)

var g errgroup.Group

func router01() http.Handler {
	e := gin类.X创建()
	e.X中间件(gin类.Recovery())
	e.X绑定GET("/", func(c *gin类.Context) {
		c.X输出JSON(
			http.StatusOK,
			gin类.H{
				"code":  http.StatusOK,
				"error": "Welcome server 01",
			},
		)
	})

	return e
}

func router02() http.Handler {
	e := gin类.X创建()
	e.X中间件(gin类.Recovery())
	e.X绑定GET("/", func(c *gin类.Context) {
		c.X输出JSON(
			http.StatusOK,
			gin类.H{
				"code":  http.StatusOK,
				"error": "Welcome server 02",
			},
		)
	})

	return e
}

func main() {
	server01 := &http.Server{
		Addr:         ":8080",
		Handler:      router01(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server02 := &http.Server{
		Addr:         ":8081",
		Handler:      router02(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		return server01.ListenAndServe()
	})

	g.Go(func() error {
		return server02.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
