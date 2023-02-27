// coding: utf-8
// @Author : lryself
// @Date : 2022/4/13 21:20
// @Software: GoLand

package rpcServer

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
	"tem_go_project/utils/errHelper"

	"tem_go_project/internal/rpcServer/middlewares"
	"tem_go_project/utils/loggers"
	"tem_go_project/utils/message"
)

type myLoggerV2 struct {
	*logrus.Logger
	v int
}

func (l myLoggerV2) V(i int) bool {
	return i > l.v
}
func (l *myLoggerV2) init(v int) {
	l.v = v
	l.Logger = loggers.GetLogger()
}

func StartRPCEngine() {
	rpcServer := grpc.NewServer(grpc_middleware.WithUnaryServerChain(
		middlewares.GrpcContext(),
		middlewares.GrpcRecover(),
		middlewares.GrpcLogger(),
	))
	var log myLoggerV2
	log.init(-1)
	grpclog.SetLoggerV2(&log)
	// 2. 注册服务
	registerServer(rpcServer)
	// 3. 新建一个listener，以tcp方式监听端口
	listener, err := net.Listen("tcp", ":"+viper.GetString("system.RpcPort"))
	errHelper.ErrExit(err)

	// 4. 运行rpcServer，传入listener
	message.Println("rpc服务已启动")
	_ = rpcServer.Serve(listener)
	message.PrintWarn("rpc服务关闭")
}

func registerServer(rpcServer *grpc.Server) {
	//pb.RegisterTokenServiceServer(rpcServer, new(service.TokenService))
}
