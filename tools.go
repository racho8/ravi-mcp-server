package main

// MCP tools definition with expected request payloads
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
	{
		Name:        "create_multiple_products",
		Description: "Create multiple products in the store",
		InputSchema: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"products": map[string]interface{}{"type": "array"},
			},
			"required": []string{"products"},
		},
	},
	{
		Name:        "update_products",
		Description: "Update multiple products at once",
		InputSchema: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"products": map[string]interface{}{"type": "array"},
			},
			"required": []string{"products"},
		},
	},
	{
		Name:        "delete_products",
		Description: "Delete multiple products at once",
		InputSchema: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"ids": map[string]interface{}{"type": "array"},
			},
			"required": []string{"ids"},
		},
	},
}
