# MCP (Model Context Protocol) Server

A Go-based mcp server wrapper that provides a unified interface for interacting with microservices using a tool-based approach. 
This server is designed to run on Google Cloud Run and interfaces with other microservices.

## Overview

The MCP Server acts as an intermediary layer that standardizes communication with underlying microservices. 
It provides a tool-based interface where each tool represents a specific microservice operation/end point.

## Features

- Tool-based API interface
- Integration with Google Cloud Run
- Authentication support
- JSON request/response handling
- Configurable microservice endpoints
- Automatic token forwarding to microservices

## Available Tools

The server provides the following tools:
- `welcome_message`: Get welcome message from the product service
- `health_check`: Check product service health
- `create_product`: Create a new product
- `get_product`: Retrieve a product by ID
- `update_product`: Update an existing product
- `delete_product`: Delete a product
- `list_products`: Get all products

## API Endpoints

- `GET /tools`: List all available tools and their schemas
- `POST /mcp`: Execute tool operations

## Environment Variables

- `MICROSERVICE_URL`: URL of the target microservice
- `PORT`: Server port (defaults to 8080)

## Deployment

The application is configured for deployment to Google Cloud Run using GitHub Actions. The deployment workflow:
1. Builds a Docker container
2. Pushes to Google Artifact Registry
3. Deploys to Cloud Run

### Required Secrets

- `GCP_SA_KEY`: Google Cloud service account key
- `GCP_PROJECT_ID`: Google Cloud project ID
- `MICROSERVICE_URL`: Target microservice URL

## Usage Example

```bash
# Execute a tool
curl -X POST https://[YOUR-SERVER]/mcp \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $(gcloud auth print-identity-token)" \
  -d '{"tool_name": "create_product", "params": {"name": "Laptop", "price": 999}}'

```

## Some handy commands

### Update the environment variable for the microservice URL
gcloud run services update ravi-mcp-server \
    --region=europe-west3 \
    --set-env-vars "MICROSERVICE_URL=https://product-service-256110662801.europe-west3.run.app"

### To describe the service to see the current environment variables
gcloud run services describe ravi-mcp-server \
  --region=europe-west3 \
  --format="value(spec.template.spec.containers[0].env)"

### Set the environment variable for the microservice URL
export MICROSERVICE_URL=https://product-service-256110662801.europe-west3.run.app


### To add cloud run.invoker role for any one in the organisation:
gcloud run services add-iam-policy-binding ravi-mcp-server \
  --region=europe-west3 \
  --member="domain:your-organization-domain.com" \
  --role="roles/run.invoker"


