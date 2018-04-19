gen_proto:
	protoc -I./proto --go_out=plugins=grpc:./distance_service/proto ./proto/*.proto && \
	grpc_tools_ruby_protoc -I ./proto --ruby_out=./price_service/proto --grpc_out=./price_service/proto ./proto/*.proto

build_containers:
	cd price_service && docker build -t kiopro/price_service:latest . && \
	cd .. && \
	cd distance_service && \
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o distance_service . && \
	docker build -t kiopro/distance-service:latest .
