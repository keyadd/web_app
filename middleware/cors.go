package middleware

import (
	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//method := c.Request.Method
		//origin := c.Request.Header.Get("Origin")
		//var filterHost = [...]string{"http://localhost.*"}
		// filterHost 做过滤器，防止不合法的域名访问
		//var isAccess = false
		//for _, v := range filterHost {
		//	match, _ := regexp.MatchString(v, origin)
		//	if match {
		//		isAccess = true
		//	}
		//}
		//if isAccess {
		// 核心处理方式
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Allow-Methods", "*")
		c.Set("content-type", "application/json")
		//}
		//放行所有OPTIONS方法
		//if method == "OPTIONS" {
		//	c.AbortWithStatus(http.StatusNoContent)
		//}

		c.Next()
	}
}
