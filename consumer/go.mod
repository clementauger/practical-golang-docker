module github.com/clementauger/practical-golang-docker/producer

go 1.12

require (
	github.com/clementauger/practical-golang-docker/model v0.0.0
	github.com/gorilla/mux v1.7.3
)

replace github.com/clementauger/practical-golang-docker/model v0.0.0 => ../model
