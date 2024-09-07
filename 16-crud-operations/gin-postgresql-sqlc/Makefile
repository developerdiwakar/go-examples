# Migration Related Commands

migrate_up:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/contact_db?sslmode=disable" -verbose up

migrate_down:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/contact_db?sslmode=disable" -verbose down