#=============================== 构建阶段 ===============================
FROM registry.cn-shenzhen.aliyuncs.com/aliensidea/golib:latest AS aliens-build

ENV SOURCE src

# 设置我们应用程序的工作目录
WORKDIR /go/$SOURCE

# 添加所有需要编译的应用代码
COPY ./$SOURCE  .

# 编译一个静态的go应用（在二进制构建中包含C语言依赖库）
RUN CGO_ENABLED=0 GOOS=linux go build -v  -a -installsuffix cgo -o ./server

#=============================== 生产阶段 ===============================

FROM scratch AS aliens-prod

ENV SOURCE src

#作者
MAINTAINER aliensboot

# 从build阶段拷贝二进制文件
COPY --from=aliens-build /go/$SOURCE/server .

ADD ./config ./config

CMD ["./server"]