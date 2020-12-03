module github.com/isurusiri/services/foo

go 1.12

require (
	github.com/gorilla/mux v1.8.0
	github.com/isurusiri/services/shared v0.0.0-00010101000000-000000000000
)

replace github.com/isurusiri/services/shared => ../shared
