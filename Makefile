generate_grpc_code:
	protoc \
		--go_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_out=. \
		--go-grpc_opt=paths=source_relative \
		./chat/chat.proto

start_server:
	go run server/server.go

start_client:
	go run client/client.go

list_services:
	grpcurl -plaintext localhost:9000 list
	
request_server:
	grpcurl -plaintext -d '{"message": "hello from client"}' \
    localhost:9000 chat.ChatService/SendMessage
