gen:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/*.proto

# alternatively this command can be used as well.
# gen:
#     protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:pb

