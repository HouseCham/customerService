generate_grpc_code_windows:
	protoc --go_out=api/core/grpc --go_opt=paths=source_relative --go-grpc_out=api/core/grpc --go-grpc_opt=paths=source_relative --proto_path=C:\dev\freelance\customerService\api\ customerService.proto

create_docker_network:
	sudo docker network create grpc_network

first_build_docker_image:
	docker build -t customer_service:[version] .

run_docker_image:
	sudo docker run -d \
	--name customer-microservice \
	-p 3001:3001 \
	--restart unless-stopped \
	customer_service:1

# --network grpc_network \

test_service:
	sudo docker run -p 3001:3001 -it --rm customer_service:1

#access to the postgres server CLI
showLogs:
	sudo docker logs customer-microservice