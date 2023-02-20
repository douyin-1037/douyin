FROM golang:1.19-alpine as builder

ENV GO111MODULE="on" \
GOARCH="amd64" \
CGO_ENABLED="1" \
GOOS="linux" \
GOPROXY="https://goproxy.io,direct"

WORKDIR /app

COPY . .

RUN go build -o gateway ./gateway/cmd

FROM alpine as prod

WORKDIR /app

COPY --from=builder /app/conf/conf.yaml .
COPY --from=builder /app/gateway .

# 指定运行时环境变量
ENV GIN_MODE=release

EXPOSE 8070

ENTRYPOINT ["./gateway"]