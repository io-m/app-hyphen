## run: runs the main function
run:
	@echo "Running Main function..."
	go run ./cmd/main.go

arango-up:
	docker-compose -f docker-compose.yml up

arango-down:
	docker-compose -f docker-compose.yml down
