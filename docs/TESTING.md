# Quick Start Testing Guide

## For Postman Users

1. **Import Collection**: Import `postman_collection.json` into Postman
2. **Set Variables**: 
   - `base_url` = `http://localhost:8080` (for local testing)
   - `production_url` = `https://ravi-mcp-server-256110662801.europe-west3.run.app`
3. **Start Server**: Run your local server first
4. **Run Tests**: Execute requests in order

## For Command Line Users

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

## Expected Responses

### Initialize Response
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

### Tools List Response
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

### Tool Call Response (JSON Data)
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

### Tool Call Response (Text Data)
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

### Error Response
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

## Testing Workflow

1. **Initialize** - Test connection
2. **Tools/List** - Verify all tools are available  
3. **Health Check** - Test basic connectivity to microservice
4. **List Products** - Test GET operations
5. **Create Product** - Test POST operations (note the returned ID)
6. **Get Product** - Test GET by ID (use ID from step 5)
7. **Update Product** - Test PUT operations (use ID from step 5)
8. **Delete Product** - Test DELETE operations (use ID from step 5)
9. **Error Tests** - Verify proper error handling

## Files Created for Testing

- `postman_collection.json` - Import into Postman
- `test_commands.sh` - Manual curl commands (display only)
- `run_tests.sh` - Automated test execution
- `validate_mcp.sh` - Protocol validation script
