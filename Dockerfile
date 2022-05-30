FROM golang:1.17-alpine
MAINTAINER pumpkin_mak

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
	GOPROXY="https://goproxy.cn,direct"
    
## 设置工作目录
WORKDIR /src/go-logistics
## 把当前（宿主机上）目录下的文件都复制到docker上刚创建的目录下
COPY . .
RUN go get -d -v ./...
RUN go install -v ./... 

CMD ["go-logistics"]
