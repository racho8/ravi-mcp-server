#!/bin/bash

# MCP Server Protocol Validation Script
# This script validates that your MCP server correctly implements the JSON-RPC 2.0 protocol

echo "=== MCP Server Protocol Validation ==="
echo

# Server URL (can be overridden with environment variable)
if [ -z "$MCP_SERVER_URL" ]; then
    SERVER_URL="http://localhost:8080/mcp"
    echo "Using default server URL: $SERVER_URL"
    echo "Tip: Set MCP_SERVER_URL environment variable to test different servers"
else
    SERVER_URL="$MCP_SERVER_URL"
    echo "Using custom server URL: $SERVER_URL"
fi

# Check if we need authentication headers
AUTH_HEADER=""
if [[ "$SERVER_URL" == *"run.app"* ]]; then
    echo "Production server detected, adding authentication..."
    if command -v gcloud &> /dev/null; then
        TOKEN=$(gcloud auth print-access-token 2>/dev/null)
        if [ $? -eq 0 ] && [ ! -z "$TOKEN" ]; then
            AUTH_HEADER="-H \"Authorization: Bearer $TOKEN\""
            echo "✅ Authentication token added"
        else
            echo "⚠️  Warning: Could not get GCP token. You may see auth errors."
        fi
    else
        echo "⚠️  Warning: gcloud not found. You may see auth errors for production server."
    fi
fi

echo "Testing server at: $SERVER_URL"
echo

# Test 1: Initialize
echo "1. Testing 'initialize' method..."
eval curl -s -X POST $SERVER_URL \
  -H "Content-Type: application/json" \
  $AUTH_HEADER \
  -d '{
    "jsonrpc": "2.0",
    "id": 1,
    "method": "initialize",
    "params": {
      "protocolVersion": "2024-11-05",
      "capabilities": {},
      "clientInfo": {
        "name": "test-client",
        "version": "1.0.0"
      }
    }
  }' | jq .

echo -e "\n---\n"

# Test 2: Tools List
echo "2. Testing 'tools/list' method..."
eval curl -s -X POST $SERVER_URL \
  -H "Content-Type: application/json" \
  $AUTH_HEADER \
  -d '{
    "jsonrpc": "2.0",
    "id": 2,
    "method": "tools/list"
  }' | jq .

echo -e "\n---\n"

# Test 3: Tool Call
echo "3. Testing 'tools/call' method (list_products)..."
eval curl -s -X POST $SERVER_URL \
  -H "Content-Type: application/json" \
  $AUTH_HEADER \
  -d '{
    "jsonrpc": "2.0",
    "id": 3,
    "method": "tools/call",
    "params": {
      "name": "list_products"
    }
  }' | jq .

echo -e "\n---\n"

# Test 4: Invalid Method
echo "4. Testing invalid method (should return error)..."
eval curl -s -X POST $SERVER_URL \
  -H "Content-Type: application/json" \
  $AUTH_HEADER \
  -d '{
    "jsonrpc": "2.0",
    "id": 4,
    "method": "invalid_method"
  }' | jq .

echo -e "\n=== Validation Complete ==="
echo
echo "💡 Usage Tips:"
echo "  Local testing:     ./validate_mcp.sh"
echo "  Production testing: MCP_SERVER_URL=https://your-server.run.app/mcp ./validate_mcp.sh"
