run: 
	go run .
test:
	go test -v ./...
server:
	docker-compose up --build
clean: 
	docker network remove neti_default