IMAGE?=app_land_x
VERSION?=0

JWT_SECRET?=JWT_SECRET

REDIS_PASSWORD?=redis

DB_USERNAME?=mongo
DB_PASSWORD?=mongo

RABBIT_USERNAME?=rabbit
RABBIT_PASSWORD?=rabbit

APP=./app

.PHONY: dev
dev:
	JWT_SECRET=$(JWT_SECRET) \
	REDIS_URI=redis://localhost:6379 \
  MONGODB_URI=mongodb://localhost:27017 \
	go run $(APP)

.PHONY: fmt
fmt:
	go fmt $(APP)
	go mod tidy

.PHONY: test
test:
	go test -v -cover -count=1 ./...

.PHONY: build
build:
	docker build --build-arg APP_DIR=./app -t $(IMAGE):$(VERSION) .

.PHONY: deploy
deploy:
	make build
	IMAGE=$(IMAGE) \
	VERSION=$(VERSION) \
	JWT_SECRET=$(JWT_SECRET) \
	REDIS_PASSWORD=$(REDIS_PASSWORD) \
	DB_USERNAME=$(DB_USERNAME) \
	DB_PASSWORD=$(DB_PASSWORD) \
	docker stack deploy \
		--compose-file=./docker-stack.yml \
		--prune \
		stack_$(IMAGE)
