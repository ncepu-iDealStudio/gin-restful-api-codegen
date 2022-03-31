// coding: utf-8
// @Author : lryself
// @Date : 2021/4/6 21:18
// @Software: GoLand

package settings

import (
	"LRYGoCodeGen/internal/globals/vipers"
	"LRYGoCodeGen/internal/middlewares"
	"LRYGoCodeGen/internal/routers"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func InitGinEngine() (*gin.Engine, error) {
	gin.SetMode(viper.GetString("system.Mode"))
	engine := gin.Default()

	// 加载全局中间件
	engine.Use(middlewares.CorsMiddleware())
	engine.Use(middlewares.LogMiddleware())
	engine.Use(middlewares.LoadUser())

	// 初始化Session
	// 将session存在redis
	v := vipers.GetDatabaseViper()
	store, err := redis.NewStore(10, "tcp", v.GetString("redis.addr"), v.GetString("redis.password"), []byte(viper.GetString("system.Secret")))
	if err != nil {
		return nil, err
	}
	// 将session存在cookie
	//store := cookie.NewStore([]byte(viper.GetString("system.Secret")))

	store.Options(sessions.Options{
		MaxAge: viper.GetInt("system.SessionExpireTime"),
	})
	engine.Use(sessions.Sessions("mySession", store))

	routers.InitRouter(engine)

	return engine, nil
}
