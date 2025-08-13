# MCP Server Local Testing - cURL Commands
# Run these commands after starting your local server with: go run main.go
# Make sure MICROSERVICE_URL environment variable is set

echo "=== MCP Server Local Testing Guide ==="
echo "Server URL: http://localhost:8080/mcp"
echo

# 1. Initialize Connection
echo "1. Initialize Connection:"
echo "========================"
cat << 'EOF'
curl -X POST http://localhost:8080/mcp \
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
  }'
EOF
echo -e "\n"

# 2. List Available Tools
echo "2. List Available Tools:"
echo "========================"
cat << 'EOF'
curl -X POST http://localhost:8080/mcp \
  -H "Content-Type: application/json" \
  -d '{
    "jsonrpc": "2.0",
    "id": 2,
    "method": "tools/list"
  }'
EOF
echo -e "\n"

# 3. Health Check
echo "3. Health Check:"
echo "================"
cat << 'EOF'
curl -X POST http://localhost:8080/mcp \
  -H "Content-Type: application/json" \
  -d '{
    "jsonrpc": "2.0",
    "id": 3,
    "method": "tools/call",
    "params": {
      "name": "health_check"
    }
  }'
EOF
echo -e "\n"

# 4. Welcome Message
echo "4. Welcome Message:"
echo "==================="
cat << 'EOF'
curl -X POST http://localhost:8080/mcp \
  -H "Content-Type: application/json" \
  -d '{
    "jsonrpc": "2.0",
    "id": 4,
    "method": "tools/call",
    "params": {
      "name": "welcome_message"
    }
  }'
EOF
echo -e "\n"

# 5. List All Products
echo "5. List All Products:"
echo "===================="
cat << 'EOF'
curl -X POST http://localhost:8080/mcp \
  -H "Content-Type: application/json" \
  -d '{
    "jsonrpc": "2.0",
    "id": 5,
    "method": "tools/call",
    "params": {
      "name": "list_products"
    }
  }'
EOF
echo -e "\n"

# 6. Create Product - MacBook Pro
echo "6. Create Product - MacBook Pro:"
echo "================================"
cat << 'EOF'
curl -X POST http://localhost:8080/mcp \
  -H "Content-Type: application/json" \
  -d '{
    "jsonrpc": "2.0",
    "id": 6,
    "method": "tools/call",
    "params": {
      "name": "create_product",
      "arguments": {
        "name": "MacBook Pro 16-inch",
        "category": "Electronics",
        "segment": "Laptops",
        "price": 2499
      }
    }
  }'
EOF
echo -e "\n"

# 7. Create Product - iPhone
echo "7. Create Product - iPhone:"
echo "==========================="
cat << 'EOF'
curl -X POST http://localhost:8080/mcp \
  -H "Content-Type: application/json" \
  -d '{
    "jsonrpc": "2.0",
    "id": 7,
    "method": "tools/call",
    "params": {
      "name": "create_product",
      "arguments": {
        "name": "iPhone 15 Pro",
        "category": "Electronics",
        "segment": "Smartphones",
        "price": 999
      }
    }
  }'
EOF
echo -e "\n"

# 8. Get Product by ID (replace product-123 with actual ID from create response)
echo "8. Get Product by ID:"
echo "===================="
cat << 'EOF'
curl -X POST http://localhost:8080/mcp \
  -H "Content-Type: application/json" \
  -d '{
    "jsonrpc": "2.0",
    "id": 8,
    "method": "tools/call",
    "params": {
      "name": "get_product",
      "arguments": {
        "id": "REPLACE_WITH_ACTUAL_PRODUCT_ID"
      }
    }
  }'
EOF
echo -e "\n"

# 9. Update Product (replace product-123 with actual ID)
echo "9. Update Product:"
echo "=================="
cat << 'EOF'
curl -X POST http://localhost:8080/mcp \
  -H "Content-Type: application/json" \
  -d '{
    "jsonrpc": "2.0",
    "id": 9,
    "method": "tools/call",
    "params": {
      "name": "update_product",
      "arguments": {
        "id": "REPLACE_WITH_ACTUAL_PRODUCT_ID",
        "name": "MacBook Pro 16-inch M3",
        "price": 2799
      }
    }
  }'
EOF
echo -e "\n"

# 10. Delete Product (replace product-123 with actual ID)
echo "10. Delete Product:"
echo "=================="
cat << 'EOF'
curl -X POST http://localhost:8080/mcp \
  -H "Content-Type: application/json" \
  -d '{
    "jsonrpc": "2.0",
    "id": 10,
    "method": "tools/call",
    "params": {
      "name": "delete_product",
      "arguments": {
        "id": "REPLACE_WITH_ACTUAL_PRODUCT_ID"
      }
    }
  }'
EOF
echo -e "\n"

# Error Tests
echo "=== ERROR TESTS ==="
echo

# 11. Error Test - Invalid Method
echo "11. Error Test - Invalid Method:"
echo "================================"
cat << 'EOF'
curl -X POST http://localhost:8080/mcp \
  -H "Content-Type: application/json" \
  -d '{
    "jsonrpc": "2.0",
    "id": 11,
    "method": "invalid_method"
  }'
EOF
echo -e "\n"

# 12. Error Test - Invalid Tool
echo "12. Error Test - Invalid Tool:"
echo "=============================="
cat << 'EOF'
curl -X POST http://localhost:8080/mcp \
  -H "Content-Type: application/json" \
  -d '{
    "jsonrpc": "2.0",
    "id": 12,
    "method": "tools/call",
    "params": {
      "name": "invalid_tool"
    }
  }'
EOF
echo -e "\n"

# 13. Error Test - Missing Required Parameter
echo "13. Error Test - Missing Required Parameter:"
echo "============================================"
cat << 'EOF'
curl -X POST http://localhost:8080/mcp \
  -H "Content-Type: application/json" \
  -d '{
    "jsonrpc": "2.0",
    "id": 13,
    "method": "tools/call",
    "params": {
      "name": "get_product"
    }
  }'
EOF

echo -e "\n=== Testing Complete ==="
