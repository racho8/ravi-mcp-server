// Package main - models.go
//
// This file defines all data structures and models used throughout the MCP server.
//
// Key Responsibilities:
//   - Define JSON-RPC 2.0 protocol structures (request, response, error)
//   - Define MCP-specific structures (tool schemas, initialization, capabilities)
//   - Define server configuration structures
//   - Provide Go struct tags for JSON marshaling/unmarshaling
//
// Structure Categories:
//
//   1. JSON-RPC 2.0 Protocol:
//      - JSONRPCRequest: Standard JSON-RPC request format
//      - JSONRPCResponse: Standard JSON-RPC response format
//      - JSONRPCError: Standard error structure with code, message, and data
//
//   2. MCP Protocol Structures:
//      - InitializeParams: Client initialization parameters
//      - InitializeResult: Server initialization response with capabilities
//      - ToolSchema: Complete tool definition with schema and metadata
//      - ToolCallParams: Parameters for executing a tool
//
//   3. Capability Structures:
//      - ServerCapabilities: Advertised server capabilities
//      - ClientInfo: Client identification information
//      - ServerInfo: Server identification information
//
//   4. Configuration:
//      - Config: Server configuration (microservice URL, port)
//
// JSON Tags:
//   - All structs include `json` tags for proper serialization
//   - Optional fields marked with `omitempty`
//
// Usage:
//   - handlers.go: Unmarshals requests and marshals responses
//   - utils.go: Uses response structures for formatting
//   - main.go: Uses Config for server initialization
package main

// data models or struct definitions for the MCP server

// JSONRPCRequest and JSONRPCResponse structures
type JSONRPCRequest struct {
	JSONRPC string      `json:"jsonrpc"`
	ID      interface{} `json:"id"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params,omitempty"`
}

type JSONRPCResponse struct {
	JSONRPC string        `json:"jsonrpc"`
	ID      interface{}   `json:"id"`
	Result  interface{}   `json:"result,omitempty"`
	Error   *JSONRPCError `json:"error,omitempty"`
}

type JSONRPCError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ToolSchema struct {
	Name          string      `json:"name"`
	Description   string      `json:"description"`
	InputSchema   interface{} `json:"inputSchema"`
	Schema        interface{} `json:"schema,omitempty"`
	SampleRequest interface{} `json:"sampleRequest,omitempty"`
}

type InitializeParams struct {
	ProtocolVersion string      `json:"protocolVersion"`
	Capabilities    interface{} `json:"capabilities"`
	ClientInfo      ClientInfo  `json:"clientInfo"`
}

type ClientInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type InitializeResult struct {
	ProtocolVersion string             `json:"protocolVersion"`
	Capabilities    ServerCapabilities `json:"capabilities"`
	ServerInfo      ServerInfo         `json:"serverInfo"`
}

type ServerCapabilities struct {
	Tools interface{} `json:"tools,omitempty"`
}

type ServerInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type ToolsListResult struct {
	Tools []ToolSchema `json:"tools"`
}

type ToolCallParams struct {
	Name      string      `json:"name"`
	Arguments interface{} `json:"arguments,omitempty"`
}

type Config struct {
	MicroserviceURL string
	Port            string
}
