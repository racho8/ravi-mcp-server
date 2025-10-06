// Package main - utils.go
//
// This file provides utility functions for handling JSON-RPC 2.0 responses in the MCP server.
//
// Key Responsibilities:
//   - Format and send successful JSON-RPC responses
//   - Format and send JSON-RPC error responses
//   - Set appropriate HTTP headers (Content-Type, CORS)
//   - Encode responses as JSON
//
// Utility Functions:
//
//   1. sendJSONRPCResponse:
//      - Constructs successful JSON-RPC 2.0 response
//      - Sets Content-Type header to application/json
//      - Encodes result data and sends to client
//      - Used by all successful handler operations
//
//   2. sendJSONRPCError:
//      - Constructs JSON-RPC 2.0 error response
//      - Includes error code, message, and optional data
//      - Sets Content-Type header to application/json
//      - Used for all error scenarios (parse errors, invalid requests, etc.)
//
// JSON-RPC 2.0 Response Format:
//   Success: { "jsonrpc": "2.0", "id": <request_id>, "result": <data> }
//   Error:   { "jsonrpc": "2.0", "id": <request_id>, "error": { "code": <code>, "message": <msg>, "data": <details> } }
//
// Usage:
//   - Called by all handlers in handlers.go
//   - Ensures consistent response formatting across all endpoints
//   - Maintains JSON-RPC 2.0 specification compliance
package main

import (
	"encoding/json"
	"net/http"
)

// utility functions for MCP server
func sendJSONRPCResponse(w http.ResponseWriter, id interface{}, result interface{}) {
	response := JSONRPCResponse{
		JSONRPC: "2.0",
		ID:      id,
		Result:  result,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

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
