# 使用命令
protoc -I asserts/proto/ --go_out=plugins=grpc:internal/rpcServer/pb UserService.proto