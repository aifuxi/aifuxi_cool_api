# 构建镜像阶段
FROM golang:1.21.0-alpine3.18 AS builder
# 设置工作区
WORKDIR /workspace
# 拷贝当前目录代码到工作区
COPY . .
# 设置 go 代理，不然安装依赖很慢，go build的时候会自动安装依赖
ENV GOPROXY https://goproxy.cn
# 执行构建命令 -o 指定输出文件名
RUN go build -o main main.go

# 运行阶段
FROM alpine:3.18
# 设置工作区
WORKDIR /workspace
# 从构建镜像中拷贝可执行程序到运行容器
COPY --from=builder /workspace/main .
COPY --from=builder /workspace/configs ./configs
# 对外暴露容器端口
EXPOSE 9003
# 容器一启动时自动执行的命令
CMD [ "/workspace/main" ]

