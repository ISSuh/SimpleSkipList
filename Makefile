# ====================================================================================
# Targets
# ====================================================================================

all:
	go build ./cmd/simpledb
	go build ./cmd/simpledb-cluster

test:
	go test -v ./...

coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out

fmt:
	go fmt ./...

generate:
	protoc -I=internal/transport/rpc/proto \
		--go_out=internal/transport/rpc/proto \
		--go_opt=paths=source_relative \
		--go-grpc_out=internal/transport/rpc/proto \
		--go-grpc_opt=paths=source_relative \
		internal/transport/rpc/proto/storage.proto

clean:
	rm -rf simpledb
	rm -rf simpledb-cluster
