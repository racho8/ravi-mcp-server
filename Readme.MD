# MCP Server

A **Model Context Protocol (MCP)** server that provides product management capabilities through a clean JSON-RPC 2.0 API. Deployed on Google Cloud Run with auto-scaling and high availability.

- Google Cloud credentials: Authentication token
<details>
<summary>Quick Test (2 minutes)</summary>

### For New Contributors: Get Access

**Required for Testing:**
- MICROSERVICE_URL: Backend service endpoint
- Google Cloud credentials: Authentication token

**How to get access:**
1. Contact Repository Owner: Reach out to [@racho8](https://github.com/racho8) via:
  - GitHub Issues: [Create new issue](https://github.com/racho8/ravi-mcp-server/issues/new)
  - Email: (see GitHub profile)
  - Slack/Teams: (if you're on the same team)
2. What to include in your request:
  - Your GitHub username
  - Intended use case (testing, development, integration)
  - Google Cloud account email (for IAM permissions)
3. What you'll receive:
  - MICROSERVICE_URL endpoint
  - Google Cloud project access
  - Setup instructions specific to your use case

For Team Members: See [`docs/TEAM_ACCESS.md`](docs/TEAM_ACCESS.md) for detailed access setup.

### 1. Get Access Token (After Getting Access)

```bash
gcloud auth login
gcloud auth print-access-token
```

### 2.Available Tools
<details>
<summary>Available Tools</summary>

| Tool                    | Purpose                              | Example Usage                                 |
|-------------------------|--------------------------------------|-----------------------------------------------|
| list_products           | Get all products                     | "show me all products"                        |
| create_product          | Add new product                      | "create product: Laptop, Electronics, $999"   |
| get_product             | Get specific product                 | "get product with ID 123"                     |
| update_product          | Modify product                       | "update product 123 price to $899"            |
| delete_product          | Remove product                       | "delete product 123"                          |
| create_multiple_products| Add multiple products                | "add these products: [...]"                   |
| update_products         | Update multiple products at once     | "update these products: [...]"                |
| delete_products         | Delete multiple products at once     | "delete these products: [...]"                |
| health_check            | Server status                        | "health check"                                |

</details>

go version  # Should be 1.23+
go test -v ./...
go run main.go

<details>
<summary>RPIM-API-Service Tools</summary>

### How to Run RPIM Tools

**Required Environment Variable:**

Set the RPIM API service URL before running the server or tests:

```bash
export RPIM_API_URL="https://rpim-api-service-385153724914.europe-west3.run.app"
```

**Example: Run MCP server with RPIM tools enabled**

```bash
export RPIM_API_URL="https://rpim-api-service-385153724914.europe-west3.run.app"
export MICROSERVICE_URL="<your_product_service_url>"
export PORT=8080
go run main.go
```

**Example: Call RPIM tool via JSON-RPC**

```bash
curl -H "Authorization: Bearer <YOUR_TOKEN>" \
  <SERVER_URL>/mcp \
  -d '{"jsonrpc":"2.0","id":1,"method":"tools/call","params":{"name":"rpim_get_class_units","arguments":{}}}'
```

**Example: Call RPIM child items tool**

```bash
curl --location "$RPIM_API_URL/api/v1/item-content" \
  --header "Authorization: Bearer <YOUR_TOKEN>" \
  --header "Content-Type: application/json" \
  --data '[{"itemNo": "20431739", "itemType": "ART", "classUnitCode": "GB", "classUnitType": "RU"}]'
```

**Available RPIM Tools:**

| Tool Name                  | Description                       |
|----------------------------|-----------------------------------|
| rpim_health_check          | Health check for RPIM API         |
| rpim_get_child_items       | Get child items                   |
| rpim_get_item_keys         | Get item keys                     |
| rpim_get_class_units       | Get class units                   |
| rpim_get_item_details      | Get item details                  |
| rpim_get_local_items       | Get local items                   |
| rpim_get_updated_items     | Get updated items                 |
| rpim_get_classified_items  | Get classified items              |
| rpim_get_item_attributes   | Get item attributes               |
| rpim_get_items_by_class_unit| Get items by class unit           |

See the Postman collection for sample payloads and endpoint details.


### For New Contributors: Setup Requirements

**Required Configuration:**
- MICROSERVICE_URL: Backend service endpoint (contact repository owner)
- Google Cloud credentials: Project access for authentication

**To get started with local development:**
1. Request access from [@racho8](https://github.com/racho8) (see "Quick Test" section above)
2. Follow the setup steps below after receiving your configuration

### Setup Go Environment

```bash
./setup-go-env.sh
go version  # Should be 1.23+
```

### Run Locally (After Getting Access)

```bash
go mod download
go test -v ./...
export MICROSERVICE_URL="<MICROSERVICE_URL_PROVIDED_BY_OWNER>"
export PORT="8080"
go run main.go
curl http://localhost:8080/health
```

**Tip:** Create a `.env` file with your environment variables for convenience:
```bash
MICROSERVICE_URL=<your_microservice_url>
PORT=8080
```

</details>

<details>
<summary>Integration</summary>

## Project Structure

```
ravi-mcp-server/
├── main.go                        # Core server
├── business.go                    # Product business logic
├── handlers.go                    # HTTP handlers
├── models.go                      # Data models
├── tools.go                       # Tool schemas

├── config/                        # MCP configurations
│   └── MCP_CONFIGURATION.md      # MCP config doc
├── docs/                          # Team access & diagrams
│   └── TEAM_ACCESS.md            # Team authentication
│   └── architecture-diagram.svg  # Architecture diagram
├── tests/                         # Test scripts & Postman
│   ├── POSTMAN_IMPORT.md         # Postman import guide
│   ├── postman_collection_v2.json# Postman collection
│   ├── run_tests.sh              # Automated test runner
│   └── test_commands.sh          # Manual cURL commands
├── setup-go-env.sh               # Environment setup
├── Dockerfile                    # Container build
├── go.mod, go.sum                # Go modules
```

gcloud run services describe ravi-mcp-server \
gcloud run services add-iam-policy-binding ravi-mcp-server \
<details>
<summary>Deployment & Management</summary>

### Cloud Run Commands (For Authorized Users)

```bash
gcloud run services update ravi-mcp-server \
    --region=europe-west3 \
    --set-env-vars "MICROSERVICE_URL=<MICROSERVICE_URL>"
gcloud run services describe ravi-mcp-server \
  --region=europe-west3 \
  --format="value(spec.template.spec.containers[0].env)"
gcloud run services add-iam-policy-binding ravi-mcp-server \
  --region=europe-west3 \
  --member="user:teammate@domain.com" \
  --role="roles/run.invoker"
```

**Note:** These commands require Google Cloud project access. Contact repository owner for permissions.

### Required Secrets (For Contributors with Deploy Access)

For GitHub Actions deployment, these secrets must be configured:
- GCP_SA_KEY: Google Cloud service account key
- GCP_PROJECT_ID: Google Cloud project ID
- MICROSERVICE_URL: Backend microservice endpoint
- ALLOWED_ORIGIN: CORS allowed origin

**Setup Guide:** Contact the repository owner for security and deployment instructions.

</details>

<details>
<summary>Need More Details?</summary>

- Team Setup: [`docs/TEAM_ACCESS.md`](docs/TEAM_ACCESS.md) - Team authentication
- Configuration: [`config/MCP_CONFIGURATION.md`](config/MCP_CONFIGURATION.md) - MCP setup
- Postman Import: [`tests/POSTMAN_IMPORT.md`](tests/POSTMAN_IMPORT.md) - Postman import guide

</details>



## ⚡️ Latest Changes & Migration Notes

- Security and some test documentation files have been removed.
- Only the following tools are available:
  - `welcome_message`, `health_check`, `create_product`, `get_product`, `update_product`, `delete_product`, `list_products`, `create_multiple_products`, `update_products`, `delete_products`
- All Go files use `package main` and there are no local imports.
- The only required environment variable is `MICROSERVICE_URL`.
- The server exposes `/health` and `/mcp` endpoints.
- Use `tests/run_tests.sh` for automated testing and `tests/test_commands.sh` for manual cURL commands.
- For configuration, see `config/MCP_CONFIGURATION.md` and for team access, see `docs/TEAM_ACCESS.md`.
- The project structure has been simplified and some files have been removed.
