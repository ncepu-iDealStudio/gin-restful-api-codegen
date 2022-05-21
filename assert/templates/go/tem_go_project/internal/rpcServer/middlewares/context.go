// coding: utf-8
// @Author : lryself
// @Date : 2022/5/17 14:05
// @Software: GoLand

package middlewares

import (
	"context"
	"google.golang.org/grpc"
)

type contextKey struct{}

type Context struct {
	context.Context
}

var cKey = contextKey{}

func storeContext(c context.Context, ctx *Context) context.Context {
	return context.WithValue(c, cKey, ctx)
}

func GetContext(c context.Context) *Context {
	return c.Value(cKey).(*Context)
}

func NewContext(ctx context.Context) *Context {
	c := &Context{}
	c.Context = storeContext(ctx, c)
	return c
}

func GrpcContext() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		commonCtx := NewContext(ctx)
		return handler(commonCtx, req)
	}
}
