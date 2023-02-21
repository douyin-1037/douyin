FROM golang:1.19-alpine as builder

ENV GO111MODULE=on \
GOPROXY=https://goproxy.cn,direct

WORKDIR /douyin

COPY . .
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o message_service ./message

FROM alpine as prod

WORKDIR /douyin

COPY --from=builder /douyin/conf/conf.yml ./conf/conf.yml
COPY --from=builder /douyin/message_service .

# 指定运行时环境变量
ENV GIN_MODE=release

EXPOSE 8074

ENTRYPOINT ["./message_service"]