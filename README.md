---- Migration -----
Cara instal migration :
go install -tags 'mysql' github.com/golang-migrate/migrate@latest

Cara membuat migration :
migrate create -ext sql -dir database/migrations create_table_categories

Cara menjalankan migration :
migrate -database "mysql://root@tcp(localhost:3306)/golang_restful_api" -path database/migrations up

Cara remove dirty :
migrate -path database/migrations -database "mysql://root@tcp(localhost:3306)/golang_restful_api" force 20240320160949

Cara migrate ke versi tertentu (versi 1 misalnya) :
migrate -database "mysql://root@tcp(localhost:3306)/golang_restful_api" -path database/migrations up 1

Cara run api :
go run main.go wire_gen.go