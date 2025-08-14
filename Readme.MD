# Ravi MCP Server

A **Model Context Protocol (MCP)** server implementation in Go that provides product management capabilities. This server follows the MCP JSON-RPC 2.0 specification and is deployed on Google Cloud Run.

## **Quick Start**

### **Test the Server:**
```bash
# Navigate to testing clients
cd mcp-clients

# Option 1: Python client (command line)
python mcp_test_client.py "list all products"

# Option 2: Web client (browser)
open mcp_web_client.html
```

### **Authentication:**
```bash
# Get your access token
gcloud auth print-access-token
```

## **Project Structure**

```
ravi-mcp-server/
├── main.go                   # Core MCP server implementation
├── go.mod, go.sum           # Go dependencies
├── Dockerfile               # Container configuration
│
├── mcp-clients/             # Testing Interfaces
│   ├── README.md            # Client usage guide
│   ├── mcp_test_client.py   # Python command-line client
│   └── mcp_web_client.html  # Browser-based GUI client
│
├── config/                  # MCP configurations
│   ├── mcp.json            # Full configuration
│   ├── mcp-local.json      # Local development
│   └── mcp-simple.json     # Minimal config
│
├── docs/                    # Documentation
│   ├── architecture-diagram.svg    # System architecture diagram
│   ├── MCP_CONFIGURATION.md
│   ├── TEAM_ACCESS.md
│   ├── TESTING.md
│   └── WEB_CLIENT_GUIDE.md
│
├── tests/                   # Complete Testing Suite
│   ├── README.md            # Complete testing guide
│   ├── POSTMAN_IMPORT.md    # Postman import instructions
│   ├── postman_collection_v2.json  # Postman collection
│   ├── run_tests.sh         # Automated test runner
│   ├── test_commands.sh     # Manual cURL commands
│   ├── test_mcp_requests.json # JSON test examples
│   └── validate_mcp.sh      # MCP protocol validation
│
└── setup-go-env.sh          # Go environment setup script
```

## **Architecture Overview**

The Ravi MCP Server follows a distributed microservices architecture deployed on Google Cloud Platform. Below is the complete system architecture showing data flow from MCP clients to the database:

![Architecture Diagram](docs/architecture-diagram.svg)

### **Component Overview:**

**Local Environment:**
- **MCP Clients**: Python CLI, Web Browser, and Claude Desktop integration
- **Authentication**: Google Cloud IAM with JWT tokens

**Google Cloud Platform:**
- **Ravi MCP Server**: Go-based JSON-RPC 2.0 server (Cloud Run)
- **Product Service**: Business logic microservice (Cloud Run)
- **Cloud Datastore**: NoSQL database for product data

### **Data Flow:**
1. **Client Authentication**: Users authenticate via `gcloud auth` to get JWT tokens
2. **MCP Protocol**: Clients send JSON-RPC 2.0 requests with Bearer tokens
3. **Request Routing**: MCP Server validates tokens and routes to Product Service
4. **Business Logic**: Product Service processes CRUD operations
5. **Data Persistence**: Service interacts with Cloud Datastore for data storage
6. **Response Chain**: Results flow back through the same path to clients

## **MCP Tools Available**

| Tool | Description | Arguments |
|------|-------------|-----------|
| `list_products` | Get all products in store | None |
| `create_product` | Create new product | name, category, price |
| `get_product` | Get specific product | id |
| `update_product` | Update product details | id, (fields to update) |
| `delete_product` | Remove product | id |
| `health_check` | Server health status | None |
| `welcome_message` | Welcome message | None |

## **Server Features**

- **MCP JSON-RPC 2.0 Compliant** - Full protocol implementation
- **Google Cloud Run Deployment** - Scalable and secure
- **CORS Support** - Works with browser clients
- **JWT Authentication** - Google Cloud IAM integration
- **Product Management** - Complete CRUD operations
- **Health Monitoring** - Built-in health checks

## **Development**

### **Go Environment Setup:**
This project requires Go 1.23+ and uses specific version configuration to avoid compiler/toolchain mismatches.

```bash
# Quick setup (recommended)
./setup-go-env.sh

# Manual setup
export GOROOT="/opt/homebrew/Cellar/go/1.23.4/libexec"  # Adjust version as needed
export PATH="$GOROOT/bin:$PATH"

# Verify setup
go version  # Should show go1.23.4 or later
```

### **Local Development:**
```bash
# Install dependencies
go mod download

# Run tests
go test -v ./...

# Run server locally
export MICROSERVICE_URL="https://product-service-256110662801.europe-west3.run.app"
go run main.go

# Test with local config
# Update config/mcp-local.json with localhost URL
```

### **Deploy to Cloud Run:**

## **Testing**

### **Quick Tests:**
```bash
# Python client (developers)
cd mcp-clients
python mcp_test_client.py "health check"

# Web client (everyone)
open mcp-clients/mcp_web_client.html
```

### **Comprehensive Testing:**
```bash
# See complete testing guide
cd tests
cat README.md

# Run all tests
./run_tests.sh

# Validate MCP compliance
./validate_mcp.sh

# Postman collection
# See POSTMAN_IMPORT.md for import instructions
```

## **Team Access**

For team members to test the server, see detailed instructions in:
- **`docs/TEAM_ACCESS.md`** - Authentication setup for team members
- **`mcp-clients/README.md`** - How to use both testing clients

### **Quick Team Setup:**
1. **Get access token**: `gcloud auth print-access-token`
2. **Choose client**: Python (developers) or Web (everyone)
3. **Start testing**: Follow client-specific instructions

## **Integration**

### **Claude Desktop Integration:**
```bash
# Copy configuration
cp config/mcp.json ~/.config/claude/claude_desktop_config.json

# Restart Claude Desktop
# Ask: "Show me all products in the store"
```

### **Custom Integration:**
- **Server URL**: `https://ravi-mcp-server-256110662801.europe-west3.run.app/mcp`
- **Protocol**: MCP JSON-RPC 2.0
- **Authentication**: Google Cloud Bearer tokens
- **Content-Type**: `application/json`

## **Documentation**

- **`mcp-clients/README.md`** - Complete client usage guide
- **`docs/MCP_CONFIGURATION.md`** - MCP setup instructions
- **`docs/TEAM_ACCESS.md`** - Team authentication guide
- **`docs/TESTING.md`** - Testing procedures
- **`docs/WEB_CLIENT_GUIDE.md`** - Web client details

## **Server Endpoint**

**Production URL**: `https://ravi-mcp-server-256110662801.europe-west3.run.app/mcp`

### **Health Check:**
```bash
curl -H "Authorization: Bearer $(gcloud auth print-access-token)" \
  https://ravi-mcp-server-256110662801.europe-west3.run.app/mcp \
  -d '{"jsonrpc":"2.0","id":1,"method":"tools/call","params":{"name":"health_check","arguments":{}}}'
```

## **Getting Started**

1. **Authenticate**: `gcloud auth login`
2. **Test server**: `cd mcp-clients && python mcp_test_client.py "health check"`
3. **Explore tools**: Use web client for interactive testing
4. **Read docs**: Check `docs/` for detailed guides
5. **Integrate**: Use with Claude Desktop or custom applications

---

**Your MCP server is ready for testing and integration!**

For questions or issues, check the documentation in the `docs/` folder or test with the clients in `mcp-clients/`.
