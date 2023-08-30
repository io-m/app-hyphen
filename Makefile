## run: runs the main function
run:
	@echo "Running Main function..."
	go run ./cmd/main.go

db-up:
	docker-compose -f docker-compose.yml up

db-down:
	docker-compose -f docker-compose.yml down
