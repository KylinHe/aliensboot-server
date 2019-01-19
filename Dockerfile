#=============================== 构建阶段 ===============================
FROM golang:1.11.2 AS aliens-build

ENV SOURCE src/github.com/KylinHe/aliensboot-server
ENV GO111MODULE on
ENV GOPATH /root/go

#使用本地依赖包，减少下载翻墙的时间, mac因为和docker demon是独立的可能需要把依赖的包放入到上下文目录下拷贝过去
#COPY ./_temp_/mod /root/go/pkg/mod

# 设置我们应用程序的工作目录
WORKDIR /go/$SOURCE

# 添加所有需要编译的应用代码
COPY ./$SOURCE  .

# 编译一个静态的go应用（在二进制构建中包含C语言依赖库）
RUN CGO_ENABLED=0 GOOS=linux go build -v  -a -installsuffix cgo -o ./server

#=============================== 生产阶段 ===============================

FROM scratch AS aliens-prod

MAINTAINER aliensboot

# 从buil阶段拷贝二进制文件
COPY --from=aliens-build /go/$SOURCE/server .

ADD ./config ./config

CMD ["./server"]