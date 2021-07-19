run: 
	go run .
test:
	go test -v ./...
server:
	docker-compose up --build