grpc:
	protoc --go_out=plugins=grpc:. spike/spike.proto

build:
	mkdir -p ../out
	go build -o ../out/grpc-server
