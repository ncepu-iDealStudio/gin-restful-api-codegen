# 使用命令
protoc -I assets/proto/ --go_out=plugins=grpc:internal/rpcServer/pb UserService.proto