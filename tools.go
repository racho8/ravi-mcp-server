package main



// MCP tools definition
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
	// RPIM Tools
	{
		Name:        "rpim_health_check",
		Description: "Health check for rpim-api-service",
		InputSchema: map[string]interface{}{
			"type":       "object",
			"properties": map[string]interface{}{},
		},
	},
	{
		Name:        "rpim_get_child_items",
		Description: "Get child items from rpim-api-service",
		InputSchema: map[string]interface{}{
			"type": "array",
			"items": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"itemNo":         map[string]string{"type": "string"},
					"itemType":       map[string]string{"type": "string"},
					"classUnitCode":  map[string]string{"type": "string"},
					"classUnitType":  map[string]string{"type": "string"},
				},
				"required": []string{"itemNo", "itemType", "classUnitCode", "classUnitType"},
			},
		},
	},
	{
		Name:        "rpim_get_item_keys",
		Description: "Get item keys from rpim-api-service",
		InputSchema: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"classUnitCode": map[string]string{"type": "string"},
			},
			"required": []string{"classUnitCode"},
		},
	},
	{
		Name:        "rpim_get_class_units",
		Description: "Get class units from rpim-api-service",
		InputSchema: map[string]interface{}{
			"type":       "object",
			"properties": map[string]interface{}{},
		},
	},
	{
		Name:        "rpim_get_item_details",
		Description: "Get item details from rpim-api-service",
		InputSchema: map[string]interface{}{
			"type": "array",
			"items": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"itemNo":         map[string]string{"type": "string"},
					"itemType":       map[string]string{"type": "string"},
					"classUnitCode":  map[string]string{"type": "string"},
					"classUnitType":  map[string]string{"type": "string"},
				},
				"required": []string{"itemNo", "itemType", "classUnitCode", "classUnitType"},
			},
		},
	},
	{
		Name:        "rpim_get_local_items",
		Description: "Get local items from rpim-api-service",
		InputSchema: map[string]interface{}{
			"type": "array",
			"items": map[string]interface{}{
				"type": "object",
				"properties": map[string]interface{}{
					"itemNo":         map[string]string{"type": "string"},
					"itemType":       map[string]string{"type": "string"},
					"classUnitCode":  map[string]string{"type": "string"},
					"classUnitType":  map[string]string{"type": "string"},
				},
				"required": []string{"itemNo", "itemType", "classUnitCode", "classUnitType"},
			},
		},
	},
	{
		Name:        "rpim_get_updated_items",
		Description: "Get updated items from rpim-api-service",
		InputSchema: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"lastRunTime":    map[string]string{"type": "string"},
				"currentRunTime": map[string]string{"type": "string"},
			},
			"required": []string{"lastRunTime", "currentRunTime"},
		},
	},
	{
		Name:        "rpim_get_classified_items",
		Description: "Get all commercially classified items from rpim-api-service",
		InputSchema: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"classUnitCode":       map[string]string{"type": "string"},
				"classificationType":  map[string]string{"type": "string"},
			},
			"required": []string{"classUnitCode", "classificationType"},
		},
	},
	{
		Name:        "rpim_get_item_attributes",
		Description: "Get item attributes from rpim-api-service",
		InputSchema: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"itemType": map[string]string{"type": "string"},
				"itemNo":   map[string]string{"type": "string"},
			},
			"required": []string{"itemType", "itemNo"},
		},
	},
	{
		Name:        "rpim_get_items_by_class_unit",
		Description: "Get items by class unit from rpim-api-service",
		InputSchema: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"classUnitCode": map[string]string{"type": "string"},
			},
			"required": []string{"classUnitCode"},
		},
	},
}
