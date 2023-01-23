migration-up:
	migrate -path ./migration/postgres/ -database 'postgres://samandar:saman107@localhost:5432/exam_task?sslmode=disable' up 


migration-down:
	migrate -path ./migration/postgres/ -database 'postgres://samandar:saman107@localhost:5432/exam_task?sslmode=disable' down



run:
	go run cmd/main.go
swag:
	swag init -g api/api.go -o api/docs

