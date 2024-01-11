.PHONY: migrate-up migrate-down generate run

# Učitajte vrijednost DB_CONNECTION_STRING iz .env datoteke
DB_URL := $(shell grep '^DB_URL' .env | cut -d '=' -f2)

migrate-up:
	goose -dir sql/schema postgres "$(DB_URL)" up

migrate-down:
	goose -dir sql/schema postgres "$(DB_URL)" down

generate:
	sqlc generate

run:
	go build && ./rss-aggregator