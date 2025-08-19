package main

// All struct definitions for MCP server
type JSONRPCRequest struct {
	JSONRPC string      `json:"jsonrpc"`
	ID      interface{} `json:"id"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params,omitempty"`
}

type JSONRPCResponse struct {
	JSONRPC string      `json:"jsonrpc"`
	ID      interface{} `json:"id"`
	Result  interface{} `json:"result,omitempty"`
	Error   *JSONRPCError `json:"error,omitempty"`
}

type JSONRPCError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ToolSchema struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	InputSchema interface{} `json:"inputSchema"`
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
	ProtocolVersion string         `json:"protocolVersion"`
	Capabilities    ServerCapabilities `json:"capabilities"`
	ServerInfo      ServerInfo     `json:"serverInfo"`
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

type ToolCallResult struct {
	Content []ToolContent `json:"content"`
}

type ToolContent struct {
	Type string      `json:"type"`
	Text string      `json:"text,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

type Config struct {
	MicroserviceURL string
	Port            string
}
