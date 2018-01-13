grpc:
	protoc --proto_path=. --go_out=plugins=grpc:. spike.proto