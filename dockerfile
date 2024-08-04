FROM registry.cn-hangzhou.aliyuncs.com/jrjr/golang:1.22.5-alpine as builder

ARG APP_DIR=./app

ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /app

COPY ./go.mod /app
COPY ./go.sum /app

RUN go mod download

COPY . /app

RUN go build -o=x ${APP_DIR}


FROM scratch as release

ENV TZ=Asia/Shanghai
ENV GIN_MODE=release

WORKDIR /app

COPY --from=builder /app/x .

EXPOSE 8080

CMD ["/app/x"]