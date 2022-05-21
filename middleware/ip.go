package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"net/http"
	"strings"
	"time"
)

var (
	RedisPool *redis.Pool
	redisHost = "119.23.40.47:6379"
)

func RobotRestrict() gin.HandlerFunc {
	/*
		对恶意请求的IP进行限制
	*/
	return func(c *gin.Context) {
		var (
			count = 55 // 频次数
			cycle = 60 // 时间周期 单位（second）
		)
		ip := strings.Split(c.Request.RemoteAddr, ":")[0]
		conn := RedisPool.Get()
		rep, _ := redis.String(conn.Do("Get", ip))
		if rep == "" {
			_, err := conn.Do("setex", ip, cycle, count) //conn.Do()用的是最多的，把命令行中的args一个个写进去
			if err != nil {
				c.JSON(
					http.StatusOK,
					gin.H{
						"code": 405,
						"msg":  "server error",
					},
				)
				c.Abort() //终止请求，直接返回提示信息
				return
			}
		}
		if rep == "1" {
			c.JSON(
				http.StatusOK,
				gin.H{
					"code": 405,
					"msg":  "客官慢一些",
				},
			)
			c.Abort() //终止请求，直接返回提示信息
			return
		}
		conn.Do("DECR", ip)
		defer conn.Close()
	}
}

func newRedisPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     50,                // 最大空闲连接数
		MaxActive:   30,                //允许分配最大连接数
		IdleTimeout: 300 * time.Second, // 连接时间限制
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisHost)
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}
}

func init() {
	//创建redis连接引擎
	RedisPool = newRedisPool()
}
