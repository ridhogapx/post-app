postgres:
	docker run --name=postgres_container -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres

createdb:
	docker exec -it postgres_container createdb --username=root --owner=root post_app

dropdb:
	docker exec -it postgres_container dropdb post_app

migrateup:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/post_app?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/post_app?sslmode=disable" -verbose down