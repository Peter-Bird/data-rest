// wf-dba: pkg/server.go
package pkg

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewServer(port string, router *mux.Router) *http.Server {

	server := &http.Server{Addr: ":" + port, Handler: router}

	return server
}
