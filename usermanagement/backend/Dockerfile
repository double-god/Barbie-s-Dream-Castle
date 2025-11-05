# ----- 阶段 1: 构建器 (Builder) -----
# 我们使用一个包含完整 Go 编译环境的 "builder" 镜像
# alpine 标签是一个非常轻量的 Linux，能让下载更快
# ----- 
FROM golang:1.25-alpine AS builder


# 在容器内设置工作目录
WORKDIR /app

# 1. 单独复制 go.mod 和 go.sum 文件
# 这一步是为了利用 Docker 的层缓存。只要 go.mod 不变，Docker 就不会重新下载依赖。
COPY go.mod go.sum ./

# 2. 下载所有依赖
RUN go mod download

# 3. 复制项目中的所有源代码
COPY . .

# 4. 编译 Go 应用程序
# CGO_ENABLED=0 告诉编译器构建一个静态二进制文件，不依赖任何系统的 C 库
# GOOS=linux 确保我们编译的是 Linux 程序 (因为你的 ECS 服务器是 Linux)
# -o /app/main 指定编译输出的可执行文件名为 main，并放在 /app 目录下
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main ./main.go

# ----- 阶段 2: 最终镜像 (Final) -----
# 我们使用一个“极其”轻量的 alpine 镜像作为最终的运行环境。
# 它甚至不包含 Go 语言本身，只包含最基础的 Linux 系统。
FROM alpine:latest

# 同样设置工作目录
WORKDIR /app

# 1. (关键步骤) 从 "builder" 阶段复制编译好的程序
# 我们只复制那个编译好的 "main" 文件，所有的 Go 源代码、工具链都被丢弃了
# 这就是为什么最终镜像会非常小的原因 (通常只有 10-20MB)
COPY --from=builder /app/main .

# 2. 暴露端口
# 告诉 Docker，这个容器打算在 8080 端口上提供服务。
# 这对应你 .env 文件里的 HTTP_PORT=8080
EXPOSE 8080

# 3. 容器启动命令
# 当容器启动时，执行 ./main 这个程序
CMD ["./main"]