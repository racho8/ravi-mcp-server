// Package main - business.go
//
// This file contains all business logic for executing MCP tools and communicating with the
// backend product service microservice.
//
// Key Responsibilities:
//   - Route tool calls to appropriate business logic functions
//   - Validate tool parameters and inputs
//   - Make HTTP requests to backend product service
//   - Transform and return results to the MCP handler
//   - Handle errors from the backend service
//
// Tool Execution Flow:
//   1. executeToolCall() receives tool name and parameters
//   2. Routes to specific tool function based on tool name
//   3. Tool function validates parameters and constructs HTTP request
//   4. invokeMicroservice() makes the actual HTTP call
//   5. Response is parsed and returned to handler
//
// Tool Functions:
//
//   Service Tools:
//     - welcome_message: Returns static welcome message
//     - health_check: Returns static health status
//
//   Single Product Operations:
//     - createProduct: POST /products
//     - getProduct: GET /products/{id}
//     - updateProduct: PUT /products/{id}
//     - deleteProduct: DELETE /products/{id}
//     - listProducts: GET /products
//
//   Batch Operations:
//     - createMultipleProducts: POST /products/create-multiple
//     - updateProducts: POST /products/update
//     - deleteProducts: POST /products/delete
//
//   Query Operations:
//     - getProductsByCategory: GET /products/category/{category}
//     - getProductsBySegment: GET /products/segment/{segment}
//     - getProductByName: GET /products/{name}
//
// Helper Functions:
//   - invokeMicroservice: Generic HTTP client for backend service calls
//
// Backend Service:
//   - Base URL: https://product-service-256110662801.europe-west3.run.app
//   - TODO: Read from MICROSERVICE_URL environment variable
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// TODO: read it from environment variable MICROSERVICE_URL
const productServiceBaseURL = "https://product-service-256110662801.europe-west3.run.app"

// business logic functions for MCP server
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
	case "get_products_by_category":
		return getProductsByCategory(params)
	case "get_products_by_segment":
		return getProductsBySegment(params)
	case "get_product_by_name":
		return getProductByName(params)
	case "list_products":
		return listProducts(params)
	case "create_multiple_products":
		return createMultipleProducts(params)
	case "update_product":
		return updateProduct(params)
	case "update_products":
		return updateProducts(params)
	case "delete_product":
		return deleteProduct(params)
	case "delete_products":
		return deleteProducts(params)
	}
	return nil, fmt.Errorf("unknown tool: %s", toolName)
}

// Returns all products in the store
func listProducts(params map[string]interface{}) (interface{}, error) {
	url := productServiceBaseURL + "/products"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("product service returned status %d", resp.StatusCode)
	}
	var products []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&products); err != nil {
		return nil, err
	}
	return products, nil
}

// Returns all products matching a given category
func getProductsByCategory(params map[string]interface{}) (interface{}, error) {
	category, ok := params["category"].(string)
	if !ok || category == "" {
		return nil, fmt.Errorf("missing or invalid 'category' argument")
	}
	url := productServiceBaseURL + "/products/category/" + category
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("product service returned status %d", resp.StatusCode)
	}
	var products []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&products); err != nil {
		return nil, err
	}
	return products, nil
}

// Returns all products matching a given segment
func getProductsBySegment(params map[string]interface{}) (interface{}, error) {
	segment, ok := params["segment"].(string)
	if !ok || segment == "" {
		return nil, fmt.Errorf("missing or invalid 'segment' argument")
	}
	url := productServiceBaseURL + "/products/segment/" + segment
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("product service returned status %d", resp.StatusCode)
	}
	var products []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&products); err != nil {
		return nil, err
	}
	return products, nil
}

// Returns all products matching a given name
func getProductByName(params map[string]interface{}) (interface{}, error) {
	name, ok := params["name"].(string)
	if !ok || name == "" {
		return nil, fmt.Errorf("missing or invalid 'name' argument")
	}
	url := productServiceBaseURL + "/products/" + name
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("product service returned status %d", resp.StatusCode)
	}
	var products []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&products); err != nil {
		return nil, err
	}
	return products, nil
}

// business logic implementations
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
	// Only include fields that are present in params
	updateFields := make(map[string]interface{})
	updateFields["id"] = id
	if name, ok := params["name"].(string); ok && name != "" {
		updateFields["name"] = name
	}
	if price, ok := params["price"].(float64); ok {
		updateFields["price"] = price
	}
	if category, ok := params["category"].(string); ok && category != "" {
		updateFields["category"] = category
	}
	url := fmt.Sprintf(productServiceBaseURL+"/products/%s", id)
	return invokeMicroservice("PUT", url, updateFields)
}

func deleteProduct(params map[string]interface{}) (interface{}, error) {
	id, ok := params["id"].(string)
	if !ok || id == "" {
		return nil, fmt.Errorf("missing or invalid product id")
	}
	url := fmt.Sprintf(productServiceBaseURL+"/products/%s", id)
	return invokeMicroservice("DELETE", url, nil)
}

// TODO: add pagination support and use params to filter results

func createMultipleProducts(params map[string]interface{}) (interface{}, error) {
	url := productServiceBaseURL + "/products/create-multiple"
	return invokeMicroservice("POST", url, params)
}

func updateProducts(params map[string]interface{}) (interface{}, error) {
	url := productServiceBaseURL + "/products/update"
	return invokeMicroservice("POST", url, params)
}

// helper to make HTTP requests to microservice and parse response
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
