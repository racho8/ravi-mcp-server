#!/bin/bash

# MCP Server Protocol Validation Script
# This script validates that your MCP server correctly implements the JSON-RPC 2.0 protocol

echo "=== MCP Server Protocol Validation ==="
echo

# Server URL (change this to your actual server URL)
SERVER_URL="http://localhost:8080/mcp"

echo "Testing server at: $SERVER_URL"
echo

# Test 1: Initialize
echo "1. Testing 'initialize' method..."
curl -s -X POST $SERVER_URL \
  -H "Content-Type: application/json" \
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
curl -s -X POST $SERVER_URL \
  -H "Content-Type: application/json" \
  -d '{
    "jsonrpc": "2.0",
    "id": 2,
    "method": "tools/list"
  }' | jq .

echo -e "\n---\n"

# Test 3: Tool Call
echo "3. Testing 'tools/call' method (list_products)..."
curl -s -X POST $SERVER_URL \
  -H "Content-Type: application/json" \
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
curl -s -X POST $SERVER_URL \
  -H "Content-Type: application/json" \
  -d '{
    "jsonrpc": "2.0",
    "id": 4,
    "method": "invalid_method"
  }' | jq .

echo -e "\n=== Validation Complete ==="
