FROM golang

# 设置Go模块代理
ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o main .
EXPOSE 8080
CMD ["./main"]