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
		Description: "Check the health of the product service",
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
		Description: "Retrieve a product by ID",
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
		Schema: map[string]interface{}{
			"id": "string",
			"name": "string",
			"price": "number",
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
		Description: "Delete a product by ID",
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
		Description: "List all products in the store",
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
		Description: "Create multiple products in the store",
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
		Description: "Update multiple products at once",
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
		Description: "Delete multiple products at once",
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
		Description: "Retrieve all products matching a given category",
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
		Description: "Retrieve all products matching a given segment",
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
		Description: "Retrieve all products matching a given name",
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
