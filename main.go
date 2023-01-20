package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/shashank/ginproject/httpd/handler"
	"github.com/shashank/ginproject/platform/newsfeed"
)

func main() {
	r := gin.Default()

	r.GET("/ping", handler.PingGet())
	r.GET("/newsfeed", handler.NewsFeedGet())
	r.POST("/newsfeed", handler.NewsFeedPost())

	r.Run()

	feed := newsfeed.New()

	fmt.Println(feed)

	feed.Add(newsfeed.Item{"Hello", "How ya doing mate"})

	fmt.Println(feed)
}
