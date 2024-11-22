/*
	Get    - an empty list from the dba web service
	Post   - a workflow to the dba web service
	Get    - the workflow by id
	Post   - a second workflow to the dba web service
	Get    - the list of workflows
	Delete - the workflows
	Get    - the empty list of workflows
*/

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const jsondata = `{
	"id": 1,
	"name": "Test Workflow",
	"steps": [
		{
			"endpoint": "/start",
			"method": "POST",
			"parameters": {
				"initialData": "data1"
			},
			"dependencies": []
		},
		{
			"endpoint": "/process",
			"method": "POST",
			"parameters": {
				"input": "data2",
				"config": {
					"optionA": true,
					"optionB": 5
				}
			},
			"dependencies": ["step1"]
		},
		{
			"endpoint": "/finalize",
			"method": "GET",
			"parameters": null,
			"dependencies": ["step2"]
		}
	]
}`

func main() {
	url := "http://localhost:8083"

	fmt.Printf("in Main - URL is: %s \n", url)

	getWorkflows(url)

	// Post a workflow to the dba web service
	workflow1 := getJson()
	postWorkflow(url, workflow1)

	getWorkflowByID(url, "1")

	// // Post a second workflow to the dba web service
	// workflow2 := map[string]interface{}{
	// 	"id":   "workflow2",
	// 	"name": "Second Workflow",
	// }
	// postWorkflow(url, workflow2)

	// getWorkflows(url)

	// // Delete the workflows
	// deleteWorkflow(url, "workflow1")
	// deleteWorkflow(url, "workflow2")

	// getWorkflows(url)
}

func getWorkflows(url string) {
	resp, err := http.Get(fmt.Sprintf("%s/workflows", url))
	if err != nil {
		fmt.Println("Error getting workflows:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Workflows list:", string(body))
}

func getWorkflowByID(url, id string) {
	resp, err := http.Get(fmt.Sprintf("%s/workflows/%s", url, id))
	if err != nil {
		fmt.Println("Error getting workflow by ID:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("Workflow %s: %s\n", id, string(body))
}

func postWorkflow(url string, workflow map[string]interface{}) {
	jsonData, err := json.Marshal(workflow)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	fmt.Printf("%v\n\n", jsonData)

	resp, err := http.Post(fmt.Sprintf("%s/workflows", url), "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error posting workflow:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Posted workflow:", string(body), "End.")
}

func deleteWorkflow(url, id string) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/workflows/%s", url, id), nil)
	if err != nil {
		fmt.Println("Error creating DELETE request:", err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error deleting workflow:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("Deleted workflow %s: %s\n", id, string(body))
}

func getJson() map[string]interface{} {
	var data map[string]interface{}

	err := json.Unmarshal([]byte(jsondata), &data)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil
	}

	return data
}
