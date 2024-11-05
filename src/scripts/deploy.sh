#!/bin/bash

# 部署脚本

# 构建 Go 项目
echo "Building Go project..."
go build -o GameBackendProject ./src/go/main.go

# 启动服务
echo "Starting server..."
./GameBackendProject &

# 保存进程 ID
echo $! > server.pid

echo "Deployment complete."