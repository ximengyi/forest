# Compile stage
FROM golang:1.17.8-alpine3.15 AS buildStage


# The latest alpine images don't have some tools like (`git` and `bash`).
# Adding git, bash and openssh to the image
RUN sed -i "s/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g" /etc/apk/repositories
RUN apk add --no-cache git make bash ca-certificates tzdata 

# 镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOSUMDB="sum.golang.google.cn" \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.cn,direct" \
    TZ=Asia/Shanghai \
    APP_ENV=docker 


# 移动到工作目录
WORKDIR /go/src/forest
ADD . /go/src/forest/
RUN cd /go/src/forest/ && go build -o forest


# 打包阶段
FROM alpine:latest
RUN sed -i "s/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g" /etc/apk/repositories
RUN apk add --no-cache  ca-certificates tzdata 
RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo 'Asia/Shanghai' >/etc/timezone

WORKDIR /app
COPY --from=buildStage /go/src/forest  /app/
RUN mkdir -p /opt/logs/ && chmod 777 /opt/logs
# ENTRYPOINT ./translate
CMD ["/app/forest"]


# 3. test: curl -i http://localhost:8700/healthy

#docker run --rm -it -p 8700:8700 