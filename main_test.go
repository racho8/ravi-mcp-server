package main

import (
	"testing"
)

func TestConfig(t *testing.T) {
	config := Config{
		MicroserviceURL: "https://test.example.com",
		Port:            "8080",
	}

	if config.MicroserviceURL != "https://test.example.com" {
		t.Errorf("Expected MicroserviceURL to be 'https://test.example.com', got '%s'", config.MicroserviceURL)
	}

	if config.Port != "8080" {
		t.Errorf("Expected Port to be '8080', got '%s'", config.Port)
	}
}

func TestToolsAvailable(t *testing.T) {
	expectedTools := []string{"welcome_message", "health_check", "create_product", "get_product", "update_product", "delete_product", "list_products"}
	
	if len(tools) != len(expectedTools) {
		t.Errorf("Expected %d tools, got %d", len(expectedTools), len(tools))
	}

	for i, tool := range tools {
		if i < len(expectedTools) && tool.Name != expectedTools[i] {
			t.Errorf("Expected tool %d to be '%s', got '%s'", i, expectedTools[i], tool.Name)
		}
	}
}

func TestJSONRPCStructs(t *testing.T) {
	// Test JSONRPCRequest
	req := JSONRPCRequest{
		JSONRPC: "2.0",
		ID:      1,
		Method:  "test",
	}

	if req.JSONRPC != "2.0" {
		t.Errorf("Expected JSONRPC to be '2.0', got '%s'", req.JSONRPC)
	}

	// Test JSONRPCResponse
	resp := JSONRPCResponse{
		JSONRPC: "2.0",
		ID:      1,
		Result:  "success",
	}

	if resp.JSONRPC != "2.0" {
		t.Errorf("Expected JSONRPC to be '2.0', got '%s'", resp.JSONRPC)
	}
}
