package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const productServiceBaseURL = "https://product-service-256110662801.europe-west3.run.app"
const rpimApiBaseURL = "https://rpim-api-service-256110662801.europe-west3.run.app"


// Business logic functions for MCP server
func executeToolCall(toolName string, params map[string]interface{}) (interface{}, error) {
	switch toolName {
	case "welcome_message":
		return map[string]string{"message": "Welcome to the MCP Product Service!"}, nil
	case "health_check":
		return map[string]string{"status": "ok"}, nil
	case "create_product":
		return createProduct(params)
	case "get_product":
		return getProduct(params)
	case "update_product":
		return updateProduct(params)
	case "delete_product":
		return deleteProduct(params)
	case "list_products":
		return listProducts(params)
	case "create_multiple_products":
		return createMultipleProducts(params)
	case "update_products":
		return updateProducts(params)
	case "delete_products":
		return deleteProducts(params)
	case "rpim_health_check":
		return rpimHealthCheck(params)
	case "rpim_get_child_items":
		return rpimGetChildItems(params)
	case "rpim_get_item_keys":
		return rpimGetItemKeys(params)
	case "rpim_get_class_units":
		return rpimGetClassUnits(params)
	case "rpim_get_item_details":
		return rpimGetItemDetails(params)
	case "rpim_get_local_items":
		return rpimGetLocalItems(params)
	case "rpim_get_updated_items":
		return rpimGetUpdatedItems(params)
	case "rpim_get_classified_items":
		return rpimGetClassifiedItems(params)
	case "rpim_get_item_attributes":
		return rpimGetItemAttributes(params)
	case "rpim_get_items_by_class_unit":
		return rpimGetItemsByClassUnit(params)
	default:
		return nil, fmt.Errorf("unknown tool: %s", toolName)
	}
}

// Example business logic implementations
// rpim-api-service business logic
func rpimHealthCheck(params map[string]interface{}) (interface{}, error) {
	url := rpimApiBaseURL + "/health"
	return invokeMicroservice("GET", url, nil)
}

func rpimGetChildItems(params map[string]interface{}) (interface{}, error) {
	url := rpimApiBaseURL + "/api/v1/item-content"
	return invokeMicroservice("POST", url, params)
}

func rpimGetItemKeys(params map[string]interface{}) (interface{}, error) {
	url := rpimApiBaseURL + "/items/keys"
	return invokeMicroservice("POST", url, params)
}

func rpimGetClassUnits(params map[string]interface{}) (interface{}, error) {
	url := rpimApiBaseURL + "/class-units"
	return invokeMicroservice("GET", url, nil)
}

func rpimGetItemDetails(params map[string]interface{}) (interface{}, error) {
	url := rpimApiBaseURL + "/items/details"
	return invokeMicroservice("POST", url, params)
}

func rpimGetLocalItems(params map[string]interface{}) (interface{}, error) {
	url := rpimApiBaseURL + "/items/local"
	return invokeMicroservice("POST", url, params)
}

func rpimGetUpdatedItems(params map[string]interface{}) (interface{}, error) {
	url := rpimApiBaseURL + "/items/updated"
	return invokeMicroservice("POST", url, params)
}

func rpimGetClassifiedItems(params map[string]interface{}) (interface{}, error) {
	url := rpimApiBaseURL + "/items/classified"
	return invokeMicroservice("POST", url, params)
}

func rpimGetItemAttributes(params map[string]interface{}) (interface{}, error) {
	url := rpimApiBaseURL + "/items/attributes"
	return invokeMicroservice("POST", url, params)
}

func rpimGetItemsByClassUnit(params map[string]interface{}) (interface{}, error) {
	url := rpimApiBaseURL + "/items/by-class-unit"
	return invokeMicroservice("POST", url, params)
}

func deleteProducts(params map[string]interface{}) (interface{}, error) {
	url := productServiceBaseURL + "/products/delete"
	return invokeMicroservice("POST", url, params)
}

func createProduct(params map[string]interface{}) (interface{}, error) {
	url := productServiceBaseURL + "/products"
	return invokeMicroservice("POST", url, params)
}

func getProduct(params map[string]interface{}) (interface{}, error) {
	id, ok := params["id"].(string)
	if !ok || id == "" {
		return nil, fmt.Errorf("missing or invalid product id")
	}
	url := fmt.Sprintf(productServiceBaseURL+"/products/%s", id)
	return invokeMicroservice("GET", url, nil)
}
func updateProduct(params map[string]interface{}) (interface{}, error) {
	id, ok := params["id"].(string)
	if !ok || id == "" {
		return nil, fmt.Errorf("missing or invalid product id")
	}
	url := fmt.Sprintf(productServiceBaseURL+"/products/%s", id)
	return invokeMicroservice("PUT", url, params)
}

func deleteProduct(params map[string]interface{}) (interface{}, error) {
	id, ok := params["id"].(string)
	if !ok || id == "" {
		return nil, fmt.Errorf("missing or invalid product id")
	}
	url := fmt.Sprintf(productServiceBaseURL+"/products/%s", id)
	return invokeMicroservice("DELETE", url, nil)
}

func listProducts(params map[string]interface{}) (interface{}, error) {
	url := productServiceBaseURL + "/products"
	return invokeMicroservice("GET", url, nil)
}

func createMultipleProducts(params map[string]interface{}) (interface{}, error) {
	url := productServiceBaseURL + "/products/create-multiple"
	return invokeMicroservice("POST", url, params)
}

func updateProducts(params map[string]interface{}) (interface{}, error) {
	url := productServiceBaseURL + "/products/update"
	return invokeMicroservice("POST", url, params)
}


// Helper to make HTTP requests to microservice and parse response
func invokeMicroservice(method, url string, params map[string]interface{}) (interface{}, error) {
	var reqBody *bytes.Buffer
	if params != nil {
		body, err := json.Marshal(params)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal params: %v", err)
		}
		reqBody = bytes.NewBuffer(body)
	} else {
		reqBody = &bytes.Buffer{}
	}

	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to call microservice: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read microservice response: %v", err)
	}

	var result interface{}
	if err := json.Unmarshal(respBody, &result); err != nil {
		// If not JSON, return raw text
		return map[string]string{"response": string(respBody)}, nil
	}
	return result, nil
}
