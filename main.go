package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var gconfig = GetConfig()

func main() {
	var mysql Database
	var redis URedis
	mysql.Open(gconfig)
	redis.Open(gconfig)

	if mysql.Enable {
		defer mysql.Connection.Close()
	}

	if redis.Enable {
		defer redis.Client.Close()
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"status": true, "message": ""})
	})

	router.POST("/short", func(context *gin.Context) {
		token := context.GetHeader("token")
		find := false

		for _, t := range gconfig.Token {
			if t == token {
				find = true
			}
		}

		if !find {
			context.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "token mismatch."})
			return
		}

		long := context.PostForm("long")

		if long == "" {
			context.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "long can not null."})
			return
		}

		var random string

		for {
			random = Generation(7)
			if _, err := GetLong(&mysql, &redis, random); err != nil {
				break
			}
		}

		short := context.PostForm("short")

		if short != "" {
			if _, err := GetLong(&mysql, &redis, short); err != nil {
				random = short
			} else {
				context.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "short already exists."})
				return
			}
		}

		if err := Set(&mysql, &redis, random, long); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error()})
			return
		}

		context.JSON(http.StatusOK, gin.H{"status": true, "message": "", "short": fmt.Sprintf("%s%s", gconfig.Host, random)})
	})

	router.GET("/:short", func(context *gin.Context) {
		short := context.Param("short")
		long, err := GetLong(&mysql, &redis, short)

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error()})
		} else {
			context.Redirect(http.StatusMovedPermanently, long)
		}

	})

	router.Run(gconfig.Listen)
}
