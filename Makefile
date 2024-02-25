.PHONY: build run clean

migrationup:
	cd sql/schema && goose postgres postgres://postgres:kiranraj27@localhost:5432/rssagg up

migrationdown:
	cd sql/schema && goose postgres postgres://postgres:kiranraj27@localhost:5432/rssagg down

generate:
	sqlc generate

build:
	go build -o rssagg && ./rssagg

clean:
	rm -f rssagg
