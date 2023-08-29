## run: runs the main function
run:
	@echo "Running Main function..."
	go run ./cmd/main.go

arango-up:
	docker-compose -f infrastructure/arango/docker-compose.yml up

arango-down:
	docker-compose -f infrastructure/arango/docker-compose.yml down
