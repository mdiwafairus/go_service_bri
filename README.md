Database menggunakan postgreSql (GORM).
Framework menggunakan Fiber .

Logging dibuat secara global di level database query (GORM) dan aplikasi secara global dengan file zerolog_gorm.go & zerolog.go
- log error kedalam file error.log
- log info & warn kedalam file app.log

handling database up/down di level middleware, route untuk cek database "healthcare"

rate limit di level middleware 