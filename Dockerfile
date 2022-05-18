FROM golang:1.17-alpine
MAINTAINER pumpkin_mak

## 在docker的根目录下创建相应的使用目录
RUN mkdir -p /www/go-logistics
## 设置工作目录
WORKDIR /www/go-logistics
## 把当前（宿主机上）目录下的文件都复制到docker上刚创建的目录下
COPY . /www/go-logistics

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
	GOPROXY="https://goproxy.cn,direct"

COPY go.mod .
COPY go.sum .
RUN go mod download

#暴露端口
EXPOSE 9090

CMD ["go","run","main.go"]
