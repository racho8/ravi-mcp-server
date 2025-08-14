# MCP Testing Clients

This folder contains **two independent client interfaces** for testing your MCP server. Both clients connect to the same server but provide different user experiences.

## 🎯 **Overview**

| Client | Interface | Best For | Users |
|--------|-----------|----------|-------|
| **Python Client** | Command Line | Development, Automation, CI/CD | Developers, Scripts |
| **Web Client** | Browser GUI | Demos, Teams, Interactive Testing | Everyone |

## 🐍 **Python Client** (`mcp_test_client.py`)

### **Quick Start:**
```bash
# Navigate to clients folder
cd mcp-clients

# Run with natural language commands
python mcp_test_client.py "list all products"
python mcp_test_client.py "health check"
python mcp_test_client.py "show available tools"
```

### **Features:**
- ✅ **Natural Language Input**: Type commands in plain English
- ✅ **Auto-Authentication**: Automatically gets GCP token via `gcloud`
- ✅ **Terminal Output**: Clean JSON responses in command line
- ✅ **Script-Friendly**: Perfect for automation and CI/CD
- ✅ **No Browser Required**: Works in headless environments

### **Prerequisites:**
```bash
# Install Google Cloud SDK
gcloud auth login
gcloud config set project YOUR_PROJECT_ID
```

### **Available Commands:**
- `"list all products"` → Call list_products tool
- `"health check"` → Check server health
- `"welcome message"` → Get welcome message
- `"show available tools"` → List all MCP tools

---

## 🌐 **Web Client** (`mcp_web_client.html`)

### **Quick Start:**
```bash
# Open in browser (choose one):
open mcp_web_client.html           # macOS
xdg-open mcp_web_client.html       # Linux
start mcp_web_client.html          # Windows

# Or drag the file into your browser
```

### **Features:**
- ✅ **Dual Output Formats**: Visual tables + Raw JSON
- ✅ **Interactive GUI**: Point-and-click interface
- ✅ **Live Authentication**: Test connection in real-time
- ✅ **Professional Styling**: Perfect for demos and presentations
- ✅ **Mobile-Friendly**: Works on any device with a browser
- ✅ **Team-Friendly**: No command-line skills required

### **Authentication Setup:**
1. **Get your token**: `gcloud auth print-access-token`
2. **Paste in web client**: Copy token to the "Access Token" field
3. **Test connection**: Click "Test Auth" button
4. **Start testing**: Use quick actions or custom tool calls

### **Output Modes:**
- **📊 Visual Mode**: Formatted tables, colors, timestamps (Default)
- **📋 JSON Mode**: Raw technical data for debugging
- **Toggle anytime**: Switch between modes with toolbar buttons

---

## 🔄 **How They Relate**

### **Completely Independent:**
- ✅ **No shared dependencies** - run separately
- ✅ **Different authentication** - Python auto-gets token, Web requires manual input
- ✅ **Separate processes** - can run simultaneously
- ✅ **No data sharing** - each maintains its own state

### **Same Backend:**
- 🎯 **Same MCP Server**: `https://ravi-mcp-server-256110662801.europe-west3.run.app/mcp`
- 🎯 **Same Protocol**: JSON-RPC 2.0 format
- 🎯 **Same Tools**: list_products, health_check, welcome_message, etc.
- 🎯 **Same Authentication**: GCP Bearer tokens

### **Choose Based on Use Case:**

#### Use **Python Client** for:
- 🔧 **Development workflows** and debugging
- 🤖 **Automated testing** and CI/CD pipelines
- ⚡ **Quick command-line checks**
- 📝 **Scripting and batch operations**

#### Use **Web Client** for:
- 👥 **Team demonstrations** and stakeholder meetings
- 🎨 **Interactive testing** with visual feedback
- 📊 **Business user access** (no technical skills needed)
- 🚀 **Professional presentations** and demos

---

## 🚀 **Quick Test Guide**

### **Test Both Clients:**

1. **Authenticate both**:
   ```bash
   # Get token for both clients
   gcloud auth print-access-token
   ```

2. **Test Python client**:
   ```bash
   cd mcp-clients
   python mcp_test_client.py "health check"
   python mcp_test_client.py "list all products"
   ```

3. **Test Web client**:
   - Open `mcp_web_client.html` in browser
   - Paste token and click "Test Auth"
   - Try quick actions: "Health Check", "List Products"

4. **Verify independence**:
   - Create product in Python client
   - View it in Web client (both see same data)
   - Both work simultaneously without interference

---

## 📁 **File Structure**

```
mcp-clients/
├── README.md                 # This guide
├── mcp_test_client.py       # Python command-line client
├── mcp_web_client.html      # Browser-based GUI client
└── examples/                # Usage examples (future)
```

---

## 🔧 **Troubleshooting**

### **Common Issues:**

#### Authentication Errors:
```bash
# Refresh your GCP authentication
gcloud auth login
gcloud auth print-access-token
```

#### Python Client Issues:
```bash
# Check gcloud is installed
gcloud --version

# Check Python version (3.6+)
python --version
```

#### Web Client Issues:
- **CORS errors**: Server now supports CORS, try different browser
- **Token expired**: Get fresh token with `gcloud auth print-access-token`
- **Connection failed**: Verify server URL and network connection

---

## 📖 **Additional Documentation**

- **Server Implementation**: See `main.go` in root directory
- **Team Access Guide**: See `docs/TEAM_ACCESS.md`
- **Web Client Details**: See `docs/WEB_CLIENT_GUIDE.md`
- **Deployment Guide**: See `README.md` in root directory

---

**Both clients provide full access to your MCP server - choose the interface that best fits your workflow!** 🎯
