// coding: utf-8
// @Author : lryself
// @Date : 2022/4/13 22:06
// @Software: GoLand

package rpcReq

import (
	"context"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"tem_go_project/internal/rpcServer/pb"
	"time"
)

func VerifyToken(token string) (*pb.VerifyTokenResponse, error) {
	conn, err := grpc.Dial(viper.GetString("remote.SsoCenter"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	// 延迟关闭连接
	defer conn.Close()

	// 初始化客户端
	tokenService := pb.NewTokenServiceClient(conn)

	// 初始化上下文，设置请求超时时间为1秒
	//ctx := context.Background()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// 延迟关闭请求会话
	defer cancel()

	// 调用接口
	r, err := tokenService.VerifyToken(ctx, &pb.VerifyTokenRequest{Token: token})
	if err != nil {
		return nil, err
	}

	// 返回消息
	return r, nil
}

func FreeToken(token string) error {
	conn, err := grpc.Dial(viper.GetString("remote.UserCenter"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	// 延迟关闭连接
	defer conn.Close()

	// 初始化客户端
	tokenService := pb.NewTokenServiceClient(conn)

	// 初始化上下文，设置请求超时时间为1秒
	//ctx := context.Background()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// 延迟关闭请求会话
	defer cancel()

	// 调用接口
	_, err = tokenService.FreeToken(ctx, &pb.VerifyTokenRequest{Token: token})
	return err
}
