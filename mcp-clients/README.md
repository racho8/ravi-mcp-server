# MCP Testing Clients

This folder contains **two independent client interfaces** for testing your MCP server. Both clients connect to the same server but provide different user experiences.

## 🎯 **Quick Start Guide**

<details>
<summary><strong>🌐 Option 1: Web Client (Easiest)</strong></summary>

```bash
# Open in browser
open mcp_web_client.html  # macOS
# xdg-open mcp_web_client.html   # Linux  
# start mcp_web_client.html      # Windows

# Get your token and paste in web client
gcloud auth print-access-token
```
</details>

<details>
<summary><strong>🐍 Option 2: Python Client (Developers)</strong></summary>

```bash
# Run with natural language commands
python mcp_test_client.py "list all products"
python mcp_test_client.py "health check"
```
</details>

## 📊 **Client Comparison**

| Feature | **Web Client** | **Python Client** |
|---------|----------------|-------------------|
| **Interface** | Browser GUI | Command Line |
| **Best For** | Demos, Teams, Non-technical users | Development, Automation, Scripts |
| **Authentication** | Manual token paste | Auto via gcloud |
| **Output Format** | Visual tables + JSON toggle | Clean JSON |
| **Prerequisites** | Browser, gcloud auth | Python, gcloud CLI |
| **Learning Curve** | Beginner-friendly | Developer-oriented |

---

## 🌐 **Web Client Guide** (`mcp_web_client.html`)

<details>
<summary><strong>📋 Step-by-Step Usage</strong></summary>

#### 1. **Open the Web Client**
```bash
cd mcp-clients
open mcp_web_client.html
```

#### 2. **Get Authentication Token**
```bash
gcloud auth print-access-token
```

#### 3. **Authenticate in Browser**
1. Copy the access token from terminal
2. Paste into "Access Token" field in web client
3. Click "Test Auth" to verify connection
4. Look for "✅ Authentication successful!"

#### 4. **Test Your Server**

**Quick Actions (Click and Go):**
- **📦 List Products** - Shows all products
- **❤️ Health Check** - Server status  
- **👋 Welcome Message** - Welcome response
- **🔧 Show Available Tools** - List all MCP tools

**Custom Tool Calls:**
1. Select tool from dropdown menu
2. Enter arguments in JSON format (if needed)
3. Click "Call Tool" to execute
</details>

<details>
<summary><strong>✨ Web Client Features</strong></summary>

#### **Dual Output Modes:**
- **📊 Visual Mode**: Formatted tables, colors, timestamps (Default)
- **📋 JSON Mode**: Raw technical data for debugging
- **Toggle anytime**: Switch between modes instantly

#### **Example Usage Scenarios:**

**Create a Product:**
1. Select "create_product" from dropdown
2. Enter: `{"name":"MacBook Pro","category":"Electronics","price":2499}`
3. Click "Call Tool"

**Update Product Price:**
1. Select "update_product" 
2. Enter: `{"id":"PRODUCT_ID","price":1999}`
3. Click "Call Tool"
</details>

<details>
<summary><strong>📖 Tool Arguments Reference</strong></summary>

| Tool | Arguments | Example |
|------|-----------|---------|
| `list_products` | None | `{}` |
| `create_product` | name, category, price | `{"name":"iPhone","category":"Electronics","price":999}` |
| `get_product` | id | `{"id":"abc123"}` |
| `update_product` | id + fields to update | `{"id":"abc123","price":899}` |
| `delete_product` | id | `{"id":"abc123"}` |
| `health_check` | None | `{}` |
</details>

---

## 🐍 **Python Client Guide** (`mcp_test_client.py`)

<details>
<summary><strong>⚡ Quick Commands</strong></summary>

```bash
cd mcp-clients

# Natural language interface
python mcp_test_client.py "list all products"
python mcp_test_client.py "health check" 
python mcp_test_client.py "show available tools"
python mcp_test_client.py "welcome message"
```
</details>

<details>
<summary><strong>✅ Features</strong></summary>

- ✅ **Natural Language Input**: Plain English commands
- ✅ **Auto-Authentication**: Gets GCP token automatically via `gcloud`
- ✅ **Terminal Output**: Clean JSON responses
- ✅ **Script-Friendly**: Perfect for automation and CI/CD
- ✅ **No Browser Required**: Works in headless environments
</details>

<details>
<summary><strong>📋 Prerequisites</strong></summary>

```bash
# Install Google Cloud SDK and authenticate
gcloud auth login
gcloud config set project YOUR_PROJECT_ID
```
</details>

---

## 🔄 **How Both Clients Work Together**

### **Completely Independent:**
- ✅ **No shared dependencies** - run separately
- ✅ **Different authentication approaches**
- ✅ **Can run simultaneously** without interference
- ✅ **No data sharing** between clients

### **Same Backend:**
- 🎯 **Same MCP Server**: `https://ravi-mcp-server-256110662801.europe-west3.run.app/mcp`
- 🎯 **Same Protocol**: JSON-RPC 2.0 format
- 🎯 **Same Tools**: All CRUD operations + health checks
- 🎯 **Same Data**: Both see same products and responses

### **Choose Based on Your Needs:**

#### Use **Web Client** for:
- � **Team demos** and stakeholder presentations
- 🎨 **Interactive testing** with visual feedback  
- 📊 **Business user access** (no command-line skills needed)
- � **Quick testing** and exploration

#### Use **Python Client** for:
- � **Development workflows** and debugging
- 🤖 **Automated testing** and CI/CD pipelines
- ⚡ **Quick command-line checks**
- � **Scripting** and batch operations

---

## 🚀 **Complete Test Workflow**

<details>
<summary><strong>🧪 Test Both Clients (5 minutes)</strong></summary>

1. **Setup Authentication:**
   ```bash
   gcloud auth login
   gcloud auth print-access-token  # Copy this token
   ```

2. **Test Python Client:**
   ```bash
   cd mcp-clients
   python mcp_test_client.py "health check"
   python mcp_test_client.py "list all products"
   ```

3. **Test Web Client:**
   - Open `mcp_web_client.html` in browser
   - Paste token and click "Test Auth"
   - Try "List Products" and "Health Check" buttons

4. **Verify They're Independent:**
   - Create product in Python: `python mcp_test_client.py "create product"`
   - View it in Web client (refresh and list products)
   - Both see the same data but work independently
</details>

---

## 🔧 **Troubleshooting**

<details>
<summary><strong>🔐 Authentication Issues</strong></summary>

```bash
# Refresh GCP authentication
gcloud auth login
gcloud auth print-access-token
```
</details>

<details>
<summary><strong>🐍 Python Client Problems</strong></summary>

```bash
# Check prerequisites
gcloud --version    # Should work
python --version    # Should be 3.6+
```
</details>

<details>
<summary><strong>🌐 Web Client Issues</strong></summary>

- **Token expired**: Get fresh token with `gcloud auth print-access-token`
- **CORS errors**: Try different browser or use Python client
- **Connection failed**: Verify server URL and network
</details>

---

## � **File Structure**

```
mcp-clients/
├── README.md                 # This complete guide
├── mcp_test_client.py       # Python CLI client
└── mcp_web_client.html      # Browser GUI client
```

---

**Both clients provide full access to your MCP server - choose the interface that fits your workflow!** 🎯

For server details and architecture, see the main project README.md in the root directory.
