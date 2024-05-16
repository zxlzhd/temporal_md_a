# 使用官方的 Golang 镜像作为构建环境
FROM golang:1.22.3 as builder

# 设置工作目录
WORKDIR /app

# 复制 go 模块和依赖文件
COPY go.mod ./
COPY go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 构建应用
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o temporal_md_a .

# 使用 scratch 作为运行环境
FROM scratch

# 从 builder 镜像中复制构建好的应用
COPY --from=builder /app/temporal_md_a .

# 应用监听的端口
#EXPOSE 8080

# 运行应用
CMD ["./temporal_md_a"]
