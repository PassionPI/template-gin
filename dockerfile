FROM golang:alpine as builder-backend

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

RUN go build -o=x ./app


FROM scratch as release

ENV GIN_MODE=release

WORKDIR /app

COPY --from=builder-backend /app/x .

EXPOSE 8080

CMD ["/app/x"]