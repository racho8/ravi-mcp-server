package main

import (
	"encoding/json"
	"net/http"
)

// Utility functions for MCP server
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
