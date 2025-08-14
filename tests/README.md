# Testing Suite

This directory contains comprehensive testing tools for the MCP server. Choose the testing method that best fits your needs.

## 🎯 **Quick Start**

### **Automated Testing (Recommended):**
```bash
# Run all tests
./run_tests.sh

# Validate MCP protocol compliance
./validate_mcp.sh
```

### **Manual Testing:**
```bash
# View manual cURL commands
./test_commands.sh

# Import Postman collection
# See POSTMAN_IMPORT.md for instructions
```

## 📁 **Files in This Directory**

| File | Purpose | Usage |
|------|---------|-------|
| **`README.md`** | This guide | Documentation |
| **`run_tests.sh`** | Automated test runner | `./run_tests.sh` |
| **`validate_mcp.sh`** | MCP protocol validator | `./validate_mcp.sh` |
| **`test_commands.sh`** | Manual cURL commands | `./test_commands.sh` |
| **`postman_collection_v2.json`** | Postman collection | Import into Postman |
| **`POSTMAN_IMPORT.md`** | Postman setup guide | Follow instructions |
| **`test_mcp_requests.json`** | JSON test examples | Reference data |

## 🚀 **Testing Methods**

### **1. Automated Shell Testing**

#### **Full Test Suite (`run_tests.sh`)**
- **Purpose**: Complete automated testing of all MCP tools
- **Best for**: CI/CD, regular testing, comprehensive validation
- **Requirements**: Running MCP server, `curl`, `jq`

```bash
# Make sure server is running first
go run ../main.go &

# Run all tests
./run_tests.sh

# Stop server
pkill -f "go run ../main.go"
```

#### **Protocol Validation (`validate_mcp.sh`)**
- **Purpose**: Validates MCP JSON-RPC 2.0 protocol compliance
- **Best for**: Ensuring protocol adherence, debugging protocol issues
- **Tests**: Initialize, tools/list, tools/call, error handling

```bash
./validate_mcp.sh
```

### **2. Manual cURL Testing**

#### **Interactive Commands (`test_commands.sh`)**
- **Purpose**: Displays manual cURL commands for copy-paste testing
- **Best for**: Learning the API, custom testing, debugging specific requests
- **Output**: Ready-to-use cURL commands

```bash
# Display all commands
./test_commands.sh

# Example: Copy a command and run it
curl -X POST http://localhost:8080/mcp \
  -H "Content-Type: application/json" \
  -d '{"jsonrpc":"2.0","id":1,"method":"tools/list"}'
```

### **3. Postman Testing**

#### **GUI Testing (`postman_collection_v2.json`)**
- **Purpose**: Interactive testing with Postman's user interface
- **Best for**: Visual testing, team collaboration, request modification
- **Setup**: See `POSTMAN_IMPORT.md` for detailed instructions

```bash
# See import guide
cat POSTMAN_IMPORT.md
```

### **4. JSON Reference Testing**

#### **Request Examples (`test_mcp_requests.json`)**
- **Purpose**: JSON examples for all MCP requests
- **Best for**: Understanding request formats, integration reference
- **Content**: Properly formatted JSON-RPC 2.0 requests

```bash
# View example requests
cat test_mcp_requests.json | jq .
```

## 🎯 **Testing Scenarios**

### **Development Workflow:**
```bash
# 1. Start server
cd ..
go run main.go &

# 2. Quick validation
cd tests
./validate_mcp.sh

# 3. Full test suite
./run_tests.sh

# 4. Clean up
pkill -f "go run main.go"
```

### **Debugging Issues:**
```bash
# 1. Manual step-by-step testing
./test_commands.sh

# 2. Copy specific commands to test individually
# 3. Use Postman for interactive debugging
```

### **Team Collaboration:**
```bash
# 1. Share Postman collection
# Import postman_collection_v2.json

# 2. Use web client for non-technical users
cd ../mcp-clients
open mcp_web_client.html
```

## 🔧 **Prerequisites**

### **Local Development:**
```bash
# Required tools
which curl    # HTTP client
which jq      # JSON processor
which go      # Go runtime

# Start server
export MICROSERVICE_URL="https://product-service-256110662801.europe-west3.run.app"
go run ../main.go
```

### **Production Testing:**
```bash
# Update server URLs in test files to point to production
# Example: https://ravi-mcp-server-256110662801.europe-west3.run.app/mcp

# Add authentication headers if needed
# -H "Authorization: Bearer $(gcloud auth print-access-token)"
```

## 📋 **Test Coverage**

### **Core MCP Protocol:**
- ✅ **Initialize** - Server initialization and capabilities
- ✅ **Tools List** - Available tools enumeration
- ✅ **Tools Call** - Tool execution with parameters
- ✅ **Error Handling** - Invalid methods and parameters

### **Product Management Tools:**
- ✅ **list_products** - Get all products
- ✅ **create_product** - Create new product
- ✅ **get_product** - Get specific product by ID
- ✅ **update_product** - Update product fields
- ✅ **delete_product** - Remove product
- ✅ **health_check** - Server health status
- ✅ **welcome_message** - Welcome message

### **Error Scenarios:**
- ✅ **Invalid methods** - Non-existent JSON-RPC methods
- ✅ **Invalid tools** - Non-existent tool names
- ✅ **Missing parameters** - Required parameters not provided
- ✅ **Invalid JSON** - Malformed request bodies

## 🔍 **Troubleshooting**

### **Common Issues:**

#### **Connection Refused:**
```bash
# Check if server is running
curl http://localhost:8080/health
# or
lsof -i :8080
```

#### **Command Not Found:**
```bash
# Install missing tools
brew install curl jq  # macOS
apt install curl jq   # Ubuntu
```

#### **Authentication Errors:**
```bash
# Get fresh token
gcloud auth print-access-token

# Update test scripts with proper auth headers
```

#### **JSON Parse Errors:**
```bash
# Validate JSON syntax
echo '{"test":"json"}' | jq .

# Check server response format
curl -v http://localhost:8080/mcp
```

### **Debug Mode:**
```bash
# Add -v flag to curl commands for verbose output
curl -v -X POST http://localhost:8080/mcp ...

# Use jq for pretty-printing responses
curl ... | jq .
```

## 🎉 **Integration with Other Testing**

### **With MCP Clients:**
```bash
# Web client
cd ../mcp-clients
open mcp_web_client.html

# Python client
cd ../mcp-clients
python mcp_test_client.py "list all products"
```

### **With CI/CD:**
```bash
# Add to GitHub Actions or similar
./tests/run_tests.sh
./tests/validate_mcp.sh
```

### **With Monitoring:**
```bash
# Regular health checks
watch -n 30 './tests/validate_mcp.sh'
```

---

**Choose your testing method based on your needs:**
- **Developers** → `run_tests.sh` and `validate_mcp.sh`
- **QA/Testing** → Postman collection
- **Learning/Debug** → `test_commands.sh`
- **Integration** → `test_mcp_requests.json`

All testing tools work independently and can be used together for comprehensive testing! 🚀
