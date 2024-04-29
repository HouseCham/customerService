generate_grpc_code_windows:
	protoc --go_out=api/core/grpc --go_opt=paths=source_relative --go-grpc_out=api/core/grpc --go-grpc_opt=paths=source_relative --proto_path=C:\dev\freelance\customerService\api\ customerService.proto

generate_grpc_code_linux:
	protoc --go_out=api/core/grpc --go_opt=paths=source_relative --go-grpc_out=api/core/grpc --go-grpc_opt=paths=source_relative --proto_path=/mnt/c/dev/freelance/customerService/api/ customerService.proto

first_build_docker_image:
	docker build -t costumer_service:[version] .

run_docker_image:
	docker run -e PORT=3001 -p 3001:3001 grpc-costumer-service:latest