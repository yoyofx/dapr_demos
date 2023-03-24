FROM golang:1.18-alpine AS builder
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn,direct
WORKDIR /publish
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o app .
# 最终容器, 基础镜像alpine 中包含 tini
FROM alpine:latest as production
COPY --from=builder /publish/app /
RUN chmod a-x app && chmod a+x app
# Run
CMD ["/app"]