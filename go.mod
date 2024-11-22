//go.mod

module data-rest

go 1.23.2

require github.com/gorilla/mux v1.8.1

require github.com/Peter-Bird/Flash-DB v0.0.1

require github.com/Peter-Bird/models v0.0.1

replace github.com/Peter-Bird/db v0.0.1 => /home/julian/pkg/db

replace github.com/Peter-Bird/models v0.0.1 => /home/julian/pkg/models
