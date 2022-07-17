init:
	go mod download
.PHONY: init

build:
	go build -o build/clingy
.PHONY: build

start:
	go run main.go
.PHONY: start

clean:
	rm build/clingy
.PHONY: clean

lint:
	docker run \
		--rm \
		-v $(shell pwd):/app \
		-w /app \
		golangci/golangci-lint:v1.46 \
		golangci-lint run
.PHONY: lint

test: lint
	go test ./...
.PHONY: test

pretty:
	go fmt ./...
.PHONY: pretty

release:
	@echo "don't forget to create and push a git tag! e.g."
	@echo "  git tag -a v0.1.0 -m 'First release'"
	@echo "  git push origin v0.1.0"
	@sleep 3
	goreleaser release
.PHONY: release

build-docker:
	docker build -t clingy .

run-docker: build-docker
	# See readme on usage
	xhost local:root
	docker run \
		-e DISPLAY=${DISPLAY} \
		-v /tmp/.X11-unix:/tmp/.X11-unix \
		-v ${PWD}/.clingy.yaml:/home/clingy/.clingy.yaml \
		clingy
.PHONY: build-docker run-docker