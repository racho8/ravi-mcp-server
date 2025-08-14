# Quick Start Testing Guide

<details>
<summary><strong>📮 For Postman Users</strong></summary>

1. **Import Collection**: Import `postman_collection.json` into Postman
2. **Set Variables**: 
   - `base_url` = `http://localhost:8080` (for local testing)
   - `production_url` = `https://ravi-mcp-server-256110662801.europe-west3.run.app`
3. **Start Server**: Run your local server first
4. **Run Tests**: Execute requests in order
</details>

<details>
<summary><strong>💻 For Command Line Users</strong></summary>

### Quick Test (Automated)
```bash
# Start server in one terminal
export MICROSERVICE_URL="https://product-service-256110662801.europe-west3.run.app"
go run main.go

# Run automated tests in another terminal
./run_tests.sh
```

### Manual Testing
```bash
# View all test commands
cat test_commands.sh

# Or run individual commands like:
curl -X POST http://localhost:8080/mcp \
  -H "Content-Type: application/json" \
  -d '{
    "jsonrpc": "2.0",
    "id": 1,
    "method": "initialize",
    "params": {
      "protocolVersion": "2024-11-05",
      "capabilities": {},
      "clientInfo": {"name": "test-client", "version": "1.0.0"}
    }
  }'
```
</details>

## Expected Responses

<details>
<summary><strong>🔧 Initialize Response</strong></summary>

```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "result": {
    "protocolVersion": "2024-11-05",
    "capabilities": {
      "tools": {}
    },
    "serverInfo": {
      "name": "ravi-mcp-server",
      "version": "1.0.0"
    }
  }
}
```
</details>

<details>
<summary><strong>🛠️ Tools List Response</strong></summary>

```json
{
  "jsonrpc": "2.0",
  "id": 2,
  "result": {
    "tools": [
      {
        "name": "welcome_message",
        "description": "Get the welcome message from the product service",
        "inputSchema": {
          "type": "object",
          "properties": {}
        }
      },
      // ... more tools
    ]
  }
}
```
</details>

<details>
<summary><strong>📊 Tool Call Response (JSON Data)</strong></summary>

```json
{
  "jsonrpc": "2.0",
  "id": 5,
  "result": {
    "content": [
      {
        "type": "text",
        "data": [
          {
            "id": "product-123",
            "name": "MacBook Pro",
            "category": "Electronics",
            "price": 2499
          }
        ]
      }
    ]
  }
}
```
</details>

<details>
<summary><strong>📝 Tool Call Response (Text Data)</strong></summary>

```json
{
  "jsonrpc": "2.0",
  "id": 3,
  "result": {
    "content": [
      {
        "type": "text",
        "text": "Service is healthy"
      }
    ]
  }
}
```
</details>

<details>
<summary><strong>❌ Error Response</strong></summary>

```json
{
  "jsonrpc": "2.0",
  "id": 11,
  "error": {
    "code": -32601,
    "message": "Method not found",
    "data": "Unknown method: invalid_method"
  }
}
```
</details>

## Testing Workflow

<details>
<summary><strong>📋 Step-by-Step Testing Process</strong></summary>

1. **Initialize** - Test connection
2. **Tools/List** - Verify all tools are available  
3. **Health Check** - Test basic connectivity to microservice
4. **List Products** - Test GET operations
5. **Create Product** - Test POST operations (note the returned ID)
6. **Get Product** - Test GET by ID (use ID from step 5)
7. **Update Product** - Test PUT operations (use ID from step 5)
8. **Delete Product** - Test DELETE operations (use ID from step 5)
9. **Error Tests** - Verify proper error handling
</details>

## Files Created for Testing

<details>
<summary><strong>📁 Testing Resources</strong></summary>

- `postman_collection.json` - Import into Postman
- `test_commands.sh` - Manual curl commands (display only)
- `run_tests.sh` - Automated test execution
- `validate_mcp.sh` - Protocol validation script
</details>
