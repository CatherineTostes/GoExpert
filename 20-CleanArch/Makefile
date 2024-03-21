migrateUp:
	migrate -path=migrations -database "mysql://root:root@tcp(localhost:3306)/orders" -verbose up

migrateDown:
	migrate -path=migrations -database "mysql://root:root@tcp(localhost:3306)/orders" -verbose down

.PHONY: migrateUp migrateDown