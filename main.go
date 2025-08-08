package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/compute/metadata"
	"github.com/gorilla/mux"
)

// ToolSchema defines the structure of an MCP tool
type ToolSchema struct {
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Parameters  map[string]interface{} `json:"parameters"`
	Returns     map[string]interface{} `json:"returns"`
}

// MCPRequest represents an incoming MCP request
type MCPRequest struct {
	ToolName string                 `json:"tool_name"`
	Params   map[string]interface{} `json:"params"`
}

// MCPResponse represents the response to an MCP client
type MCPResponse struct {
	Result interface{} `json:"result"`
	Error  string      `json:"error,omitempty"`
}

// Config holds server configuration
type Config struct {
	MicroserviceURL string
	Port            string
}

// tools defines the MCP tools for all microservice endpoints
var tools = []ToolSchema{
	{
		Name:        "welcome_message",
		Description: "Get the welcome message from the product service",
		Parameters:  map[string]interface{}{"type": "object", "properties": map[string]interface{}{}},
		Returns: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"message": map[string]string{"type": "string"},
			},
		},
	},
	{
		Name:        "health_check",
		Description: "Check the health of the product service",
		Parameters:  map[string]interface{}{"type": "object", "properties": map[string]interface{}{}},
		Returns: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"status": map[string]string{"type": "string"},
			},
		},
	},
	{
		Name:        "create_product",
		Description: "Create a new product in the store",
		Parameters: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"name":  map[string]string{"type": "string"},
				"price": map[string]string{"type": "number"},
			},
			"required": []string{"name", "price"},
		},
		Returns: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"id":    map[string]string{"type": "string"},
				"name":  map[string]string{"type": "string"},
				"price": map[string]string{"type": "number"},
			},
		},
	},
	{
		Name:        "get_product",
		Description: "Retrieve a product by ID",
		Parameters: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"id": map[string]string{"type": "string"},
			},
			"required": []string{"id"},
		},
		Returns: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"id":    map[string]string{"type": "string"},
				"name":  map[string]string{"type": "string"},
				"price": map[string]string{"type": "number"},
			},
		},
	},
	{
		Name:        "update_product",
		Description: "Update an existing product by ID",
		Parameters: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"id":    map[string]string{"type": "string"},
				"name":  map[string]string{"type": "string"},
				"price": map[string]string{"type": "number"},
			},
			"required": []string{"id"},
		},
		Returns: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"id":    map[string]string{"type": "string"},
				"name":  map[string]string{"type": "string"},
				"price": map[string]string{"type": "number"},
			},
		},
	},
	{
		Name:        "delete_product",
		Description: "Delete a product by ID",
		Parameters: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"id": map[string]string{"type": "string"},
			},
			"required": []string{"id"},
		},
		Returns: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"message": map[string]string{"type": "string"},
			},
		},
	},
	{
		Name:        "list_products",
		Description: "List all products in the store",
		Parameters:  map[string]interface{}{"type": "object", "properties": map[string]interface{}{}},
		Returns: map[string]interface{}{
			"type": "array",
			"items": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"id":    map[string]string{"type": "string"},
					"name":  map[string]string{"type": "string"},
					"price": map[string]string{"type": "number"},
				},
			},
		},
	},
}

// getToolsHandler returns the list of available tools
func getToolsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tools)
}

// mcpHandler processes MCP requests
func mcpHandler(config Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req MCPRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Validate tool name
		var tool ToolSchema
		for _, t := range tools {
			if t.Name == req.ToolName {
				tool = t
				break
			}
		}
		if tool.Name == "" {
			http.Error(w, "Tool not found", http.StatusNotFound)
			return
		}

		// Map tool to microservice endpoint
		var method, url string
		var body []byte
		switch req.ToolName {
		case "welcome_message":
			method = "GET"
			url = config.MicroserviceURL + "/"
		case "health_check":
			method = "GET"
			url = config.MicroserviceURL + "/healthz"
		case "create_product":
			method = "POST"
			url = config.MicroserviceURL + "/products"
			body, _ = json.Marshal(req.Params)
		case "get_product":
			method = "GET"
			url = fmt.Sprintf("%s/products/%s", config.MicroserviceURL, req.Params["id"])
		case "update_product":
			method = "PUT"
			url = fmt.Sprintf("%s/products/%s", config.MicroserviceURL, req.Params["id"])
			body, _ = json.Marshal(req.Params)
		case "delete_product":
			method = "DELETE"
			url = fmt.Sprintf("%s/products/%s", config.MicroserviceURL, req.Params["id"])
		case "list_products":
			method = "GET"
			url = config.MicroserviceURL + "/products"
		default:
			http.Error(w, "Unsupported tool", http.StatusBadRequest)
			return
		}

		// Call microservice
		client := &http.Client{}
		httpReq, err := http.NewRequest(method, url, bytes.NewBuffer(body))
		if err != nil {
			http.Error(w, "Failed to create request", http.StatusInternalServerError)
			return
		}

		// Add Google Cloud authentication
		token, err := metadata.Get("instance/service-accounts/default/token")
		if err == nil {
			httpReq.Header.Set("Authorization", "Bearer "+token)
		}

		httpReq.Header.Set("Content-Type", "application/json")
		resp, err := client.Do(httpReq)
		if err != nil {
			http.Error(w, "Failed to call microservice", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		// Read response
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Failed to read response", http.StatusInternalServerError)
			return
		}

		// Prepare MCP response
		mcpResp := MCPResponse{}
		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			json.Unmarshal(respBody, &mcpResp.Result)
		} else {
			mcpResp.Error = string(respBody)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(mcpResp)
	}
}

func main() {
	config := Config{
		MicroserviceURL: os.Getenv("MICROSERVICE_URL"), // Set to your microservice URL
		Port:            os.Getenv("PORT"),             // Default to 8080 if not set
	}
	if config.Port == "" {
		config.Port = "8080"
	}

	router := mux.NewRouter()
	router.HandleFunc("/tools", getToolsHandler).Methods("GET")
	router.HandleFunc("/mcp", mcpHandler(config)).Methods("POST")

	log.Printf("Starting MCP server on port %s", config.Port)
	if err := http.ListenAndServe(":"+config.Port, router); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
