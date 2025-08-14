# Postman Import Instructions

# Postman Import Instructions

<details>
<summary><strong>✅ Option 1: Import the Fixed Collection (Recommended)</strong></summary>

Use the file `postman_collection_v2.json` in this directory which follows the proper Postman Collection v2.1 format.

### Steps:
1. Open Postman
2. Click "Import" button (top left)
3. Select "File" tab
4. Choose `tests/postman_collection_v2.json`
5. Click "Import"
</details>

<details>
<summary><strong>🔧 Option 2: Manual Setup (If import still fails)</strong></summary>

### 1. Create New Collection
- Click "New" → "Collection"
- Name: "Ravi MCP Server Test Collection"

### 2. Add Variables
- Go to collection settings → Variables tab
- Add variable: `base_url` = `http://localhost:8080`
- Add variable: `production_url` = `https://ravi-mcp-server-256110662801.europe-west3.run.app`

### 3. Add Requests Manually
Copy-paste these requests one by one:

#### Request 1: Initialize Connection
```
POST {{base_url}}/mcp
Content-Type: application/json

{
  "jsonrpc": "2.0",
  "id": 1,
  "method": "initialize",
  "params": {
    "protocolVersion": "2024-11-05",
    "capabilities": {},
    "clientInfo": {
      "name": "postman-client",
      "version": "1.0.0"
    }
  }
}
```

#### Request 2: List Tools
```
POST {{base_url}}/mcp
Content-Type: application/json

{
  "jsonrpc": "2.0",
  "id": 2,
  "method": "tools/list"
}
```

#### Request 3: Health Check
```
POST {{base_url}}/mcp
Content-Type: application/json

{
  "jsonrpc": "2.0",
  "id": 3,
  "method": "tools/call",
  "params": {
    "name": "health_check"
  }
}
```

#### Request 4: Create Product
```
POST {{base_url}}/mcp
Content-Type: application/json

{
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
}
```

#### Request 5: List Products
```
POST {{base_url}}/mcp
Content-Type: application/json

{
  "jsonrpc": "2.0",
  "id": 5,
  "method": "tools/call",
  "params": {
    "name": "list_products"
  }
}
```
</details>

<details>
<summary><strong>💻 Option 3: Use cURL Commands</strong></summary>

If Postman import continues to fail, use the automated test script:

```bash
# Start server
export MICROSERVICE_URL="https://product-service-256110662801.europe-west3.run.app"
go run main.go

# Run tests from project root
cd ..
./tests/run_tests.sh
```
</details>

## Troubleshooting Postman Import

<details>
<summary><strong>🔍 Common Issues and Solutions</strong></summary>

If you're still having issues:

1. **Check Postman Version**: Make sure you're using a recent version of Postman
2. **Try Different Import Methods**: 
   - Import via "Link" instead of "File"
   - Copy-paste the JSON content directly
3. **Validate JSON**: Use a JSON validator to ensure the file is valid
4. **Manual Setup**: Create requests manually as shown in Option 2
</details>

## Files in This Directory

<details>
<summary><strong>📁 Available Test Files</strong></summary>

- `postman_collection_v2.json` - Proper Postman Collection v2.1 format
- `run_tests.sh` - Automated test runner
- `test_commands.sh` - Manual cURL commands
- `test_mcp_requests.json` - JSON test examples
- `validate_mcp.sh` - MCP protocol validation
</details>
