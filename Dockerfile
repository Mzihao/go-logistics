FROM golang:1.17-alpine

WORKDIR /src/logistics

RUN go env -w GO111MODULE=on

RUN go env -w GOPROXY=https://goproxy.cn,https://goproxy.io,direct

ADD go.mod .

RUN go mod download

COPY . .

CMD "./run.sh"