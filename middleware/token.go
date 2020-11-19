package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CheckToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Header.Get("token") == "test" {
			log.Println("校验 token 成功")
			// 验证通过继续执行下一个中间件
			c.Next()
		} else {
			// 验证不通过时终止请求
			c.Abort()
			log.Println("校验 token 失败")
			c.JSON(http.StatusUnauthorized, "验证失败")
		}
	}
}
