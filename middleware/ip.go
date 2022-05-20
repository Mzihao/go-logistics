package middleware

import (
	"github.com/gin-gonic/gin"
	"time"
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
		conn := dao.RedisPool.Get()
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
————————————————
版权声明：本文为CSDN博主「乐享极致」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/weixin_44313745/article/details/115404222