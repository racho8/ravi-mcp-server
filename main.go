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

// --- MCP JSON-RPC Protocol Structs ---

// JSONRPCRequest represents a JSON-RPC 2.0 request
type JSONRPCRequest struct {
	JSONRPC string      `json:"jsonrpc"`
	ID      interface{} `json:"id"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params,omitempty"`
}

// JSONRPCResponse represents a JSON-RPC 2.0 response
type JSONRPCResponse struct {
	JSONRPC string      `json:"jsonrpc"`
	ID      interface{} `json:"id"`
	Result  interface{} `json:"result,omitempty"`
	Error   *JSONRPCError `json:"error,omitempty"`
}

// JSONRPCError represents a JSON-RPC 2.0 error
type JSONRPCError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// ToolSchema defines the structure of an MCP tool
type ToolSchema struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	InputSchema interface{} `json:"inputSchema"`
}

// InitializeParams for the initialize method
type InitializeParams struct {
	ProtocolVersion string      `json:"protocolVersion"`
	Capabilities    interface{} `json:"capabilities"`
	ClientInfo      ClientInfo  `json:"clientInfo"`
}

// ClientInfo contains information about the client
type ClientInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// InitializeResult for the initialize response
type InitializeResult struct {
	ProtocolVersion string         `json:"protocolVersion"`
	Capabilities    ServerCapabilities `json:"capabilities"`
	ServerInfo      ServerInfo     `json:"serverInfo"`
}

// ServerCapabilities defines what the server can do
type ServerCapabilities struct {
	Tools interface{} `json:"tools,omitempty"`
}

// ServerInfo contains information about the server
type ServerInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// ToolsListResult for tools/list response
type ToolsListResult struct {
	Tools []ToolSchema `json:"tools"`
}

// ToolCallParams for tools/call method
type ToolCallParams struct {
	Name      string      `json:"name"`
	Arguments interface{} `json:"arguments,omitempty"`
}

// ToolCallResult for tools/call response
type ToolCallResult struct {
	Content []ToolContent `json:"content"`
}

// ToolContent represents the result content from a tool call
type ToolContent struct {
	Type string      `json:"type"`
	Text string      `json:"text,omitempty"`
	Data interface{} `json:"data,omitempty"`
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
		InputSchema: map[string]interface{}{
			"type":       "object",
			"properties": map[string]interface{}{},
		},
	},
	{
		Name:        "health_check",
		Description: "Check the health of the product service",
		InputSchema: map[string]interface{}{
			"type":       "object",
			"properties": map[string]interface{}{},
		},
	},
	{
		Name:        "create_product",
		Description: "Create a new product in the store",
		InputSchema: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"name":     map[string]string{"type": "string"},
				"category": map[string]string{"type": "string"},
				"segment":  map[string]string{"type": "string"},
				"price":    map[string]string{"type": "number"},
			},
			"required": []string{"name", "category", "price"},
		},
	},
	{
		Name:        "get_product",
		Description: "Retrieve a product by ID",
		InputSchema: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"id": map[string]string{"type": "string"},
			},
			"required": []string{"id"},
		},
	},
	{
		Name:        "update_product",
		Description: "Update an existing product by ID",
		InputSchema: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"id":    map[string]string{"type": "string"},
				"name":  map[string]string{"type": "string"},
				"price": map[string]string{"type": "number"},
			},
			"required": []string{"id"},
		},
	},
	{
		Name:        "delete_product",
		Description: "Delete a product by ID",
		InputSchema: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"id": map[string]string{"type": "string"},
			},
			"required": []string{"id"},
		},
	},
	{
		Name:        "list_products",
		Description: "List all products in the store",
		InputSchema: map[string]interface{}{
			"type":       "object",
			"properties": map[string]interface{}{},
		},
	},
}

// mcpHandler is the main handler for all JSON-RPC requests to the /mcp endpoint.
// It implements the MCP JSON-RPC 2.0 protocol with proper method routing.
func mcpHandler(config Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			sendJSONRPCError(w, nil, -32700, "Parse error", "Failed to read request body")
			return
		}

		var req JSONRPCRequest
		if err := json.Unmarshal(body, &req); err != nil {
			sendJSONRPCError(w, nil, -32700, "Parse error", "Invalid JSON")
			return
		}

		// Validate JSON-RPC 2.0 format
		if req.JSONRPC != "2.0" {
			sendJSONRPCError(w, req.ID, -32600, "Invalid Request", "Invalid JSON-RPC version")
			return
		}

		log.Printf("Received JSON-RPC request: method=%s, id=%v", req.Method, req.ID)

		// Route to appropriate handler based on method
		switch req.Method {
		case "initialize":
			handleInitialize(w, req)
		case "tools/list":
			handleToolsList(w, req)
		case "tools/call":
			handleToolCall(w, req, config)
		default:
			sendJSONRPCError(w, req.ID, -32601, "Method not found", fmt.Sprintf("Unknown method: %s", req.Method))
		}
	}
}

// handleInitialize handles the initial handshake from the client
func handleInitialize(w http.ResponseWriter, req JSONRPCRequest) {
	var params InitializeParams
	if req.Params != nil {
		paramBytes, _ := json.Marshal(req.Params)
		if err := json.Unmarshal(paramBytes, &params); err != nil {
			sendJSONRPCError(w, req.ID, -32602, "Invalid params", "Failed to parse initialize params")
			return
		}
	}

	result := InitializeResult{
		ProtocolVersion: "2024-11-05",
		Capabilities: ServerCapabilities{
			Tools: map[string]interface{}{},
		},
		ServerInfo: ServerInfo{
			Name:    "ravi-mcp-server",
			Version: "1.0.0",
		},
	}

	sendJSONRPCResponse(w, req.ID, result)
	log.Println("Sent initialize response to client.")
}

// handleToolsList handles requests to list available tools
func handleToolsList(w http.ResponseWriter, req JSONRPCRequest) {
	result := ToolsListResult{
		Tools: tools,
	}

	sendJSONRPCResponse(w, req.ID, result)
	log.Println("Sent tools list to client.")
}

// handleToolCall handles requests to invoke a specific tool
func handleToolCall(w http.ResponseWriter, req JSONRPCRequest, config Config) {
	var params ToolCallParams
	if req.Params == nil {
		sendJSONRPCError(w, req.ID, -32602, "Invalid params", "Missing tool call parameters")
		return
	}

	paramBytes, _ := json.Marshal(req.Params)
	if err := json.Unmarshal(paramBytes, &params); err != nil {
		sendJSONRPCError(w, req.ID, -32602, "Invalid params", "Failed to parse tool call params")
		return
	}

	log.Printf("Received tool call: %s", params.Name)

	// Map tool to microservice endpoint and execute
	result, err := executeToolCall(params, config)
	if err != nil {
		sendJSONRPCError(w, req.ID, -32603, "Internal error", err.Error())
		return
	}

	sendJSONRPCResponse(w, req.ID, result)
}

// executeToolCall executes the actual tool call against the microservice
func executeToolCall(params ToolCallParams, config Config) (ToolCallResult, error) {
	// Map tools to microservice endpoints
	method, url, body, err := mapToolsToEndpoints(params, config)
	if err != nil {
		return ToolCallResult{}, err
	}

	// Call the microservice
	client := &http.Client{}
	httpReq, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return ToolCallResult{}, fmt.Errorf("failed to create request: %v", err)
	}

	log.Printf("Calling microservice: Method=%s, URL=%s", method, url)

	// Add authentication token if available
	token, err := metadata.Get("instance/service-accounts/default/token")
	if err == nil {
		httpReq.Header.Set("Authorization", "Bearer "+token)
	}

	httpReq.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(httpReq)
	if err != nil {
		return ToolCallResult{}, fmt.Errorf("failed to call microservice: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return ToolCallResult{}, fmt.Errorf("failed to read response: %v", err)
	}

	// Create tool result based on response
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		// Success response
		var content ToolContent
		
		switch params.Name {
		case "health_check":
			// Health check might return plain text, so include as text
			content = ToolContent{
				Type: "text",
				Text: string(respBody),
			}
		default:
			// For JSON responses, prefer structured data over raw text
			var jsonData interface{}
			if json.Unmarshal(respBody, &jsonData) == nil {
				// Successfully parsed as JSON, use structured data only
				content = ToolContent{
					Type: "text",
					Data: jsonData,
				}
			} else {
				// Not valid JSON, fallback to text
				content = ToolContent{
					Type: "text",
					Text: string(respBody),
				}
			}
		}

		return ToolCallResult{
			Content: []ToolContent{content},
		}, nil
	} else {
		// Error response
		return ToolCallResult{
			Content: []ToolContent{
				{
					Type: "text",
					Text: fmt.Sprintf("Error: %s", string(respBody)),
				},
			},
		}, nil
	}
}

// mapToolsToEndpoints maps the tool call to the correct microservice endpoint
func mapToolsToEndpoints(params ToolCallParams, config Config) (string, string, []byte, error) {
	var method, url string
	var body []byte

	arguments, ok := params.Arguments.(map[string]interface{})
	if !ok && params.Arguments != nil {
		return "", "", nil, fmt.Errorf("invalid arguments format")
	}

	switch params.Name {
	case "welcome_message":
		method = "GET"
		url = config.MicroserviceURL + "/"
	case "health_check":
		method = "GET"
		url = config.MicroserviceURL + "/healthz"
	case "create_product":
		method = "POST"
		url = config.MicroserviceURL + "/products"
		if arguments != nil {
			body, _ = json.Marshal(arguments)
		}
	case "get_product":
		method = "GET"
		if arguments == nil || arguments["id"] == nil {
			return "", "", nil, fmt.Errorf("missing required parameter: id")
		}
		url = fmt.Sprintf("%s/products/%s", config.MicroserviceURL, arguments["id"])
	case "update_product":
		method = "PUT"
		if arguments == nil || arguments["id"] == nil {
			return "", "", nil, fmt.Errorf("missing required parameter: id")
		}
		url = fmt.Sprintf("%s/products/%s", config.MicroserviceURL, arguments["id"])
		body, _ = json.Marshal(arguments)
	case "delete_product":
		method = "DELETE"
		if arguments == nil || arguments["id"] == nil {
			return "", "", nil, fmt.Errorf("missing required parameter: id")
		}
		url = fmt.Sprintf("%s/products/%s", config.MicroserviceURL, arguments["id"])
	case "list_products":
		method = "GET"
		url = config.MicroserviceURL + "/products"
	default:
		return "", "", nil, fmt.Errorf("unsupported tool: %s", params.Name)
	}

	return method, url, body, nil
}

// sendJSONRPCResponse sends a successful JSON-RPC response
func sendJSONRPCResponse(w http.ResponseWriter, id interface{}, result interface{}) {
	response := JSONRPCResponse{
		JSONRPC: "2.0",
		ID:      id,
		Result:  result,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// sendJSONRPCError sends an error JSON-RPC response
func sendJSONRPCError(w http.ResponseWriter, id interface{}, code int, message string, data interface{}) {
	response := JSONRPCResponse{
		JSONRPC: "2.0",
		ID:      id,
		Error: &JSONRPCError{
			Code:    code,
			Message: message,
			Data:    data,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// Handle preflight OPTIONS request
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
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
	
	// Add CORS middleware
	router.Use(corsMiddleware)
	
	// Health check endpoint for monitoring and testing
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"status": "healthy",
			"service": "ravi-mcp-server",
		})
	}).Methods("GET")
	
	// Add OPTIONS handler for preflight requests
	router.HandleFunc("/mcp", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods("OPTIONS")
	
	router.HandleFunc("/mcp", mcpHandler(config)).Methods("POST")

	log.Printf("Starting MCP server on port %s", config.Port)
	log.Printf("MCP JSON-RPC 2.0 Protocol supported methods:")
	log.Printf("  - initialize")
	log.Printf("  - tools/list")
	log.Printf("  - tools/call")
	
	if err := http.ListenAndServe(":"+config.Port, router); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
