# Ravi MCP Server - Sequence Diagram

This diagram shows the unified flow for all major MCP methods: `initialize`, `tools/list`, and `tools/call`.

```mermaid
sequenceDiagram
    participant Client
    participant main.go
    participant handlers.go
    participant (tools.go/business.go)
    participant utils.go

    Client->>main.go: POST /mcp (method: initialize/tools.list/tools.call)
    main.go->>handlers.go: mcpHandler(config)
    handlers.go->>(tools.go/business.go): Handle request (based on method)
    (tools.go/business.go)-->>handlers.go: Result (tools list, server info, or tool result)
    handlers.go->>utils.go: sendJSONRPCResponse (result)
    utils.go-->>Client: JSON response (result)
```

**How to use:**
- Copy the Mermaid code above into https://mermaid.live/ or any Mermaid-compatible tool to view the diagram visually.

**Explanation:**
- The client sends a POST request to `/mcp` with one of the three methods.
- `main.go` routes the request to `handlers.go`.
- `handlers.go` decides which logic to call (`tools.go` for listing tools, `business.go` for tool calls, or internal for initialize).
- The result is sent back to the client using `utils.go`.

