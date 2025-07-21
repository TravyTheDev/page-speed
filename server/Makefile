run:
	@go run . run

up: 
	@goose -dir=db/migrations sqlite db/app.db up

down: 
	@goose -dir=db/migrations sqlite db/app.db down

seed: 
	@go run . seed