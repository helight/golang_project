package main

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

// 限流
func RateLimiter() gin.HandlerFunc {
	//例如： 每秒产生1个令牌，最多存储10个令牌。
	l := rate.NewLimiter(1, 10)
	return func(context *gin.Context) {
		//当没有可用的令牌时返回false，也就是当没有可用的令牌时，禁止通行
		if !l.Allow() {
			context.JSON(429, gin.H{"error": "request limit"})
			context.Abort()
		}
		//用可用的令牌时放行
		context.Next()
	}
}

func main() {

	fmt.Printf("launch Gin")

	Router := gin.Default()
	PublicGroup := Router.Group("xxx").Use(RateLimiter())
	PublicGroup.GET("/get", HandleGet)
	PublicGroup.POST("/getall", HandleGetAllData)

	//如果使用浏览器调试，那么响应Get方法
	//r.GET("/getall",HandleGetAllData)
	Router.Run(":9000")
}

func HandleGet(c *gin.Context) {
	fmt.Println("---header/--- \r\n")
	for k, v := range c.Request.Header {
		fmt.Println(k, v)
	}
	headerTest := c.Request.Header.Get("X-Real-Ip")
	fmt.Printf("X-Real-Ip:%s;", headerTest)

	c.JSON(200, gin.H{
		"receive": "65536",
	})

}

func HandleGetAllData(c *gin.Context) {
	//log.Print("handle log")
	body, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Println("---body/--- \r\n " + string(body))

	fmt.Println("---header/--- \r\n")
	for k, v := range c.Request.Header {
		fmt.Println(k, v)
	}
	headerTest := c.Request.Header.Get("test")

	fmt.Printf("header_test:%s;", headerTest)

	//fmt.Println("header \r\n",c.Request.Header)

	c.JSON(200, gin.H{
		"receive": "1024",
	})

}
