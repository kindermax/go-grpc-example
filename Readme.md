# Go grpc example

### Run

Via docker-compose and make
 
```bash
# build docker image
make build_base

make server
make client
```

### Build grpc

```bash
# build docker image
make build_grpc_tools
# build and gen grpc stubs from .proto files
make grpc-gen
```