// wf-dba: pkg/services.go
package pkg

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/Peter-Bird/db"
	"github.com/Peter-Bird/models"
)

type Service struct {
	Database db.Repository
}

// Note :- Currently using a specific database structure (models.workflow)
func NewService(database *db.FlashDB) *Service {

	service := &Service{Database: database}

	return service
}

func mapToStruct(data map[string]interface{}, target interface{}) error {
	// Convert map to JSON bytes
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return errors.New("failed to marshal map to JSON: " + err.Error())
	}

	// Unmarshal JSON bytes into the target struct
	err = json.Unmarshal(jsonBytes, target)
	if err != nil {
		return errors.New("failed to unmarshal JSON to struct: " + err.Error())
	}

	return nil
}

func (s *Service) Get(id string) (models.Workflow, error) {
	if id == "" {
		return models.Workflow{}, errors.New("workflow ID cannot be empty")
	}

	data, err := s.Database.Get(id)
	if err != nil {
		return models.Workflow{}, err // Return the error if the record is not found or malformed
	}

	var workflow models.Workflow
	err = mapToStruct(data, &workflow)
	if err != nil {
		return models.Workflow{}, err // Handle conversion errors
	}

	return workflow, nil
}

func (s *Service) List() ([]models.Workflow, error) {

	dataList, err := s.Database.List()
	if err != nil {
		return nil, err
	}

	var workflows []models.Workflow

	for _, data := range dataList {
		var workflow models.Workflow
		err := mapToStruct(data, &workflow)
		if err != nil {
			return nil, errors.New("failed to convert database record to workflow struct: " + err.Error())
		}
		workflows = append(workflows, workflow)
	}

	return workflows, nil
}

func structToMap(data interface{}) (map[string]interface{}, error) {
	// Marshal the struct to JSON
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON into a map
	var result map[string]interface{}
	err = json.Unmarshal(jsonBytes, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Service) Post(ctx context.Context, workflow models.Workflow) error {
	// Validate the workflow
	if err := validateWorkflow(workflow); err != nil {
		return err
	}

	// Convert workflow to map[string]interface{}
	data, err := structToMap(workflow)
	if err != nil {
		return errors.New("failed to convert workflow to map: " + err.Error())
	}

	// Save to the database
	return s.Database.Save(workflow.Id, data)
}

func validateWorkflow(workflow models.Workflow) error {
	if workflow.Id == "" {
		return errors.New("workflow id is required")
	}
	if workflow.Name == "" {
		return errors.New("workflow name is required")
	}

	return nil
}
func (s *Service) Delete(id string) error {
	log.Println("In Delete Service")

	if id == "" {
		return errors.New("ID cannot be empty")
	}

	err := s.Database.Delete(id) // Pass the context to the database layer
	if err != nil {
		log.Printf("Error deleting ID %s: %v\n", id, err)
		return err
	}

	log.Printf("Successfully deleted ID %s\n", id)
	return nil
}

func (s *Service) Truncate() error {
	log.Println("In Truncate Service")

	err := s.Database.Truncate() // Pass the context to the database layer
	if err != nil {
		log.Printf("Error truncating database: %v\n", err)
		return err
	}

	log.Printf("Successfully deleted database.\n")
	return nil
}
