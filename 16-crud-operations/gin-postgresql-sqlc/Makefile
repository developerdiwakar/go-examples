# Migration Related Commands


migration:
	migrate create -ext sql -dir db/migration -seq $(name)

migrate.up:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/contact_db?sslmode=disable" -verbose up

migrate.down:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/contact_db?sslmode=disable" -verbose down


source.link:
	echo "https://dev.to/geoff89/deploying-a-golang-restful-api-with-gin-sqlc-and-postgresql-1lbl"