// wf-dba: pkg/routes.go
package pkg

import (
	"github.com/gorilla/mux"
)

func NewRouter(handler *Handler) *mux.Router {

	//mux := http.NewServeMux()
	router := mux.NewRouter()
	router.HandleFunc("/workflows", handler.PostHandler).Methods("POST")
	router.HandleFunc("/workflows/{id}", handler.GetHandler).Methods("GET")
	router.HandleFunc("/workflows", handler.ListHandler).Methods("GET")
	router.HandleFunc("/workflows", handler.TruncateHandler).Methods("DELETE")
	router.HandleFunc("/workflows/{id}", handler.DeleteHandler).Methods("DELETE")

	return router
}
