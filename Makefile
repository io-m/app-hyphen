## run: runs the main function
run:
	@echo "Running Main function..."
	go run ./cmd/main.go

db-up:
	docker-compose -f ./infrastructure/docker-compose-template.yml up

db-down:
	docker-compose -f ./infrastructure/docker-compose-template.yml down
