CURRENT_DIR=$(shell pwd)

APP=$(shell basename ${CURRENT_DIR})
APP_CMD_DIR=${CURRENT_DIR}/cmd

TAG=latest
ENV_TAG=latest

migration-up:
	migrate -path ./migrations/postgres -database 'postgres://khumoyun:admin@111@localhost:5432/shop_product?sslmode=disable' up

migration-down:
	migrate -path ./migrations/postgres -database 'postgres://khumoyun:admin@111@localhost:5432/shop_product?sslmode=disable' down

build:
	CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go

swag-init:
	swag init -g api/api.go -o api/docs

run:
	go run cmd/main.go