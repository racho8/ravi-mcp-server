package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// All HTTP handler functions for MCP server
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

		if req.JSONRPC != "2.0" {
			sendJSONRPCError(w, req.ID, -32600, "Invalid Request", "Invalid JSON-RPC version")
			return
		}

		log.Printf("Received JSON-RPC request: method=%s, id=%v", req.Method, req.ID)

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

func handleToolsList(w http.ResponseWriter, req JSONRPCRequest) {
	result := ToolsListResult{
		Tools: tools,
	}

	sendJSONRPCResponse(w, req.ID, result)
	log.Println("Sent tools list to client.")
}

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

	// pass tool name and arguments only to executeToolCall
	args, _ := params.Arguments.(map[string]interface{})
	result, err := executeToolCall(params.Name, args)
	if err != nil {
		sendJSONRPCError(w, req.ID, -32603, "Internal error", err.Error())
		return
	}

	sendJSONRPCResponse(w, req.ID, result)
}
