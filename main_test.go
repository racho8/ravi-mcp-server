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
	// List of all tool names currently registered in tools.go
	// Update this list if you add/remove tools in tools.go
	expectedTools := []string{
		"welcome_message",           // Product service: welcome message
		"health_check",              // Product service: health check
		"create_product",            // Product service: create product
		"get_product",               // Product service: get product by ID
		"update_product",            // Product service: update product by ID
		"delete_product",            // Product service: delete product by ID
		"list_products",             // Product service: list all products
		"create_multiple_products",  // Product service: batch create
		"update_products",           // Product service: batch update
		"delete_products",           // Product service: batch delete
		// rpim-api-service tools below
		"rpim_health_check",         // RPIM: health check
		"rpim_get_child_items",      // RPIM: get child items
		"rpim_get_item_keys",        // RPIM: get item keys
		"rpim_get_class_units",      // RPIM: get class units
		"rpim_get_item_details",     // RPIM: get item details
		"rpim_get_local_items",      // RPIM: get local items
		"rpim_get_updated_items",    // RPIM: get updated items
		"rpim_get_classified_items", // RPIM: get classified items
		"rpim_get_item_attributes",  // RPIM: get item attributes
		"rpim_get_items_by_class_unit", // RPIM: get items by class unit
	}

	// Ensure the number of tools matches
	if len(tools) != len(expectedTools) {
		t.Errorf("Expected %d tools, got %d. If you added/removed tools, update expectedTools in main_test.go.", len(expectedTools), len(tools))
	}

	// Ensure tool names match
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
