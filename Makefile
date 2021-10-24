APP?=litclock-service

default:
	echo ${APP}

all: test build

.PHONEY: build
## build: build the application
build: clean
	@echo "Building..."
	@go build -o ${APP} main.go


.PHONY: run
## run: runs go run main.go
run:
	go run main.go


.PHONY: clean
## clean: cleans the binary
clean:
	@echo "Cleaning"
	@rm -rf ${APP}

.PHONY: test
## test: runs go test with default values
test:
	go test -v -count=1 -race ./...

docker_build:
	docker build . -t ${APP}

docker_run:
	docker run -it --rm -p 8080:8080 ${APP}

.PHONY: help
## help: prints this help message
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'
