// coding: utf-8
// @Author : lryself
// @Date : 2022/4/13 21:23
// @Software: GoLand

package service

import (
	"context"
	"errors"
	"github.com/spf13/viper"
	"tem_go_project/internal/globals/database"
	"tem_go_project/internal/globals/jwt"
	"tem_go_project/internal/rpcServer/pb"
	"time"
)

type TokenService struct {
}

func (ps *TokenService) VerifyToken(ctx context.Context, request *pb.VerifyTokenRequest) (*pb.VerifyTokenResponse, error) {
	//token验证
	token := request.GetToken()
	jwtChaim, err := jwt.VerifyToken(token, []byte(viper.GetString("system.Secret")))
	if err != nil {
		return nil, err
	}

	//从数据库读取token信息
	redisManager := database.GetRedisManager()
	result, err := redisManager.HGetAll("Token_" + jwtChaim.TokenID).Result()
	if err != nil || result["token"] != token {
		return nil, errors.New("您的Token已失效！")
	}

	//刷新token有效期
	err = redisManager.Expire("Token_"+jwtChaim.TokenID, time.Duration(viper.GetInt("system.RedisExpireTime"))*time.Second).Err()
	if err != nil {
		return nil, errors.New("刷新token失败！")
	}

	return &pb.VerifyTokenResponse{
		UserID: result["userID"],
	}, nil
}

func (ps *TokenService) FreeToken(ctx context.Context, request *pb.VerifyTokenRequest) (*pb.FreeTokenResponse, error) {
	//token验证
	token := request.GetToken()
	jwtChaim, err := jwt.VerifyToken(token, []byte(viper.GetString("system.Secret")))
	if err != nil {
		return nil, err
	}

	//从数据库读取token信息
	redisManager := database.GetRedisManager()
	err = redisManager.Del("Token_" + jwtChaim.TokenID).Err()
	if err != nil {
		return nil, errors.New("登出token失败！")
	}

	return &pb.FreeTokenResponse{}, nil
}
