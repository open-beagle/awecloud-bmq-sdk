# awecloud-bmq-sdk

## proto

```bash
# install
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# update
protoc \
  --go_out=. \
  --go_opt=paths=source_relative \
  --go-grpc_out=. \
  --go-grpc_opt=paths=source_relative \
  pkg/registry.proto
```

## router

grpc使用以下路由提供服务
/awecloud/bmq/registry

```bash
# 修改前
err := c.cc.Invoke(ctx, "/Registry/Login", in, out, opts...)

# 修改后 , GRPC默认支持/pkg.Registry/Login，这样的Url前缀，下面的URL前缀需要手动修改。
err := c.cc.Invoke(ctx, "/awecloud/bmq/registry/Login", in, out, opts...)
```

## tag

```bash
git tag v0.1.1 -f
git push origin v0.1.1 -f
```
