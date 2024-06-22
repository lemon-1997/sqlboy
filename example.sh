go build ./cmd/sqlboy/main.go
./main ./examples/gorm/stmt.go -mode gorm
./main ./examples/sqlx/stmt.go -mode sqlx
