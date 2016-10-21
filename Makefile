NAME=ragnarb/lto
VERSION=0.1

default: build

.PHONY: build
build: test
	go build

.PHONY: test
test:
	go test

.PHONY: container
container:
	docker build --tag ${NAME} ${CURDIR}
	docker tag ${NAME}:latest ${NAME}:${VERSION}

.PHONY: clean
clean:
	go clean
	-docker rmi --force ${NAME}:latest
	-docker rmi --force ${NAME}:${VERSION}

