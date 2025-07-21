run:
	@go run . run

up: 
	@goose -dir=db/migrations sqlite db/app.db up

down: 
	@goose -dir=db/migrations sqlite db/app.db down

down-all:
	@goose -dir=db/migrations sqlite db/app.db down-to 0

seed: 
	@go run . seed