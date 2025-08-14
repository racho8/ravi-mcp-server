# MCP Web Client Usage Guide

## 🚀 How to Use mcp_web_client.html

The MCP Web Client is a browser-based interface for testing your MCP server. It provides a user-friendly way to interact with your product management tools.

### 📋 Prerequisites

1. **Google Cloud SDK installed** on your machine
2. **Valid Google Cloud authentication** 
3. **Access to your MCP server** (either through IAM or service account)

### 🎯 Step-by-Step Usage

#### Step 1: Open the Web Client
```bash
# Navigate to the clients directory
cd mcp-clients

# Open the web client in your browser
open mcp_web_client.html
# or on Linux/Windows:
# xdg-open mcp_web_client.html
# start mcp_web_client.html
```

#### Step 2: Get Your Access Token
```bash
# In terminal, get your access token
gcloud auth print-access-token
```

#### Step 3: Authenticate in the Web Client
1. Copy the access token from your terminal
2. Paste it into the "Access Token" field in the web client
3. Click "Test Auth" to verify connection
4. You should see "✅ Authentication successful!"

#### Step 4: Test Your MCP Server

##### Quick Actions
Click any of these buttons to test basic functionality:
- **📦 List Products** - Shows all products in your store
- **❤️ Health Check** - Verifies server is running
- **👋 Welcome Message** - Gets welcome message
- **🔧 Show Available Tools** - Lists all available MCP tools

##### Custom Tool Calls
1. **Select a tool** from the dropdown menu
2. **Enter arguments** in JSON format (if required)
3. **Click "Call Tool"** to execute

### 🎨 Example Usage Scenarios

#### Scenario 1: Check What Products Exist
1. Click "📦 List Products"
2. View the product list in the output section

#### Scenario 2: Create a New Product
1. Select "create_product" from the dropdown
2. Enter arguments: `{"name":"MacBook Pro","category":"Electronics","price":2499}`
3. Click "Call Tool"
4. Check the output for the created product ID

#### Scenario 3: Update a Product Price
1. Select "update_product" from the dropdown
2. Enter arguments: `{"id":"PRODUCT_ID_HERE","price":1999}`
3. Click "Call Tool"
4. Verify the price was updated

#### Scenario 4: Delete a Product
1. Select "delete_product" from the dropdown
2. Enter arguments: `{"id":"PRODUCT_ID_HERE"}`
3. Click "Call Tool"
4. Confirm the product was deleted

### 📊 Understanding the Output

The web client now features **dual output formats** for the best user experience:

#### 🎨 Visual Mode (Default)
- **Formatted tables** for product listings with clean columns
- **Color-coded status messages** with timestamps
- **Tool descriptions** in organized tables
- **Server connection status** with visual indicators
- **Professional styling** perfect for demos and business users

#### 📋 JSON Mode (Technical)
- **Raw JSON responses** for debugging and development
- **Complete technical data** with full object structures
- **Exact server responses** for integration testing
- **Command-line style output** for developers

#### 🔄 Toggle Between Modes
- Click **"📊 Visual"** for user-friendly formatted output
- Click **"📋 JSON"** for raw technical data
- Click **"🗑️ Clear"** to reset both output areas
- Both modes update simultaneously with each operation

#### 📈 Output Features
- **Timestamps** for every operation
- **Success/Error indicators** (✅/❌) 
- **Product tables** with Name, Price, Category, and ID columns
- **Tools overview** showing available MCP tools and their requirements
- **Connection status** with server information display

### 🔧 Tool Arguments Reference

| Tool | Required Arguments | Example |
|------|-------------------|---------|
| `list_products` | None | `{}` |
| `create_product` | name, category, price | `{"name":"iPhone","category":"Electronics","price":999}` |
| `get_product` | id | `{"id":"abc123"}` |
| `update_product` | id, plus any field to update | `{"id":"abc123","price":899}` |
| `delete_product` | id | `{"id":"abc123"}` |
| `health_check` | None | `{}` |
| `welcome_message` | None | `{}` |

### 🎯 Best Practices

1. **Always test authentication first** before trying other operations
2. **Keep your access token secure** - don't share it
3. **Refresh your token** if you get authentication errors (tokens expire)
4. **Use the health check** to verify server status before complex operations
5. **Copy product IDs** from list_products to use in other operations
6. **Use Visual mode** for presentations and business reviews
7. **Switch to JSON mode** when debugging integration issues
8. **Clear output regularly** to keep the interface clean during testing sessions

### 🔍 Troubleshooting

#### Authentication Issues
```bash
# If you get auth errors, refresh your token:
gcloud auth login
gcloud auth print-access-token
```

#### Server Connection Issues
- Verify your server is deployed and running
- Check the server URL in the web client matches your deployment
- Try the health check first

#### CORS Issues
- The web client now handles CORS properly
- If you still see CORS errors, try using a different browser
- Alternatively, use the Python client instead

### 🚀 Advanced Usage

#### Dual Output Format Benefits
1. **Visual Mode for Demos**: Use when presenting to stakeholders or during team reviews
2. **JSON Mode for Development**: Switch when debugging API integration or troubleshooting
3. **Side-by-Side Testing**: Keep both views in mind - visual for UX, JSON for technical accuracy
4. **Output Format Switching**: Toggle between modes during the same session to see both perspectives

#### Batch Operations
1. List all products first to get IDs
2. Use those IDs to update or delete multiple products
3. Monitor the output section for each operation's success
4. Use Visual mode to quickly see table updates, JSON mode to verify exact responses

#### Testing Workflows
1. Create a product
2. List products to verify it exists
3. Update the product price
4. Get the specific product to verify the update
5. Delete the product
6. List products again to verify deletion

### 📱 Mobile/Responsive Usage
The web client is responsive and works on mobile devices, making it easy to test your MCP server from anywhere.

### 🔗 Integration with Team
Share the web client HTML file with your team members along with the authentication instructions from the TEAM_ACCESS.md guide.

---

The MCP Web Client provides an intuitive way to test your server without needing command-line tools. Perfect for demonstrations, debugging, and team collaboration! 🎉
