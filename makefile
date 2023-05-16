IMAGE?=app_land_x
VERSION?=0

JWT_SECRET?=Wia3d3zRH84SuLo5n6WCfR5YNU09qLLZHlBnWeGnFZ
REDIS_PASSWORD?=redis
RABBIT_USERNAME?=rabbit
RABBIT_PASSWORD?=rabbit
DB_USERNAME?=mongo
DB_PASSWORD?=mongo

APP=./app

.PHONY: dev
dev:
	MONGODB_URI=$(MONGODB_URI) \
	JWT_SECRET=$(JWT_SECRET) \
  REDIS_URI=redis://default:$(REDIS_PASSWORD)@localhost:6379 \
  RABBIT_URI=amqp://$(RABBIT_USERNAME):$(RABBIT_PASSWORD)@localhost:5672 \
  MONGODB_URI=mongodb://$(DB_USERNAME):$(DB_PASSWORD)@localhost:27017 \
	go run $(APP)

.PHONY: fmt
fmt:
	go fmt $(APP)
	go mod tidy

.PHONY: test
test:
	go test -v -cover ./ -count=1

.PHONY: build
build:
	docker build -t app_land_x:$(VERSION) .

.PHONY: deploy
deploy:
	make build
	IMAGE=$(IMAGE) \
	VERSION=$(VERSION) \
	JWT_SECRET=$(JWT_SECRET) \
	REDIS_PASSWORD=$(REDIS_PASSWORD) \
	RABBIT_USERNAME=$(RABBIT_USERNAME) \
	RABBIT_PASSWORD=$(RABBIT_PASSWORD) \
	DB_USERNAME=$(DB_USERNAME) \
	DB_PASSWORD=$(DB_PASSWORD) \
	docker stack deploy \
		--compose-file=./docker-stack.yml \
		--prune \
		app_land_x_stack
