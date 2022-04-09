// coding: utf-8
// @Author : lryself
// @Date : 2021/4/8 11:36
// @Software: GoLand

package middlewares

import (
	"LRYGoCodeGen/internal/globals/database"
	"LRYGoCodeGen/internal/globals/jwt"
	"LRYGoCodeGen/internal/globals/parser"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"time"
)

func TokenRequire() gin.HandlerFunc {
	return func(c *gin.Context) {
		//token验证
		token := c.Request.Header.Get("Token")
		jwtChaim, err := jwt.VerifyToken(token, []byte(viper.GetString("system.Secret")))
		if err != nil {
			parser.JsonAccessDenied(c, "请重新登录！")
			c.Abort()
			return
		}

		//从数据库读取token信息
		redisManager := database.GetRedisManager()
		result, err := redisManager.Get("Token_" + jwtChaim.UserID).Result()
		if err != nil || result != token {
			parser.JsonAccessDenied(c, "请重新登录！")
			c.Abort()
			return
		}

		//刷新token有效期
		err = redisManager.Expire("Token_"+jwtChaim.UserID, time.Duration(viper.GetInt("system.RedisExpireTime"))*time.Second).Err()
		if err != nil {
			parser.JsonInternalError(c, "刷新token错误！", err)
			c.Abort()
			return
		}
		c.Next()
	}
}
