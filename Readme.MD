# Ravi MCP Server

A Model Context Protocol (MCP) server for product management via JSON-RPC 2.0.



## Overview

This project implements a MCP wrapper server using the Model Context Protocol (MCP) that allows AI agents to interact with a microservice called product service that provides the product data.

It provides tools for:

- Listing available products
- Creating, updating, and deleting products
- Bulk product operations
- Health check and welcome message

<details>
<summary>Architecture Diagram</summary>

![Architecture Diagram](docs/architecture-diagram.svg)

</details>



## Installation

<details>
<summary>Manual Installation</summary>

<b>Prerequisites</b>

- Go 1.23+
- Access to a product microservice (ask repo owner for MICROSERVICE_URL)

<b>Setup</b>
1. Clone this repository
2. Install dependencies:
  ```bash
  go mod download
  ```
3. Set environment variables:
  ```bash
  export MICROSERVICE_URL="<your_microservice_url>"
  export PORT=8080
  ```
4. Run the server:
  ```bash
  go run main.go
  # or
  go run .
  ```
5. Test health endpoint:
  ```bash
  curl http://localhost:8080/health
  ```
6. Use `/mcp` endpoint for JSON-RPC requests (see `tests/test_commands.sh` for examples)

</details>


## Available Tools

- `list_products` — List all products
- `create_product` — Add a new product
- `get_product` — Get product details by ID
- `update_product` — Update product by ID
- `delete_product` — Delete product by ID
- `create_multiple_products` — Add multiple products
- `update_products` — Update multiple products
- `delete_products` — Delete multiple products
- `health_check` — Check server status
- `welcome_message` — Get welcome message



## Using Ravi MCP Server with GitHub Copilot Agent

<details>
<summary>Show details</summary>

To use this remote MCP server with GitHub Copilot agent mode (or any MCP-compatible client):

<b>1. Add <code>mcp.json</code> to your workspace (recommended location: <code>.vscode/mcp.json</code>)</b>

Example structure:
```json
{
  "servers": {
    "ravi-mcp-server": {
      "command": "curl",
      "args": [
        "-X", "POST",
        "https://ravi-mcp-server-256110662801.europe-west3.run.app/mcp",
        "-H", "Content-Type: application/json",
        "-H", "Authorization: Bearer GCP_AUTH_TOKEN",
        "-d", "@.vscode/payload.json"
      ]
    }
  }
}
```

<b>2. Authentication Token</b>

- You must provide a valid Google Cloud Bearer token for authentication.
- Recommended: Use `gcloud auth print-identity-token` to generate a token.
- Set the token in your environment or update the `mcp.json` as needed.

<b>3. Example Usage</b>

- With configuration in place, you can use natural language commands in Copilot agent mode, e.g.:
  - "List all products"
  - "Create a new product called 'iPhone 15' for $999"

<b>4. Sample payload.json</b>

Your `.vscode/payload.json` file is referenced in `mcp.json` and should contain a valid MCP JSON-RPC request. Example:

```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "method": "tools/list"
}
```

You can modify this file to call any tool, e.g.:

```json
{
  "jsonrpc": "2.0",
  "id": 2,
  "method": "tools/call",
  "params": {
    "name": "list_products"
  }
}
```

<b>5. More Info</b>

- See [`config/MCP_CONFIGURATION.md`](config/MCP_CONFIGURATION.md) for advanced configuration and troubleshooting.
- For team access setup, see [`docs/TEAM_ACCESS.md`](docs/TEAM_ACCESS.md).

</details>


## Troubleshooting

<details>
<summary>Show details</summary>

If you encounter issues:

1. Ensure your `MICROSERVICE_URL` is correct and accessible
2. Check Go version (should be 1.23+)
3. Review error messages in server logs

</details>