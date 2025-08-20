#!/bin/bash
# validate_mcp.sh - Validates MCP server protocol and endpoints
# Usage: ./validate_mcp.sh <server_url>

SERVER_URL="${1:-http://localhost:8080}"

function validate() {
  local method="$1"
  local payload="$2"
  echo "Validating $method..."
  response=$(curl -s -X POST "$SERVER_URL/mcp" -H "Content-Type: application/json" -d "$payload")
  echo "$response" | jq . 2>/dev/null || echo "$response"
  echo "---"
}

# Validate initialize
validate "initialize" '{"jsonrpc":"2.0","id":1,"method":"initialize","params":{"protocolVersion":"2024-11-05","capabilities":{},"clientInfo":{"name":"validator","version":"1.0.0"}}}'

# Validate tools/list
validate "tools/list" '{"jsonrpc":"2.0","id":2,"method":"tools/list"}'

# Validate health_check
validate "health_check" '{"jsonrpc":"2.0","id":3,"method":"tools/call","params":{"name":"health_check"}}'

# Validate list_products
validate "list_products" '{"jsonrpc":"2.0","id":4,"method":"tools/call","params":{"name":"list_products"}}'

# Validate error handling (invalid method)
validate "invalid_method" '{"jsonrpc":"2.0","id":5,"method":"invalid_method"}'

# Validate error handling (invalid tool)
validate "invalid_tool" '{"jsonrpc":"2.0","id":6,"method":"tools/call","params":{"name":"invalid_tool"}}'

# Validate error handling (missing params)
validate "missing_params" '{"jsonrpc":"2.0","id":7,"method":"tools/call","params":{"name":"get_product"}}'

# End of script
