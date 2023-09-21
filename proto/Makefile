server:
	protoc *.proto --go_out=../server --go-grpc_out=../server
client:
	protoc *.proto --go_out=../client --go-grpc_out=../client
evan:
	evans --proto=./calculator.proto --port=5050
all: server client
