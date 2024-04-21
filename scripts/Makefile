generate_grpc_code:
	protoc --go_out=api/core/grpc --go_opt=paths=source_relative --go-grpc_out=api/core/grpc --go-grpc_opt=paths=source_relative --proto_path=C:\dev\freelance\customerService\api\ customerService.proto

generate_docker_image:
	docker run -e PORT=3001 -p 3001:3001 grpc-costumer-service:latest