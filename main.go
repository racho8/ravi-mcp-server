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

// --- MCP Protocol Structs ---

// ToolSchema defines the structure of an MCP tool, as returned by the server.
type ToolSchema struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Parameters  interface{} `json:"parameters"`
	Returns     interface{} `json:"returns"`
}

// InitializeRequest is the first request a client sends to the server for initial handshake.
// This differs from the tool invocation requests and is used to establish the connection.
type InitializeRequest struct {
}

// ToolInvocationRequest represents a request to call a specific tool.
type ToolInvocationRequest struct {
	Name   string      `json:"name"`
	Params interface{} `json:"params"`
}

// MCPRequest is the top-level request from the client.
type MCPRequest struct {
	StreamID   string                 `json:"streamId"`
	Initialize *InitializeRequest     `json:"initialize,omitempty"`
	ToolCode   *ToolInvocationRequest `json:"tool_code,omitempty"`
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
				"name":     map[string]string{"type": "string"},
				"category": map[string]string{"type": "string"},
				"segment":  map[string]string{"type": "string"},
				"price":    map[string]string{"type": "number"},
			},
			"required": []string{"name", "category", "price"},
		},
		Returns: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"id":       map[string]string{"type": "string"},
				"name":     map[string]string{"type": "string"},
				"category": map[string]string{"type": "string"},
				"segment":  map[string]string{"type": "string"},
				"price":    map[string]string{"type": "number"},
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
				"id":       map[string]string{"type": "string"},
				"name":     map[string]string{"type": "string"},
				"category": map[string]string{"type": "string"},
				"segment":  map[string]string{"type": "string"},
				"price":    map[string]string{"type": "number"},
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
				"id":       map[string]string{"type": "string"},
				"name":     map[string]string{"type": "string"},
				"category": map[string]string{"type": "string"},
				"segment":  map[string]string{"type": "string"},
				"price":    map[string]string{"type": "number"},
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
					"id":       map[string]string{"type": "string"},
					"name":     map[string]string{"type": "string"},
					"category": map[string]string{"type": "string"},
					"segment":  map[string]string{"type": "string"},
					"price":    map[string]string{"type": "number"},
				},
			},
		},
	},
}

// mcpRouter is the main handler for all POST requests to the /mcp endpoint.
// It inspects the request body to determine if it's an 'initialize' request or
// a 'tool_code' invocation and routes to the appropriate handler.
func mcpRouter(config Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusBadRequest)
			return
		}

		var req MCPRequest
		if err := json.Unmarshal(body, &req); err != nil {
			http.Error(w, "Invalid JSON in request body", http.StatusBadRequest)
			return
		}

		r.Body = io.NopCloser(bytes.NewBuffer(body))

		log.Printf("Received request: %+v", req)

		if req.Initialize != nil {

			// This is the initial handshake from the client
			handleInitialize(w, r)
		} else if req.ToolCode != nil {

			// This is a tool invocation request
			handleToolInvocation(w, r, config, req.ToolCode)
		} else {

			// Handle unsupported or malformed requests
			http.Error(w, "Unsupported request type", http.StatusBadRequest)
		}
	}
}

// handleInitialize handles the initial handshake from the client
func handleInitialize(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"tools": tools,
	})
	log.Println("Sent tool list to client.")
}

// handleToolInvocation handles requests to invoke a specific tool
func handleToolInvocation(w http.ResponseWriter, r *http.Request, config Config, toolInvocation *ToolInvocationRequest) {
	log.Printf("Received tool invocation request: %+v", *toolInvocation)

	// map tools to microservice endpoints
	method, url, body, done := mapToolsToEndpoints(w, toolInvocation, config)
	if done {
		return
	}

	// Call the microservice, e.g., product service
	client := &http.Client{}
	httpReq, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}

	log.Printf("Calling microservice: Method=%s, URL=%s", method, url)

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

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response", http.StatusInternalServerError)
		return
	}

	mcpResp := MCPResponse{}
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {

		// handle responses based on the tool name
		switch toolInvocation.Name {
		case "list_products":
			var products []map[string]interface{}
			if err := json.Unmarshal(respBody, &products); err != nil {
				log.Printf("Failed to unmarshal list_products response: %v", err)
				mcpResp.Error = "Failed to parse microservice response"
			} else {
				mcpResp.Result = products
			}
		case "health_check":

			// Health check returns a simple string, so we'll just return it directly
			mcpResp.Result = string(respBody)
		default:

			// For all other endpoints, assume a JSON object response
			var result map[string]interface{}
			if err := json.Unmarshal(respBody, &result); err != nil {
				log.Printf("Failed to unmarshal JSON response: %v", err)
				mcpResp.Error = "Failed to parse microservice response"
			} else {
				mcpResp.Result = result
			}
		}
	} else {
		mcpResp.Error = string(respBody)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mcpResp)
}

// mapToolsToEndpoints maps the tool invocation to the correct microservice endpoint.
func mapToolsToEndpoints(w http.ResponseWriter, req *ToolInvocationRequest, config Config) (string, string, []byte, bool) {
	var method, url string
	var body []byte

	params, ok := req.Params.(map[string]interface{})
	if !ok {
		http.Error(w, "Invalid params format", http.StatusBadRequest)
		return "", "", nil, true
	}

	switch req.Name {
	case "welcome_message":
		method = "GET"
		url = config.MicroserviceURL + "/"
	case "health_check":
		method = "GET"
		url = config.MicroserviceURL + "/healthz" // TODO: Fix it , this is not working
	case "create_product":
		method = "POST"
		url = config.MicroserviceURL + "/products"
		body, _ = json.Marshal(params)
	case "get_product":
		method = "GET"
		url = fmt.Sprintf("%s/products/%s", config.MicroserviceURL, params["id"])
	case "update_product":
		method = "PUT"
		url = fmt.Sprintf("%s/products/%s", config.MicroserviceURL, params["id"])
		body, _ = json.Marshal(params)
	case "delete_product":
		method = "DELETE"
		url = fmt.Sprintf("%s/products/%s", config.MicroserviceURL, params["id"])
	case "list_products":
		method = "GET"
		url = config.MicroserviceURL + "/products"
	default:
		http.Error(w, "Unsupported tool", http.StatusBadRequest)
		return "", "", nil, true
	}
	return method, url, body, false
}

func main() {
	log.Printf("MICROSERVICE_URL: %s", os.Getenv("MICROSERVICE_URL"))
	log.Printf("PORT: %s", os.Getenv("PORT"))

	config := Config{
		MicroserviceURL: os.Getenv("MICROSERVICE_URL"),
		Port:            os.Getenv("PORT"),
	}
	if config.Port == "" {
		config.Port = "8080"
	}

	router := mux.NewRouter()
	router.HandleFunc("/mcp", mcpRouter(config)).Methods("POST")

	log.Printf("Starting MCP server on port %s", config.Port)
	if err := http.ListenAndServe(":"+config.Port, router); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
