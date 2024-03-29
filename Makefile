.PHONY: build migrationup generate clean

build:
	go build -o rssagg && ./rssagg

migrationup:
	cd sql/schema && goose postgres postgres://myuser:sample@localhost:5432/rssagg up

migrationdown:
	cd sql/schema && goose postgres postgres://myuser:sample@localhost:5432/rssagg down

generate:
	sqlc generate

clean:
	rm -f rssagg
