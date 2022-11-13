gen-cal:
	protoc proto/message.proto --go-grpc_out=.
	protoc proto/message.proto --go_out=.
run-server:
	go run server.go
run-client:
	go run ./client/client.go