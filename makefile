IMAGE?=app_land_x
VERSION?=0

JWT_SECRET?=Wia3d3zRH84SuLo5n6WCfR5YNU09qLLZHlBnWeGnFZ
DB_USERNAME?=mongo
DB_PASSWORD?=mongo

APP=./app

.PHONY: dev
dev:
	JWT_SECRET=$(JWT_SECRET) \
  MONGODB_URI=mongodb://localhost:27017 \
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
	docker build -t $(IMAGE):$(VERSION) .

.PHONY: deploy
deploy:
	make build
	IMAGE=$(IMAGE) \
	VERSION=$(VERSION) \
	JWT_SECRET=$(JWT_SECRET) \
	DB_USERNAME=$(DB_USERNAME) \
	DB_PASSWORD=$(DB_PASSWORD) \
	docker stack deploy \
		--compose-file=./docker-stack.yml \
		--prune \
		stack_$(IMAGE)
