.PHONY: movie metadata rating

movie-service:
	go run movie/cmd/main.go

metadata-service:
	go run metadata/cmd/main.go

rating-service:
	go run rating/cmd/main.go

proto:
	protoc -I=api --go_out=. --go-grpc_out=. movie.proto 