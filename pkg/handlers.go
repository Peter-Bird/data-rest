// wf-dba: pkg/handlers.go
package pkg

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/Peter-Bird/models"
	"github.com/gorilla/mux"
)

type Handler struct {
	Service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{Service: service}
}

func (h *Handler) PostHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log.Println("PostHandler invoked")

	// Ensure the request method is POST
	if r.Method != http.MethodPost {
		log.Printf("Invalid request method: %s\n", r.Method)
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Read and parse the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Failed to read request body: %v\n", err)
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	log.Println("Request body read successfully")

	var workflow models.Workflow
	if err := json.Unmarshal(body, &workflow); err != nil {
		log.Printf("Invalid JSON format: %v\n", err)
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}
	log.Printf("Request body unmarshaled into workflow: %s\n", workflow.Id)

	// Use the service to post the workflow
	if err := h.Service.Post(ctx, workflow); err != nil {
		log.Printf("Failed to post workflow %s: %v\n", workflow.Id, err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("Workflow %s posted successfully\n", workflow.Id)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Workflow posted successfully"))
}

func (h *Handler) GetHandler(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context()
	vars := mux.Vars(r)
	workflowID := vars["id"]

	log.Println("GetHandler invoked")

	if workflowID == "" {
		writeJSONError(w, http.StatusBadRequest, "Missing workflow ID")
		return
	}

	workflow, err := h.Service.Get(workflowID)
	if err != nil {
		writeJSONError(w, http.StatusNotFound, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(workflow)
}

func (h *Handler) ListHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("ListHandler invoked")

	workflows, err := h.Service.List()
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, "Failed to retrieve workflows")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(workflows)
}

func (h *Handler) TruncateHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("DeleteHandler invoked")
	h.Service.Truncate()

	// 204 No Content, indicating success without response body
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("DeleteIdHandler invoked")

	// Extract the workflow ID from the URL path variables
	vars := mux.Vars(r)
	workflowID := vars["id"]

	if workflowID == "" {
		writeJSONError(w, http.StatusBadRequest, "Missing workflow ID")
		return
	}

	// Attempt to delete the specific workflow using the service
	err := h.Service.Delete(workflowID)
	if err != nil {
		log.Printf("Failed to delete workflow %s: %v\n", workflowID, err)
		writeJSONError(w, http.StatusInternalServerError, "Failed to delete workflow")
		return
	}

	log.Printf("Workflow %s deleted successfully\n", workflowID)
	// 204 No Content, indicating success without response body
	w.WriteHeader(http.StatusNoContent)
}

/*
	Helper Functions
*/

type ErrorResponse struct {
	Error string `json:"error"`
}

func writeJSONError(w http.ResponseWriter, statusCode int, errMsg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{Error: errMsg})
}
