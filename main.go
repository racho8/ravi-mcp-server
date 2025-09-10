package main

import (
	"log"
	"net/http"
	"os"
	"encoding/json"
)

func main() {
	microserviceURL := os.Getenv("MICROSERVICE_URL")
	if microserviceURL != "" {
		log.Printf("MICROSERVICE_URL: configured")
	} else {
		log.Printf("MICROSERVICE_URL: not configured!!")
	}
	log.Printf("PORT: %s", os.Getenv("PORT"))

	config := Config{
		MicroserviceURL: microserviceURL,
		Port:            os.Getenv("PORT"),
	}
	if config.Port == "" {
		config.Port = "8080"
	}

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {

		// needed for CORS support, especially for web-based clients
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "healthy", "service": "ravi-mcp-server"}`))
	})
	http.HandleFunc("/mcp", func(w http.ResponseWriter, r *http.Request) {

		// needed for CORS support, especially for web-based clients
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		mcpHandler(config)(w, r)
	})
	http.HandleFunc("/mcp/discover", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		// Import encoding/json and tools from tools.go
		if err := json.NewEncoder(w).Encode(tools); err != nil {
			http.Error(w, "Failed to encode tools", http.StatusInternalServerError)
		}
	})

	log.Printf("Starting MCP server on port %s", config.Port)
	log.Printf("MCP JSON-RPC 2.0 Protocol supported methods:")
	log.Printf("  - initialize")
	log.Printf("  - tools/list")
	log.Printf("  - tools/call")

	if err := http.ListenAndServe(":"+config.Port, nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
