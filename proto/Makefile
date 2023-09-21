server:
	protoc *.proto --go_out=../server --go-grpc_out=../server
client:
	protoc *.proto --go_out=../client --go-grpc_out=../client
all: server client
