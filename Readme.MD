# Ravi MCP Server

A **Model Context Protocol (MCP)** server that provides product management capabilities through a clean JSON-RPC 2.0 API. Deployed on Google Cloud Run with auto-scaling and high availability.

## **What is This?**

This server acts as a bridge between AI clients (like Claude Desktop) and a product database, enabling natural language interactions with structured data operations.

**Key Features:**
- MCP JSON-RPC 2.0 compliant
- Product CRUD operations 
- Google Cloud authentication
- Browser and CLI testing clients
- Microservices architecture

## **Architecture**

![Architecture Diagram](docs/architecture-diagram.svg)

**Flow:** MCP Clients → Ravi MCP Server → Product Service → Cloud Datastore

**Components:**
- **MCP Clients**: Python CLI, Web browser, Claude Desktop integration
- **MCP Server**: Go application handling JSON-RPC requests (Cloud Run)
- **Product Service**: Business logic microservice (Cloud Run) 
- **Cloud Datastore**: NoSQL database for product storage

## **Quick Test (2 minutes)**

### **1. Get Access Token**
```bash
gcloud auth print-access-token
```

### **2. Choose Your Testing Method**

**Option A: Web Client (Easiest)**
```bash
cd mcp-clients
open mcp_web_client.html
# Paste your token and test in browser
```

**Option B: Python Client**
```bash
cd mcp-clients  
python mcp_test_client.py "list all products"
```

**Option C: Direct API Call**
```bash
curl -H "Authorization: Bearer $(gcloud auth print-access-token)" \
  https://ravi-mcp-server-256110662801.europe-west3.run.app/mcp \
  -d '{"jsonrpc":"2.0","id":1,"method":"tools/call","params":{"name":"health_check","arguments":{}}}'
```

## **Available Tools**

| Tool | Purpose | Example Usage |
|------|---------|---------------|
| `list_products` | Get all products | "show me all products" |
| `create_product` | Add new product | "create product: Laptop, Electronics, $999" |
| `get_product` | Get specific product | "get product with ID 123" |
| `update_product` | Modify product | "update product 123 price to $899" |
| `delete_product` | Remove product | "delete product 123" |
| `health_check` | Server status | "health check" |

## **Local Development**

### **Setup Go Environment**
```bash
# Quick setup (handles version issues)
./setup-go-env.sh

# Verify
go version  # Should be 1.23+
```

### **Run Locally**
```bash
# Install dependencies
go mod download

# Run tests
go test -v ./...

# Start server
export MICROSERVICE_URL="https://product-service-256110662801.europe-west3.run.app"
go run main.go
```

## **Integration**

### **Claude Desktop**
```bash
cp config/mcp.json ~/.config/claude/claude_desktop_config.json
# Restart Claude Desktop, then ask: "Show me all products"
```

### **Custom Applications**
- **Endpoint**: `https://ravi-mcp-server-256110662801.europe-west3.run.app/mcp`
- **Protocol**: MCP JSON-RPC 2.0
- **Auth**: Google Cloud Bearer tokens

## **Project Structure**

```
ravi-mcp-server/
├── main.go                        # Core server
├── mcp-clients/                   # Testing clients
│   ├── mcp_test_client.py        # Python CLI
│   └── mcp_web_client.html       # Web GUI
├── config/                        # MCP configurations
├── docs/                          # Detailed guides
├── tests/                         # Complete testing suite
└── setup-go-env.sh               # Environment setup
```

## **Deployment & Management**

### **Cloud Run Commands**
```bash
# Update environment variables
gcloud run services update ravi-mcp-server \
    --region=europe-west3 \
    --set-env-vars "MICROSERVICE_URL=https://product-service-256110662801.europe-west3.run.app"

# Check current environment variables
gcloud run services describe ravi-mcp-server \
  --region=europe-west3 \
  --format="value(spec.template.spec.containers[0].env)"

# Add IAM access (if needed)
gcloud run services add-iam-policy-binding ravi-mcp-server \
  --region=europe-west3 \
  --member="domain:your-organization.com" \
  --role="roles/run.invoker"
```

### **Required Secrets (GitHub Actions)**
- `GCP_SA_KEY`: Google Cloud service account key
- `GCP_PROJECT_ID`: Google Cloud project ID  
- `MICROSERVICE_URL`: Target microservice URL

## **Need More Details?**

- **Testing**: [`tests/README.md`](tests/README.md) - Comprehensive testing guide
- **Client Usage**: [`mcp-clients/README.md`](mcp-clients/README.md) - Complete client guide
- **Team Setup**: [`docs/TEAM_ACCESS.md`](docs/TEAM_ACCESS.md) - Team authentication
- **Configuration**: [`config/MCP_CONFIGURATION.md`](config/MCP_CONFIGURATION.md) - MCP setup

---

**Ready to test?** Start with the web client for easiest experience, or dive into the detailed guides above!
