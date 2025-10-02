migrate:
	docker compose --env-file .postgres.env run --rm migrate
migration:
	docker run --rm -v ./db/migrations/:/migrations/ migrate/migrate:v4.18.3 create -ext sql -dir /migrations -seq $(name)
