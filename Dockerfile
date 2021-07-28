FROM golang:alpine

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 移动到工作目录：/build
WORKDIR /build

#将项目复制到容器中
COPY . .

# 将我们的代码编译成二进制可执行文件 bubble
RUN go build -o web_app .

