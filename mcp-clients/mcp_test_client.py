#!/usr/bin/env python3
"""
Simple MCP Client for testing your server
Usage: python mcp_test_client.py "list all products"
"""

import json
import subprocess
import sys

def call_mcp_server(method, params=None):
    """Call your MCP server with JSON-RPC"""
    if params is None:
        params = {}
    
    # Get GCP token
    try:
        token_result = subprocess.run(['gcloud', 'auth', 'print-access-token'], 
                                    capture_output=True, text=True, check=True)
        token = token_result.stdout.strip()
    except subprocess.CalledProcessError:
        print("Error: Unable to get GCP access token")
        return None
    
    # Prepare the JSON-RPC request
    request = {
        "jsonrpc": "2.0",
        "id": 1,
        "method": method,
        "params": params
    }
    
    # Call your MCP server
    curl_cmd = [
        'curl', '-s', '-X', 'POST',
        'https://ravi-mcp-server-256110662801.europe-west3.run.app/mcp',
        '-H', 'Content-Type: application/json',
        '-H', f'Authorization: Bearer {token}',
        '-d', json.dumps(request)
    ]
    
    try:
        result = subprocess.run(curl_cmd, capture_output=True, text=True, check=True)
        return json.loads(result.stdout)
    except (subprocess.CalledProcessError, json.JSONDecodeError) as e:
        print(f"Error calling MCP server: {e}")
        return None

def natural_language_to_tool(command):
    """Convert natural language to MCP tool calls"""
    command = command.lower()
    
    if "list" in command and "product" in command:
        return call_mcp_server("tools/call", {"name": "list_products", "arguments": {}})
    elif "create" in command and "product" in command:
        # Simple parsing - you could make this more sophisticated
        return "Please use: create_product(name, category, price)"
    elif "health" in command or "check" in command:
        return call_mcp_server("tools/call", {"name": "health_check", "arguments": {}})
    elif "welcome" in command:
        return call_mcp_server("tools/call", {"name": "welcome_message", "arguments": {}})
    elif "tools" in command or "available" in command:
        return call_mcp_server("tools/list", {})
    else:
        return "Available commands: list products, health check, welcome message, show tools"

if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("Usage: python mcp_test_client.py 'your natural language command'")
        print("Examples:")
        print("  python mcp_test_client.py 'list all products'")
        print("  python mcp_test_client.py 'health check'")
        print("  python mcp_test_client.py 'show available tools'")
        sys.exit(1)
    
    command = " ".join(sys.argv[1:])
    result = natural_language_to_tool(command)
    
    if result:
        print(json.dumps(result, indent=2))
