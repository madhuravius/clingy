init:
	go mod download
	go generate ./...
.PHONY: init

build:
	echo "Dependencies used: "
	go list -m all
	echo "Building"
	go build -v -o build/clingy
.PHONY: build

docs:
	go run *.go run \
		-i ./cmd/test_data/01_basic_flow_will_pass.yaml \
		-o ./docs/example-outputs/carousel \
		-r carousel \
		-u False
	go run *.go run \
		-i ./cmd/test_data/01_basic_flow_will_pass.yaml \
		-o ./docs/example-outputs/html-simple \
		-u False
	go run *.go run \
		-i ./cmd/test_data/01_basic_flow_will_pass.yaml \
		-o ./docs/example-outputs/images-only \
		-r images-only \
		-u False
.PHONY: docs

start:
	go run main.go
.PHONY: start

clean:
	rm build/clingy || true
	go run *.go clean
	go mod tidy
.PHONY: clean

lint:
	docker run \
		--rm \
		-v $(shell pwd):/app \
		-w /app \
		golangci/golangci-lint:v1.46 \
		golangci-lint run
.PHONY: lint

test:
	gotestsum --format testname -- -coverprofile=cover.out ./...
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

start-docs:
	docker run \
		--rm \
		-it \
		-p 8000:8000 \
		-v ${PWD}:/docs squidfunk/mkdocs-material
.PHONY: start-docs