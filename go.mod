module bootdev

require internal/database v1.0.0

require golang.org/x/crypto v0.22.0 // indirect

replace internal/database => ./internal/database

go 1.22.0
