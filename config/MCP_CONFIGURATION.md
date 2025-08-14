# MCP Configuration Guide

This guide helps you configure your MCP server to work with GitHub Copilot and other MCP clients using natural language commands.

## 🎯 **Quick Configuration**

<details>
<summary><strong>🤖 GitHub Copilot Setup</strong></summary>

```bash
# Copy the appropriate config
cp config/mcp.json ~/.config/claude/claude_desktop_config.json

# Restart Claude Desktop and ask:
# "Show me all products in the store"
```
</details>

<details>
<summary><strong>🔗 Custom MCP Client Setup</strong></summary>

- **Server URL**: `https://ravi-mcp-server-256110662801.europe-west3.run.app/mcp`
- **Protocol**: MCP JSON-RPC 2.0
- **Authentication**: Google Cloud Bearer tokens
</details>

## 📋 **Configuration Files**

1. **`mcp.json`** - Full configuration with tool descriptions
2. **`mcp-simple.json`** - Minimal configuration
3. **`mcp-local.json`** - Local development configuration

## 🔧 **MCP Protocol Implementation**

<details>
<summary><strong>📝 Protocol Compliance</strong></summary>

This server implements **MCP JSON-RPC 2.0 Protocol** with required methods:

#### **1. `initialize`** - Handshake
```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "method": "initialize",
  "params": {
    "protocolVersion": "2024-11-05",
    "capabilities": {},
    "clientInfo": {"name": "test-client", "version": "1.0.0"}
  }
}
```

#### **2. `tools/list`** - Available Tools
```json
{
  "jsonrpc": "2.0",
  "id": 2,
  "method": "tools/list"
}
```

#### **3. `tools/call`** - Execute Tool
```json
{
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
}
```
</details>

## 🛠️ **Available Tools**

| Tool Name | Description | Required Parameters |
|-----------|-------------|-------------------|
| `welcome_message` | Get welcome message | None |
| `health_check` | Check service health | None |
| `create_product` | Create new product | `name`, `category`, `price` |
| `get_product` | Get product by ID | `id` |
| `update_product` | Update existing product | `id` + optional fields |
| `delete_product` | Delete product | `id` |
| `list_products` | Get all products | None |

## ⚙️ **Environment Variables**

- `MICROSERVICE_URL`: Target microservice URL
- `PORT`: Server port (defaults to 8080)

## 🔍 **Troubleshooting**

<details>
<summary><strong>❗ Common Issues</strong></summary>

1. **Method not found**: Use exact method names (`initialize`, `tools/list`, `tools/call`)
2. **Invalid JSON-RPC**: Always include `"jsonrpc": "2.0"` and unique `id`
3. **Authentication errors**: Verify GCP service account permissions
4. **Microservice connectivity**: Check `MICROSERVICE_URL` environment variable
</details>

<details>
<summary><strong>🤖 GitHub Copilot Integration</strong></summary>

1. Configure MCP server URL in Copilot settings
2. Copilot automatically discovers tools via MCP protocol
3. Use natural language: "Create a new product called 'iPhone 15'"
</details>

For testing examples and deployment commands, see [`tests/README.md`](../tests/README.md) and main project README.
```

### Option 2: Direct HTTP Configuration

For MCP clients that support direct HTTP connections:

```json
{
  "servers": {
    "ravi-mcp-server": {
      "transport": {
        "type": "http",
        "url": "https://ravi-mcp-server-256110662801.europe-west3.run.app/mcp"
      },
      "auth": {
        "type": "bearer",
        "token": "${GCLOUD_ACCESS_TOKEN}"
      }
    }
  }
}
```

## Natural Language Examples

Once configured, you can use these natural language commands:

### Product Management
- **"Show me all products in the store"** → Calls `list_products`
- **"Create a new MacBook Pro for $2499 in Electronics category"** → Calls `create_product`
- **"Get details for product ID abc123"** → Calls `get_product`
- **"Update product xyz789 price to $1299"** → Calls `update_product`
- **"Delete product with ID def456"** → Calls `delete_product`

### Service Operations
- **"Is the product service healthy?"** → Calls `health_check`
- **"Show me the welcome message"** → Calls `welcome_message`

### Advanced Queries
- **"Create 3 products: iPhone 15 ($999), MacBook Pro ($2499), and iPad Pro ($1299) all in Electronics"** → Multiple `create_product` calls
- **"List all products and then create a new Gaming Chair for $299"** → Calls `list_products` then `create_product`

## Setup Instructions

### For GitHub Copilot

1. **Install GitHub Copilot** with MCP support
2. **Add configuration** to your Copilot settings:
   ```bash
   # Add mcp.json content to your Copilot MCP configuration
   ```
3. **Set authentication**:
   ```bash
   export GCLOUD_ACCESS_TOKEN=$(gcloud auth print-identity-token)
   ```

### For VS Code with MCP Extension

1. **Install MCP extension** in VS Code
2. **Copy `mcp.json`** to your workspace or global MCP configuration
3. **Restart VS Code** to load the configuration

### For Claude Desktop (Anthropic)

1. **Open Claude Desktop settings**
2. **Add MCP server configuration**:
   ```json
   {
     "mcpServers": {
       "ravi-mcp-server": {
         "command": "curl",
         "args": [
           "-X", "POST", 
           "https://ravi-mcp-server-256110662801.europe-west3.run.app/mcp",
           "-H", "Content-Type: application/json",
           "-d", "@-"
         ]
       }
     }
   }
   ```

## Team Access Setup

### For Team Members to Use Your MCP Server

#### Option 1: Add Team Members to Google Cloud (Recommended)
```bash
# Add each team member's email to Cloud Run access
gcloud run services add-iam-policy-binding ravi-mcp-server \
  --region=europe-west3 \
  --member="user:teammate@company.com" \
  --role="roles/run.invoker"
```

Each team member then needs to:
1. Install Google Cloud SDK
2. Run `gcloud auth login`
3. Use the MCP configuration with `$(gcloud auth print-access-token)`

#### Option 2: Shared Service Account (Easier for Testing)
1. Create a service account for team testing
2. Generate a key file
3. Share the key file securely with team members
4. Team members use `GOOGLE_APPLICATION_CREDENTIALS`

See [TEAM_ACCESS.md](TEAM_ACCESS.md) for detailed setup instructions.

## Authentication Setup

### Google Cloud Authentication

```bash
# For production (using service account)
export GOOGLE_APPLICATION_CREDENTIALS="/path/to/your/service-account-key.json"

# For development (using user credentials)
gcloud auth login
export GCLOUD_ACCESS_TOKEN=$(gcloud auth print-identity-token)
```

### Environment Variables

```bash
# Set these in your environment
export MCP_SERVER_URL="https://ravi-mcp-server-256110662801.europe-west3.run.app/mcp"
export GCLOUD_PROJECT_ID="your-project-id"
```

## Testing the Configuration

### Test with cURL
```bash
curl -X POST https://ravi-mcp-server-256110662801.europe-west3.run.app/mcp \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $(gcloud auth print-identity-token)" \
  -d '{
    "jsonrpc": "2.0",
    "id": 1,
    "method": "tools/list"
  }'
```

### Test Natural Language
Once configured in your MCP client:
- Type: **"List all products"**
- Expected: Client calls your MCP server and shows product list
- Type: **"Create a new iPhone 15 for $999"**
- Expected: Client creates the product and confirms creation

## Troubleshooting

### Common Issues

1. **Authentication Errors**
   ```bash
   # Refresh your token
   gcloud auth login
   gcloud auth print-identity-token
   ```

2. **Server Not Found**
   - Verify your server URL is correct
   - Check if server is deployed and running

3. **MCP Client Not Detecting Server**
   - Restart your MCP client
   - Check configuration file syntax
   - Verify file is in correct location

### Debug Commands

```bash
# Test server directly
curl -X POST https://ravi-mcp-server-256110662801.europe-west3.run.app/mcp \
  -H "Content-Type: application/json" \
  -d '{"jsonrpc":"2.0","id":1,"method":"initialize","params":{"protocolVersion":"2024-11-05","capabilities":{},"clientInfo":{"name":"test","version":"1.0.0"}}}'

# Check server logs
gcloud run services logs read ravi-mcp-server --region=europe-west3
```

## Next Steps

1. **Deploy your updated server** to production
2. **Configure your preferred MCP client** using the provided configurations
3. **Test natural language commands** to ensure everything works
4. **Share with your team** by providing them the MCP configuration

Your MCP server is now ready for natural language interactions! 🎉
