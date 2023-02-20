FROM golang:1.19-alpine as builder

ENV GO111MODULE=on \
GOPROXY=https://goproxy.cn,direct

WORKDIR /app

COPY . .
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o gateway_service ./gateway/cmd

FROM alpine as prod

WORKDIR /app

COPY --from=builder /app/conf/conf.yml .
COPY --from=builder /app/gateway_service .

# 指定运行时环境变量
ENV GIN_MODE=release

EXPOSE 8070

ENTRYPOINT ["./gateway_service"]