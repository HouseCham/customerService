build-image:
	docker buildx build --platform linux/amd64,linux/arm64,linux/arm/v7 -t chamito/customer-service:1.0.0 --push .