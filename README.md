Go Version : go1.24.2
FIber Version : v2.x

Database menggunakan postgreSql (GORM).

Logging dibuat secara global di level database query (GORM) dan aplikasi secara global dengan file zerolog_gorm.go & zerolog.go
- log error kedalam file error.log
- log info & warn kedalam file app.log

handling database up/down di level middleware, route untuk cek database "healthcare"

rate limit di level middleware 