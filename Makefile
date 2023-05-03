gen:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/*.proto

server:
	go run cmd/server/main.go -port 8080

client:
	go run cmd/client/main.go -address 0.0.0.0:8080

# alternatively this command can be used as well.
# gen:
#     protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:pb

