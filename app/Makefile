NAME    := http-server
VERSION := 0.0.1

GIT_COMMIT_ID := $(shell git rev-parse --short HEAD)

BUILD_DIR   := ./build
BIN_DIR     := ${BUILD_DIR}/_output/bin
BUILD_PATH  := ./cmd/${NAME}.go

REGISTRY := localhost:5000

.PHONY: build

version:
	@echo ${VERSION}-${GIT_COMMIT_ID}

build:
	GOOS=linux go build -o ${BIN_DIR}/${NAME} ${BUILD_PATH}

image:
	docker build -f ${BUILD_DIR}/Dockerfile -t ${NAME}:${VERSION} .

push:
	docker tag ${NAME}:${VERSION} ${REGISTRY}/${NAME}:${VERSION}-${GIT_COMMIT_ID}
	docker push ${REGISTRY}/${NAME}:${VERSION}-${GIT_COMMIT_ID}

lint:
	golangci-lint run ./...

test:
	go test -cover ./...  -coverprofile=coverage.out

coverage:
	go tool cover -html=coverage.out -o coverage.html

clean:
	go clean
	rm -f ${BIN_DIR}/${NAME}
