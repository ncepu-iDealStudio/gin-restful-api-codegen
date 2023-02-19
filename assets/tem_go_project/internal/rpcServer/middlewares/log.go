package middlewares

import (
	"context"
	"gitee.com/lryself/go-utils/loggers"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"time"
)

func GrpcLogger() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		startTime := time.Now()
		defer func() {
			remoteCon, ok := peer.FromContext(ctx)
			if !ok {
				remoteCon.Addr.String()
			}
			md, ok := metadata.FromIncomingContext(ctx)
			if !ok {
				md.Len()
			}

			log := loggers.GetLogger()
			if err != nil {
				log.WithFields(logrus.Fields{
					"client_ip":    remoteCon.Addr.String(),
					"req_method":   info.FullMethod,
					"req_string":   req,
					"error":        err.Error(),
					"latency_time": time.Since(startTime),
				}).Error()
			} else {
				log.WithFields(logrus.Fields{
					"client_ip":    remoteCon.Addr.String(),
					"req_method":   info.FullMethod,
					"req_string":   req,
					"res_string":   resp,
					"latency_time": time.Since(startTime),
				}).Info()
			}
		}()
		resp, err = handler(ctx, req)
		return resp, err
	}
}
