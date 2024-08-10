.PHONY: movie metadata rating

movie-service:
	go run movie/cmd/main.go

metadata-service:
	go run metadata/cmd/main.go

rating-service:
	go run rating/cmd/main.go


get-rating:
	grpcurl -plaintext -d '{"record_id": "1", "record_type": "movie"}' localhost:8082 RatingService/GetAggregatedRating

put-rating:
	grpcurl -plaintext -d '{"record_id":"1", "record_type": "movie", "user_id": "alex", "rating_value": 5}' localhost:8082 RatingService/PutRating

proto:
	protoc -I=api --go_out=. --go-grpc_out=. movie.proto 

db-up:
	docker run --name movieexample_db -e MYSQL_ROOT_PASSWORD=password -e MYSQL_DATABASE=movieexample -p 3305:3306 -d mysql:8.4

db-down:
	docker stop movieexample_db
	docker rm movieexample_db

db-migrate-up`:
	docker exec -i movieexample_db mysql movieexample -h localhost -P 3306 --protocol=tcp -uroot -p"password" < schema/schema.sql

db-migrate-down:
	docker exec -i movieexample_db mysql movieexample -h localhost -P 3306 --protocol=tcp -u root -p"password" < schema/schema_drop.sql
	
db-show:
	docker exec -i movieexample_db mysql movieexample -h localhost -P 3306 --protocol=tcp -uroot -p"password" -e "SHOW tables"
