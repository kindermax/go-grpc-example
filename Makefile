export DOCKER_BUILDKIT = 1
export CURRENT_UID = $(shell echo "`id -u`:`id -g`")

build_grpc_tools:
	docker build -t grpc-tools . -f Dockerfile.grpc-tools

build_base:
	docker build -t go-grpc-example . -f Dockerfile

grpc-gen:
	docker-compose run --rm grpc-gen

server:
	docker-compose up server

client:
	docker-compose run --rm client