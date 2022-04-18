run-server:
	go run cmd/server/main.go

run-proxy:
	go run cmd/proxy/proxy.go

proto-generate:
	protoc --proto_path=internal/proto internal/proto/*.proto --go_out=plugins=grpc:internal/pb --grpc-gateway_out=:internal/pb --swagger_out=:internal/swagger

proto-clean:
	rm internal/pb/*.go

swagger-clean:
	rm internal/swagger/*.swagger.json

run-paquetes:
	go mod tidy