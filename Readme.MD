# Ravi MCP Server

This project is a simple example of a Model Context Protocol (MCP) server for product management. 
It lets you use easy, standard requests (JSON-RPC 2.0) to list, add, update, or delete products. 

The server connects AI agents or other programs to a product microservice, making it easier to manage product data automatically or through natural language commands. 
This proof-of-concept shows how MCP can help organize and automate product management tasks.

**API Version:** v1.0.0

See [docs/api.md](docs/api.md) for a full API reference, including all methods, required parameters, and example payloads. If you change the API, increment the version and update the documentation.

## Overview

This project implements a MCP wrapper server using the Model Context Protocol (MCP) that allows AI agents to interact with a microservice called product service that provides the product data.

It provides tools for:

- Listing available products
- Creating, updating, and deleting products
- Bulk product operations
- Health check and welcome message

<summary>Architecture Diagram</summary>

![Architecture Diagram](docs/architecture-diagram.svg)


## How to Run the Remote MCP Server with GitHub Copilot Agent
<details>

1. **Clone the repository and open it in VSCode.**

   ```bash
    git clone https://github.com/racho8/ravi-mcp-server.git
    cd <repository_name>
    code .
   ```

2. **Setup Configuration Files**

   Copy mcp.json and payload.json from docs/configuration to a new folder called .vscode in your repository.  
     ```bash
     mkdir .vscode
     cp docs/configuration/mcp.json .vscode/
     cp docs/configuration/payload.json .vscode/
     ```

3. **Start the MCP Server**

      Start the ravi-mcp-server by clicking the start button in the mcp.json file in vscode.
      
      **Note**: Sometime, it might show as Error next to the server, but it still works.


4. **Export GCP Authentication Token** (Very Important, otherwise it won't work)

      Run the following command to export your GCP authentication token. 

    ```bash
    export gcp_auth_token=$(gcloud auth print-identity-token)
    ```
      **Note**: Ensure your account has the Cloud Run Invoker role (contact the repository owner if unsure). 


* **Test with GitHub Copilot Chat in Agent Mode**

  Open GitHub Copilot chat in agent mode and test with the following natural language commands:

Use the remote ravi-mcp-server endpoint in mcp.json to list all tools by sending a tools/list request.
Then you can fire nlp querries like: 

       - show me all the available tools
       - Show all the products
       - List all products under Electronics category
       - Increment the price of all the products under Electronics category by 2%
       - Round up the price of all the products under Laptops segment to nearest 100
    
  The agent should automatically generate and execute the corresponding cURL requests in the background and display the results.


* Hack(Optional): If the agent asks you to run the cURL command everytime you give ask a question, you can give below prompt to bypass it:


   > Prompt: *Automatically generate and execute the required cURL requests in the background for any command I provide related to ravi-mcp-server. Display the results directly in plain text format, structured as well-organized tables or lists for better readability. Do not ask for approval before executing each and every cURL command.*

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

**Note**: Contact repository owner for MICROSERVICE_URL and GCP authentication token.
  ```bash
  export MICROSERVICE_URL="<your_microservice_url>"
  export gcp_auth_token=$(gcloud auth print-identity-token)
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

## Troubleshooting

<details>
<summary>Show details</summary>

If you encounter issues while running the mcp server in local:

1. Ensure your `MICROSERVICE_URL` is correct and accessible
2. Check Go version (should be 1.23+)
3. Review error messages in server logs

</details>

## Best Practices for Natural Language Querying

<details>
<summary>Show details</summary>

To smoothly query and see results in natural language, follow these steps before interacting with the MCP server:

1. **Export your GCP authentication token:**
  ```sh
   export gcp_auth_token=$(gcloud auth print-identity-token)
   (Make sure you get necessary roles assigned to your token, contact repo owner if unsure)
  ```

2. **Set the context for your agent (e.g., GitHub Copilot agent chat):**
  > Prompt: *Whenever I ask you about a tool in ravi-mcp-server, I want you to automatically generate and execute the corresponding cURL requests asynchronously in the background, actively displaying the results in plain text format, preferably structured as well-organized tables or lists for enhanced readability. No prior approval of individual cURL commands is necessary prior to execution.*


3. **Use natural language commands to query the MCP server. Example commands:**
  - show me all the available tools
  - Show all the products
  - List all products under Electronics category
  - Increment the price of all the products under Electronics category by 3%
  - Round up the price of all the products under Laptops segment to nearest 100

</details>