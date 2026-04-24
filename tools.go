// Package main - tools.go
//
// This file defines all available MCP tools and their schemas for the product management service.
//
// Key Responsibilities:
//   - Define the global 'tools' array with complete metadata for each tool
//   - Specify input schemas for parameter validation
//   - Provide sample requests for documentation and testing
//   - Describe tool capabilities and requirements
//
// Tool Categories:
//   1. Service Health Tools:
//      - welcome_message: Get welcome message
//      - health_check: Check service health
//
//   2. Single Product Operations:
//      - create_product: Create a new product
//      - get_product: Retrieve product by ID
//      - update_product: Update existing product
//      - delete_product: Delete product by ID
//      - list_products: List all products
//
//   3. Batch Operations:
//      - create_multiple_products: Batch create products
//      - update_products: Batch update products
//      - delete_products: Batch delete products
//
//   4. Query/Filter Operations:
//      - get_products_by_category: Filter by category
//      - get_products_by_segment: Filter by segment
//      - get_product_by_name: Search by name
//
// Tool Schema Structure:
//   - Name: Unique identifier for the tool
//   - Description: Human-readable description of tool functionality
//   - InputSchema: JSON schema for parameter validation (JSON Schema format)
//   - Schema: Simplified schema representation
//   - SampleRequest: Example JSON-RPC request with sample parameters
//
// This tools array is referenced by:
//   - handleToolsList() in handlers.go (returns tools to clients)
//   - /mcp/discover endpoint in main.go (REST discovery)
//   - executeToolCall() in business.go (validates tool existence)
package main

// MCP tools definition with expected request payloads
var tools = []ToolSchema{
	{
		Name:        "welcome_message",
		Description: "Use this tool to get a welcome greeting from the product management service. Useful for verifying the server is reachable. Takes no arguments.",
		InputSchema: map[string]interface{}{
			"type":       "object",
			"properties": map[string]interface{}{},
		},
		Schema: map[string]interface{}{},
		SampleRequest: map[string]interface{}{
			"jsonrpc": "2.0",
			"id": "<id>",
			"method": "tools/call",
			"params": map[string]interface{}{
				"name": "welcome_message",
				"arguments": map[string]interface{}{},
			},
		},
	},
	{
		Name:        "health_check",
		Description: "Use this tool to check the health and availability of the product service. Returns the current service status. Call this before other operations to confirm the backend is operational. Takes no arguments.",
		InputSchema: map[string]interface{}{
			"type":       "object",
			"properties": map[string]interface{}{},
		},
		Schema: map[string]interface{}{},
		SampleRequest: map[string]interface{}{
			"jsonrpc": "2.0",
			"id": "<id>",
			"method": "tools/call",
			"params": map[string]interface{}{
				"name": "health_check",
				"arguments": map[string]interface{}{},
			},
		},
	},
	{
		Name:        "create_product",
		Description: "Use this tool to create a single new product in the catalog. Requires name, category, and price. Optionally accepts a segment. Returns the created product with its generated ID.",
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
		Schema: map[string]interface{}{
			"name":     "string",
			"category": "string",
			"segment":  "string",
			"price":    "number",
		},
		SampleRequest: map[string]interface{}{
			"jsonrpc": "2.0",
			"id": "<id>",
			"method": "tools/call",
			"params": map[string]interface{}{
				"name": "create_product",
				"arguments": map[string]interface{}{
					"name": "<product name>",
					"category": "<category name>",
					"segment": "<segment name>",
					"price": "<price>",
				},
			},
		},
	},
	{
		Name:        "get_product",
		Description: "Use this tool to retrieve a single product by its unique ID. Returns full product details including name, category, segment, and price.",
		InputSchema: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"id": map[string]string{"type": "string"},
			},
			"required": []string{"id"},
		},
		Schema: map[string]interface{}{
			"id": "string",
		},
		SampleRequest: map[string]interface{}{
			"jsonrpc": "2.0",
			"id": "<id>",
			"method": "tools/call",
			"params": map[string]interface{}{
				"name": "get_product",
				"arguments": map[string]interface{}{
					"id": "12345",
				},
			},
		},
	},
	{
		Name:        "update_product",
		Description: "Use this tool to update a single existing product by its ID. Only the provided fields (name, price, category) are modified; omitted fields remain unchanged. ID is required, all other fields are optional.",
		InputSchema: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"id":    map[string]string{"type": "string"},
				"name":  map[string]string{"type": "string"},
				"price": map[string]string{"type": "number"},
				"category": map[string]string{"type": "string"},
			},
			"required": []string{"id"},
		},
		Schema: map[string]interface{}{
			"id": "string",
			"name": "string",
			"price": "number",
			"category": "string",
		},
		SampleRequest: map[string]interface{}{
			"jsonrpc": "2.0",
			"id": "<id>",
			"method": "tools/call",
			"params": map[string]interface{}{
				"name": "update_product",
				"arguments": map[string]interface{}{
					"id": "12345",
					"name": "Laptop5",
					"category": "<category name>",
					"price": 1099,
				},
			},
		},
	},
	{
		Name:        "delete_product",
		Description: "Use this tool to permanently delete a single product from the catalog by its ID. This action cannot be undone.",
		InputSchema: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"id": map[string]string{"type": "string"},
			},
			"required": []string{"id"},
		},
		Schema: map[string]interface{}{
			"id": "string",
		},
		SampleRequest: map[string]interface{}{
			"jsonrpc": "2.0",
			"id": "<id>",
			"method": "tools/call",
			"params": map[string]interface{}{
				"name": "delete_product",
				"arguments": map[string]interface{}{
					"id": "12345",
				},
			},
		},
	},
	{
		Name:        "list_products",
		Description: "Use this tool to list all products in the catalog. Returns an array of every product with full details including ID, name, category, segment, and price. Takes no arguments. For follow-up questions about a specific product's price, category, or details, use get_product_by_name instead of calling this again.",
		InputSchema: map[string]interface{}{
			"type":       "object",
			"properties": map[string]interface{}{},
		},
		Schema: map[string]interface{}{},
		SampleRequest: map[string]interface{}{
			"jsonrpc": "2.0",
			"id": "<id>",
			"method": "tools/call",
			"params": map[string]interface{}{
				"name": "list_products",
				"arguments": map[string]interface{}{},
			},
		},
	},
	{
		Name:        "create_multiple_products",
		Description: "Use this tool to create multiple products in a single batch operation. Accepts an array of product objects, each with name, category, segment, and price. Prefer this over repeated create_product calls.",
		InputSchema: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"products": map[string]interface{}{"type": "array"},
			},
			"required": []string{"products"},
		},
		Schema: map[string]interface{}{
			"products": "array of product objects",
		},
		SampleRequest: map[string]interface{}{
			"jsonrpc": "2.0",
			"id": "<id>",
			"method": "tools/call",
			"params": map[string]interface{}{
				"name": "create_multiple_products",
				"arguments": map[string]interface{}{
					"products": []map[string]interface{}{
						{
							"name": "<product name>",
							"category": "<category name>",
							"segment": "<segment>",
							"price": "<price>",
						},
						{
							"name": "<product name>",
							"category": "<category name>",
							"segment": "<segment>",
							"price": "<price>",
						},
					},
				},
			},
		},
	},
	{
		Name:        "update_products",
		Description: "Use this tool to update multiple products in a single batch operation. Accepts an array of product objects, each identified by its ID with the fields to update. Prefer this over repeated update_product calls.",
		InputSchema: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"products": map[string]interface{}{"type": "array"},
			},
			"required": []string{"products"},
		},
		Schema: map[string]interface{}{
			"products": "array of product update objects",
		},
		SampleRequest: map[string]interface{}{
			"jsonrpc": "2.0",
			"id": "<id>",
			"method": "tools/call",
			"params": map[string]interface{}{
				"name": "update_products",
				"arguments": map[string]interface{}{
					"products": []map[string]interface{}{
						{
							"id": "<product id>",
							"category": "<category name>",
							"segment": "<segment>",
							"price": "<price>",
						},
						{
							"id": "<product id>",
							"category": "<category name>",
							"segment": "<segment>",
							"price": "<price>",
						},
					},
				},
			},
		},
	},
	{
		Name:        "delete_products",
		Description: "Use this tool to permanently delete multiple products in a single batch operation. Accepts an array of product IDs. This action cannot be undone. Prefer this over repeated delete_product calls.",
		InputSchema: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"ids": map[string]interface{}{"type": "array"},
			},
			"required": []string{"ids"},
		},
		Schema: map[string]interface{}{
			"ids": "array of product ids",
		},
		SampleRequest: map[string]interface{}{
			"jsonrpc": "2.0",
			"id": "<id>",
			"method": "tools/call",
			"params": map[string]interface{}{
				"name": "delete_products",
				"arguments": map[string]interface{}{
					"ids": []string{"<product id 1>", "<product id 2>"},
				},
			},
		},
	},
	{
		Name:        "get_products_by_category",
		Description: "Use this tool to filter products by category (e.g., Electronics, Clothing, Food). Returns an array of all products belonging to the specified category with full details.",
		InputSchema: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"category": map[string]string{"type": "string"},
			},
			"required": []string{"category"},
		},
		Schema: map[string]interface{}{
			"category": "string",
		},
		SampleRequest: map[string]interface{}{
			"jsonrpc": "2.0",
			"id": "<id>",
			"method": "tools/call",
			"params": map[string]interface{}{
				"name": "get_products_by_category",
				"arguments": map[string]interface{}{
					"category": "<category name>",
				},
			},
		},
	},
	{
		Name:        "get_products_by_segment",
		Description: "Use this tool to filter products by market segment (e.g., Premium, Budget, Enterprise). Returns an array of all products belonging to the specified segment with full details.",
		InputSchema: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"segment": map[string]string{"type": "string"},
			},
			"required": []string{"segment"},
		},
		Schema: map[string]interface{}{
			"segment": "string",
		},
		SampleRequest: map[string]interface{}{
			"jsonrpc": "2.0",
			"id": "<id>",
			"method": "tools/call",
			"params": map[string]interface{}{
				"name": "get_products_by_segment",
				"arguments": map[string]interface{}{
					"segment": "<segment name>",
				},
			},
		},
	},
	{
		Name:        "get_product_by_name",
		Description: "Use this tool to look up a specific product by its name and get its full details including price, category, and segment. Call this when the user asks about a specific product's price, availability, or details. For example: 'What is the price of iPhone 17?' or 'Tell me about Laptop5'.",
		InputSchema: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"name": map[string]string{"type": "string"},
			},
			"required": []string{"name"},
		},
		Schema: map[string]interface{}{
			"name": "string",
		},
		SampleRequest: map[string]interface{}{
			"jsonrpc": "2.0",
			"id": "<id>",
			"method": "tools/call",
			"params": map[string]interface{}{
				"name": "get_product_by_name",
				"arguments": map[string]interface{}{
					"name": "<product name>",
				},
			},
		},
	},
}
