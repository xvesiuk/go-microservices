.PHONY: app
app:
	go run cmd/app/main.go

.PHONY: swag
swag:
	swag init -d ./cmd/app,./internal/rest -o ./api/docs

.PHONY: db
db:
	docker run -d \
		--name postgres-microservices \
		-e POSTGRES_USER=dev \
		-e POSTGRES_PASSWORD=dev \
		-p 5433:5432 \
		postgres

# (L)ocal migrate 
.PHONY: lmigrate
lmigrate:
	goose -dir migrations postgres "postgresql://dev:dev@127.0.0.1:5433/postgres?sslmode=disable" up