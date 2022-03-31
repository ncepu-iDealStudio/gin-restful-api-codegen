// coding: utf-8
// @Author : lryself
// @Date : 2021/4/8 11:36
// @Software: GoLand

package middlewares

import (
	"LRYGoCodeGen/internal/globals/codes"
	"LRYGoCodeGen/internal/globals/database"
	"LRYGoCodeGen/internal/globals/jwt"
	logs "gitee.com/lryself/go-utils/loggers"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

var log = logs.GetLogger()

func TokenRequire() gin.HandlerFunc {
	return func(c *gin.Context) {
		//token验证
		token := c.Request.Header.Get("Token")
		jwtChaim, err := jwt.VerifyToken(token, []byte(viper.GetString("system.Secret")))
		if err != nil {
			log.Errorln(err)
			c.JSON(http.StatusOK, gin.H{
				"code":    codes.AccessDenied,
				"message": "您的Token已过期！",
			})
			c.Abort()
			return
		}

		//从数据库读取token信息
		redisManager := database.GetRedisManager()
		result, err := redisManager.Get("Token_" + jwtChaim.UserID).Result()
		if err != nil || result != token {
			log.Errorln(err)
			c.JSON(http.StatusOK, gin.H{
				"code":    codes.AccessDenied,
				"message": "您的Token已失效！",
			})
			c.Abort()
			return
		}

		//刷新token有效期
		err = redisManager.Expire("Token_"+jwtChaim.UserID, time.Duration(viper.GetInt("system.RedisExpireTime"))*time.Second).Err()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    codes.InternalError,
				"message": "刷新token错误！",
				"err":     err,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
