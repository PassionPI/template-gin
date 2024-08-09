# 时区, 默认为上海
TIMEZONE?=Asia/Shanghai
# 项目目录, 默认为 ./app
APP?=./app
# 镜像名称, 默认为 app-ink
IMAGE?=app-ink

# 版本号, 需要设置, 默认为 0. 
VERSION?=0
# 下面为所以来环境的变量, 可以根据需要修改
JWT_SECRET?=JWT_SECRET
REDIS_PASSWORD?=redis
PG_USERNAME?=postgres
PG_PASSWORD?=postgres


.PHONY: dev
dev:
	IMAGE=$(IMAGE) \
	JWT_SECRET=$(JWT_SECRET) \
	REDIS_URI=redis://localhost:6379 \
  POSTGRES_URI=postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable \
	go run $(APP)

.PHONY: fmt
fmt:
	go fmt $(APP)
	go mod tidy

.PHONY: test
test:
	go test -v -cover -count=1 ./...

.PHONY: lint
lint:
	make fmt
	make test

.PHONY: build
build:
	docker build --build-arg APP_DIR=$(APP) -t $(IMAGE):$(VERSION) .
	
.PHONY: deploy
deploy:
	make build
	IMAGE=$(IMAGE) \
	VERSION=$(VERSION) \
	TIMEZONE=$(TIMEZONE) \
	JWT_SECRET=$(JWT_SECRET) \
	REDIS_PASSWORD=$(REDIS_PASSWORD) \
	PG_USERNAME=$(PG_USERNAME) \
	PG_PASSWORD=$(PG_PASSWORD) \
	docker stack deploy \
		--compose-file=./docker-stack.yml \
		--prune \
		$(IMAGE)
