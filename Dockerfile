FROM golang:1.17-alpine
MAINTAINER pumpkin_mak

RUN adduser -u 10001 -D app-runner

ADD go.mod .
ADD go.sum .
RUN go mod download
WORKDIR /go/src/go-logistics
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    GIN_MODE=release

COPY . .
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -a -o go-logistics .

COPY . .
COPY --from=build /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN go build .
EXPOSE 9090

USER app-runner
CMD ["./go-logistics"]