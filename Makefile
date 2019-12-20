export DOCKER_BUILDKIT = 1
export CURRENT_UID = $(shell echo "`id -u`:`id -g`")

build_grpc_tools:
	docker build -t grpc-tools . -f Dockerfile.grpc-tools

grpc-gen:
	docker-compose run --rm grpc-gen