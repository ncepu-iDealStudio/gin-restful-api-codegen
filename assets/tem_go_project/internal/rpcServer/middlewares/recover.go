package middlewares

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"runtime/debug"
	"tem_go_project/utils/loggers"
)

func GrpcRecover() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		defer func() {
			log := loggers.GetLogger()
			if e := recover(); e != nil {
				log.Errorln("server panic", e)
				log.Errorf("%s\n", debug.Stack())
				err = errors.New(fmt.Sprintf("panic:%v", e))
			}
		}()
		resp, err = handler(ctx, req)
		return resp, err
	}
}
