go build ./cmd/sqlboy/main.go
./main.exe ./examples/gorm/stmt.go -mode gorm
./main.exe ./examples/sqlx/stmt.go -mode sqlx
