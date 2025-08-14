# Team Access Setup Guide

# Team Access Setup Guide

## For Team Members to Test the MCP Server

<details>
<summary><strong>👤 Option 1: Individual Google Cloud Access (Recommended)</strong></summary>

### Step 1: Add Team Members to Cloud Run
```bash
# Replace with actual team member emails
gcloud run services add-iam-policy-binding ravi-mcp-server \
  --region=europe-west3 \
  --member="user:john.doe@company.com" \
  --role="roles/run.invoker"

gcloud run services add-iam-policy-binding ravi-mcp-server \
  --region=europe-west3 \
  --member="user:jane.smith@company.com" \
  --role="roles/run.invoker"
```

### Step 2: Team Members Setup Their Authentication
Each team member needs to:

1. **Install Google Cloud SDK**:
   ```bash
   # macOS
   brew install --cask google-cloud-sdk
   
   # Windows/Linux - download from cloud.google.com
   ```

2. **Authenticate with Google Cloud**:
   ```bash
   gcloud auth login
   gcloud config set project your-project-id
   ```

3. **Use the MCP Configuration**:
   ```json
   {
     "mcpServers": {
       "ravi-mcp-server": {
         "command": "curl",
         "args": [
           "-X", "POST",
           "https://ravi-mcp-server-256110662801.europe-west3.run.app/mcp",
           "-H", "Content-Type: application/json",
           "-H", "Authorization: Bearer $(gcloud auth print-access-token)",
           "-d", "@-"
         ]
       }
     }
   }
   ```
</details>

<details>
<summary><strong>🔑 Option 2: Shared Service Account (Easier for Testing)</strong></summary>

### Step 1: Create a Service Account for Team Testing
```bash
# Create service account
gcloud iam service-accounts create mcp-team-testing \
  --description="Service account for team MCP testing" \
  --display-name="MCP Team Testing"

# Grant Cloud Run Invoker role
gcloud run services add-iam-policy-binding ravi-mcp-server \
  --region=europe-west3 \
  --member="serviceAccount:mcp-team-testing@your-project-id.iam.gserviceaccount.com" \
  --role="roles/run.invoker"

# Create and download key
gcloud iam service-accounts keys create team-mcp-key.json \
  --iam-account=mcp-team-testing@your-project-id.iam.gserviceaccount.com
```

### Step 2: Share the Service Account Key
Share `team-mcp-key.json` securely with team members.

### Step 3: Team Members Use Service Account
```bash
# Set service account
export GOOGLE_APPLICATION_CREDENTIALS="/path/to/team-mcp-key.json"
gcloud auth activate-service-account --key-file=/path/to/team-mcp-key.json

# Test authentication
gcloud auth print-access-token
```
</details>

<details>
<summary><strong>🌐 Option 3: Public Access (For Demo/Testing Only)</strong></summary>

### Make Service Publicly Accessible (NOT RECOMMENDED FOR PRODUCTION)
```bash
gcloud run services add-iam-policy-binding ravi-mcp-server \
  --region=europe-west3 \
  --member="allUsers" \
  --role="roles/run.invoker"
```

Then team members can use without authentication:
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
</details>

## Testing Instructions for Team Members

<details>
<summary><strong>⚡ Quick Test Commands</strong></summary>

```bash
# Test server access
curl -X POST https://ravi-mcp-server-256110662801.europe-west3.run.app/mcp \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $(gcloud auth print-access-token)" \
  -d '{
    "jsonrpc": "2.0",
    "id": 1,
    "method": "tools/list",
    "params": {}
  }'

# Expected response: List of available tools
```
</details>

<details>
<summary><strong>🤖 Using with Claude Desktop</strong></summary>

1. Copy the MCP configuration to Claude Desktop config
2. Restart Claude Desktop
3. Ask: "Show me all products in the store"
</details>

<details>
<summary><strong>🐍 Using with Python Test Client</strong></summary>

```bash
cd mcp-clients
python mcp_test_client.py "list all products"
```
</details>

## Security Best Practices

<details>
<summary><strong>🔒 Security Guidelines</strong></summary>

1. **Use Individual Google Accounts** when possible
2. **Limit service account permissions** to only Cloud Run Invoker
3. **Rotate service account keys** regularly
4. **Monitor access logs** in Google Cloud Console
5. **Remove access** when team members leave
</details>

## Recommended Approach for Your Team

<details>
<summary><strong>📋 Choose Your Setup</strong></summary>

**For Development/Testing**: Use **Option 2 (Shared Service Account)**
- Easier to set up
- Single key to manage
- Can be rotated when needed

**For Production**: Use **Option 1 (Individual IAM)**
- Better audit trail
- Individual access control
- More secure
</details>

Would you like me to help you set up one of these options?

Would you like me to help you set up one of these options?
