
sqlc.gen:
	sqlc generate

sqlc.create:
	migrate create -ext sql -dir migrations -seq create_authors_table

go.install:
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest