# go-grcp-demo

a grpc demo using golang

该demo主要演示client和server的创建以及通讯案例

ps:生成go文件可在项目主目录使用下面的命令

```go
protoc -I = proto/ proto/rpc.proto --go_out = plugins = grpc:proto
```