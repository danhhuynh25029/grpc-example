gen-cal:
	protoc --go-grpc_out=. proto/message.proto
run-server:
	go run server.go
run-client:
	go run ./client/client.go