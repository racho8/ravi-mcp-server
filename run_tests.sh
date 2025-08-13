#!/bin/bash

# Executable test script for MCP server
# Usage: ./run_tests.sh

BASE_URL="http://localhost:8080"
SERVER_PID=""

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}=== MCP Server Test Runner ===${NC}"
echo

# Function to check if server is running
check_server() {
    curl -s "$BASE_URL/mcp" > /dev/null 2>&1
    return $?
}

# Function to run a test and display results
run_test() {
    local test_name="$1"
    local curl_command="$2"
    
    echo -e "${YELLOW}Testing: $test_name${NC}"
    echo "Command: $curl_command"
    echo "Response:"
    
    response=$(eval "$curl_command" 2>/dev/null)
    if [ $? -eq 0 ] && [ ! -z "$response" ]; then
        echo "$response" | jq . 2>/dev/null || echo "$response"
        echo -e "${GREEN}✓ Test passed${NC}"
    else
        echo -e "${RED}✗ Test failed or no response${NC}"
    fi
    echo "----------------------------------------"
    echo
}

# Check if server is running
echo "Checking if MCP server is running at $BASE_URL..."
if ! check_server; then
    echo -e "${RED}❌ Server is not running at $BASE_URL${NC}"
    echo "Please start the server first with:"
    echo "  export MICROSERVICE_URL=\"https://product-service-256110662801.europe-west3.run.app\""
    echo "  go run main.go"
    echo
    exit 1
fi

echo -e "${GREEN}✓ Server is running${NC}"
echo

# Test 1: Initialize
run_test "Initialize Connection" \
'curl -s -X POST '"$BASE_URL"'/mcp \
  -H "Content-Type: application/json" \
  -d "{
    \"jsonrpc\": \"2.0\",
    \"id\": 1,
    \"method\": \"initialize\",
    \"params\": {
      \"protocolVersion\": \"2024-11-05\",
      \"capabilities\": {},
      \"clientInfo\": {
        \"name\": \"test-client\",
        \"version\": \"1.0.0\"
      }
    }
  }"'

# Test 2: Tools List
run_test "List Available Tools" \
'curl -s -X POST '"$BASE_URL"'/mcp \
  -H "Content-Type: application/json" \
  -d "{
    \"jsonrpc\": \"2.0\",
    \"id\": 2,
    \"method\": \"tools/list\"
  }"'

# Test 3: Health Check
run_test "Health Check" \
'curl -s -X POST '"$BASE_URL"'/mcp \
  -H "Content-Type: application/json" \
  -d "{
    \"jsonrpc\": \"2.0\",
    \"id\": 3,
    \"method\": \"tools/call\",
    \"params\": {
      \"name\": \"health_check\"
    }
  }"'

# Test 4: Welcome Message
run_test "Welcome Message" \
'curl -s -X POST '"$BASE_URL"'/mcp \
  -H "Content-Type: application/json" \
  -d "{
    \"jsonrpc\": \"2.0\",
    \"id\": 4,
    \"method\": \"tools/call\",
    \"params\": {
      \"name\": \"welcome_message\"
    }
  }"'

# Test 5: List Products
run_test "List All Products" \
'curl -s -X POST '"$BASE_URL"'/mcp \
  -H "Content-Type: application/json" \
  -d "{
    \"jsonrpc\": \"2.0\",
    \"id\": 5,
    \"method\": \"tools/call\",
    \"params\": {
      \"name\": \"list_products\"
    }
  }"'

# Test 6: Create Product
run_test "Create Product - MacBook Pro" \
'curl -s -X POST '"$BASE_URL"'/mcp \
  -H "Content-Type: application/json" \
  -d "{
    \"jsonrpc\": \"2.0\",
    \"id\": 6,
    \"method\": \"tools/call\",
    \"params\": {
      \"name\": \"create_product\",
      \"arguments\": {
        \"name\": \"MacBook Pro 16-inch\",
        \"category\": \"Electronics\",
        \"segment\": \"Laptops\",
        \"price\": 2499
      }
    }
  }"'

# Error Tests
echo -e "${BLUE}=== Error Handling Tests ===${NC}"
echo

# Test 7: Invalid Method
run_test "Invalid Method (should return error)" \
'curl -s -X POST '"$BASE_URL"'/mcp \
  -H "Content-Type: application/json" \
  -d "{
    \"jsonrpc\": \"2.0\",
    \"id\": 11,
    \"method\": \"invalid_method\"
  }"'

# Test 8: Invalid Tool
run_test "Invalid Tool (should return error)" \
'curl -s -X POST '"$BASE_URL"'/mcp \
  -H "Content-Type: application/json" \
  -d "{
    \"jsonrpc\": \"2.0\",
    \"id\": 12,
    \"method\": \"tools/call\",
    \"params\": {
      \"name\": \"invalid_tool\"
    }
  }"'

# Test 9: Missing Required Parameter
run_test "Missing Required Parameter (should return error)" \
'curl -s -X POST '"$BASE_URL"'/mcp \
  -H "Content-Type: application/json" \
  -d "{
    \"jsonrpc\": \"2.0\",
    \"id\": 13,
    \"method\": \"tools/call\",
    \"params\": {
      \"name\": \"get_product\"
    }
  }"'

echo -e "${BLUE}=== All Tests Complete ===${NC}"
echo
echo "Note: For tests involving specific product IDs (get, update, delete),"
echo "first create a product and use the returned ID from the response."
