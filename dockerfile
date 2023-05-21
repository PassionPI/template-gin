# 构建前端资源
FROM node:alpine as builder-frontend

WORKDIR /app

COPY frontend/package.json .
COPY frontend/pnpm-lock.yaml .

RUN npm config set registry https://registry.npmmirror.com
RUN npm install -g pnpm
RUN pnpm install

COPY frontend .

RUN pnpm run build


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

COPY --from=builder-frontend /app/dist ./frontend
COPY --from=builder-backend /app/x .

EXPOSE 8080

CMD ["/app/x"]