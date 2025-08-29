# MCP Server API Reference

**Version:** v1.0.0

## Base Endpoint

- Cloud Run: `https://ravi-mcp-server-256110662801.europe-west3.run.app/mcp`

## Methods

### 1. welcome_message
- **Description:** Get the welcome message from the product service
- **Payload Example:**
```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "method": "tools/call",
  "params": {
    "name": "welcome_message",
    "arguments": {}
  }
}
```

### 2. health_check
- **Description:** Check the health of the product service
- **Payload Example:**
```json
{
  "jsonrpc": "2.0",
  "id": 2,
  "method": "tools/call",
  "params": {
    "name": "health_check",
    "arguments": {}
  }
}
```

### 3. create_product
- **Description:** Create a new product in the store
- **Required:** `name`, `category`, `price`
- **Optional:** `segment`
- **Payload Example:**
```json
{
  "jsonrpc": "2.0",
  "id": 3,
  "method": "tools/call",
  "params": {
    "name": "create_product",
    "arguments": {
      "name": "Laptop",
      "category": "Electronics",
      "price": 999.99,
      "segment": "Laptops"
    }
  }
}
```

### 4. get_product
- **Description:** Retrieve a product by ID
- **Required:** `id`
- **Payload Example:**
```json
{
  "jsonrpc": "2.0",
  "id": 4,
  "method": "tools/call",
  "params": {
    "name": "get_product",
    "arguments": {
      "id": "<product_id>"
    }
  }
}
```

### 5. update_product
- **Description:** Update an existing product by ID
- **Required:** `id`
- **Optional:** `name`, `price`
- **Payload Example:**
```json
{
  "jsonrpc": "2.0",
  "id": 5,
  "method": "tools/call",
  "params": {
    "name": "update_product",
    "arguments": {
      "id": "<product_id>",
      "name": "New Name",
      "price": 1099.99
    }
  }
}
```

### 6. delete_product
- **Description:** Delete a product by ID
- **Required:** `id`
- **Payload Example:**
```json
{
  "jsonrpc": "2.0",
  "id": 6,
  "method": "tools/call",
  "params": {
    "name": "delete_product",
    "arguments": {
      "id": "<product_id>"
    }
  }
}
```

### 7. list_products
- **Description:** List all products in the store
- **Payload Example:**
```json
{
  "jsonrpc": "2.0",
  "id": 7,
  "method": "tools/call",
  "params": {
    "name": "list_products",
    "arguments": {}
  }
}
```

### 8. create_multiple_products
- **Description:** Create multiple products in the store
- **Required:** `products` (array)
- **Payload Example:**
```json
{
  "jsonrpc": "2.0",
  "id": 8,
  "method": "tools/call",
  "params": {
    "name": "create_multiple_products",
    "arguments": {
      "products": [
        {"name": "Laptop", "category": "Electronics", "price": 999.99},
        {"name": "Chair", "category": "Furniture", "price": 199.99}
      ]
    }
  }
}
```

### 9. update_products
- **Description:** Update multiple products at once
- **Required:** `products` (array)
- **Payload Example:**
```json
{
  "jsonrpc": "2.0",
  "id": 9,
  "method": "tools/call",
  "params": {
    "name": "update_products",
    "arguments": {
      "products": [
        {"id": "<product_id>", "name": "New Name", "price": 1099.99}
      ]
    }
  }
}
```

### 10. delete_products
- **Description:** Delete multiple products at once
- **Required:** `ids` (array)
- **Payload Example:**
```json
{
  "jsonrpc": "2.0",
  "id": 10,
  "method": "tools/call",
  "params": {
    "name": "delete_products",
    "arguments": {
      "ids": ["<product_id1>", "<product_id2>"]
    }
  }
}
```

---

**Note:**
- All requests must include a valid GCP identity token in the `Authorization` header.
- If you change the API, increment the version and update this file.
