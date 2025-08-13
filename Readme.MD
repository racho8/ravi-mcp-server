# MCP (Model Context Protocol) Server

A Go-based MCP server that implements the **MCP JSON-RPC 2.0 Protocol** to provide a unified interface for interacting with microservices using standardized tool calls. This server is designed to run on Google Cloud Run and interfaces with other microservices while being compatible with Copilot agents.

## Overview

The MCP Server acts as an intermediary layer that follows the official MCP JSON-RPC 2.0 specification, enabling seamless integration with AI agents like GitHub Copilot. It standardizes communication with underlying microservices through a tool-based interface where each tool represents a specific microservice operation.

## MCP Protocol Compliance

This implementation follows the **MCP JSON-RPC 2.0 Protocol** specification with the following required methods:

### 1. `initialize`
Initial handshake between client and server to establish capabilities and protocol version.

### 2. `tools/list` 
Returns all available tools that the server can execute.

### 3. `tools/call`
Executes a specific tool with provided arguments.

## Features

- ✅ **Full MCP JSON-RPC 2.0 Protocol compliance**
- ✅ **Standardized error handling with JSON-RPC error codes**
- ✅ **Tool-based API interface**
- ✅ **Integration with Google Cloud Run**
- ✅ **Authentication support with service account tokens**
- ✅ **Configurable microservice endpoints**
- ✅ **Automatic token forwarding to microservices**
- ✅ **Compatible with GitHub Copilot and other MCP clients**

## Available Tools

The server provides the following tools that map to microservice endpoints:

| Tool Name | Description | Required Parameters |
|-----------|-------------|-------------------|
| `welcome_message` | Get welcome message from the product service | None |
| `health_check` | Check product service health | None |
| `create_product` | Create a new product | `name`, `category`, `price` |
| `get_product` | Retrieve a product by ID | `id` |
| `update_product` | Update an existing product | `id` (+ optional fields) |
| `delete_product` | Delete a product | `id` |
| `list_products` | Get all products | None |

## API Endpoints

- `POST /mcp`: Main MCP JSON-RPC 2.0 endpoint

## Protocol Usage Examples

### 1. Initialize Connection
```bash
curl -X POST https://your-server/mcp \
  -H "Content-Type: application/json" \
  -d '{
    "jsonrpc": "2.0",
    "id": 1,
    "method": "initialize",
    "params": {
      "protocolVersion": "2024-11-05",
      "capabilities": {},
      "clientInfo": {
        "name": "your-client",
        "version": "1.0.0"
      }
    }
  }'
```

### 2. List Available Tools
```bash
curl -X POST https://your-server/mcp \
  -H "Content-Type: application/json" \
  -d '{
    "jsonrpc": "2.0",
    "id": 2,
    "method": "tools/list"
  }'
```

### 3. Call a Tool
```bash
curl -X POST https://your-server/mcp \
  -H "Content-Type: application/json" \
  -d '{
    "jsonrpc": "2.0",
    "id": 3,
    "method": "tools/call",
    "params": {
      "name": "create_product",
      "arguments": {
        "name": "MacBook Pro",
        "category": "Electronics",
        "price": 2499
      }
    }
  }'
```

## Environment Variables

- `MICROSERVICE_URL`: URL of the target microservice
- `PORT`: Server port (defaults to 8080)

## Testing Your MCP Server

Use the provided test file `test_mcp_requests.json` for comprehensive testing examples.

### Local Testing
```bash
# Set environment variables
export MICROSERVICE_URL="https://product-service-256110662801.europe-west3.run.app"
export PORT="8080"

# Run the server
go run main.go

# Test initialize in another terminal
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

## Deployment

The application is configured for deployment to Google Cloud Run using GitHub Actions. The deployment workflow:
1. Builds a Docker container
2. Pushes to Google Artifact Registry  
3. Deploys to Cloud Run

### Required Secrets

- `GCP_SA_KEY`: Google Cloud service account key
- `GCP_PROJECT_ID`: Google Cloud project ID
- `MICROSERVICE_URL`: Target microservice URL

## GitHub Copilot Integration

This MCP server is fully compatible with GitHub Copilot and other MCP clients. Once deployed, you can:

1. **Configure Copilot**: Add your MCP server URL to Copilot's MCP server configuration
2. **Use Tools**: Copilot will automatically discover and use your tools through the standard MCP protocol
3. **Natural Language**: Ask Copilot to perform operations like "create a new product" or "list all products"

### Example Copilot Usage
```
User: "Create a new product called 'iPhone 15' in the Electronics category for $999"
Copilot: Uses the create_product tool via MCP protocol to call your microservice
```

## Key Improvements Made

### 1. ✅ **Proper JSON-RPC 2.0 Protocol**
- Added correct `jsonrpc: "2.0"` field
- Implemented proper `id` handling  
- Added standard error codes (-32700, -32600, -32601, -32602, -32603)

### 2. ✅ **Standard MCP Methods**
- `initialize`: Proper handshake with capabilities negotiation
- `tools/list`: Returns tool schemas with `inputSchema` 
- `tools/call`: Executes tools with proper argument handling

### 3. ✅ **Enhanced Error Handling**
- JSON-RPC compliant error responses
- Proper error codes and messages
- Validation of required parameters

### 4. ✅ **Improved Tool Response Format**
- Returns structured `ToolCallResult` with `content` array
- Supports both text and structured data responses
- Compatible with MCP client expectations

## Utility Commands

### Update Cloud Run Environment Variables
```bash
gcloud run services update ravi-mcp-server \
    --region=europe-west3 \
    --set-env-vars "MICROSERVICE_URL=https://product-service-256110662801.europe-west3.run.app"
```

### Check Current Environment Variables
```bash
gcloud run services describe ravi-mcp-server \
  --region=europe-west3 \
  --format="value(spec.template.spec.containers[0].env)"
```

### Add Public Access (if needed)
```bash
gcloud run services add-iam-policy-binding ravi-mcp-server \
  --region=europe-west3 \
  --member="domain:your-organization-domain.com" \
  --role="roles/run.invoker"
```

### Test Production Server
```bash
# Test initialize
curl -X POST https://ravi-mcp-server-256110662801.europe-west3.run.app/mcp \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $(gcloud auth print-identity-token)" \
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

# Test tool call
curl -X POST https://ravi-mcp-server-256110662801.europe-west3.run.app/mcp \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $(gcloud auth print-identity-token)" \
  -d '{
    "jsonrpc": "2.0",
    "id": 3,
    "method": "tools/call",
    "params": {
      "name": "create_product",
      "arguments": {
        "name": "Apple iPad Pro",
        "category": "Electronics",
        "price": 1299
      }
    }
  }'
```

## Troubleshooting

### Common Issues

1. **Method not found**: Ensure you're using exact method names: `initialize`, `tools/list`, `tools/call`
2. **Invalid JSON-RPC**: Always include `"jsonrpc": "2.0"` and unique `id` field
3. **Authentication errors**: Verify your GCP service account has proper permissions
4. **Microservice connectivity**: Check `MICROSERVICE_URL` environment variable

### Debug Logs
The server provides detailed logging for:
- Incoming JSON-RPC requests
- Tool execution details  
- Microservice call attempts
- Error conditions
