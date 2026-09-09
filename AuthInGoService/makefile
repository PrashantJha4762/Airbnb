MIGRATIONS_FOLDER = "db/migrations"
DB_URL = "root:areyoumad4762@tcp(localhost:3306)/AUTH_DEV"
migration-create:
	goose -dir $(MIGRATIONS_FOLDER) create $(name) sql

migration-up:
	goose -dir $(MIGRATIONS_FOLDER) mysql $(DB_URL) up

migration-down:
	goose -dir $(MIGRATIONS_FOLDER) mysql $(DB_URL) down

