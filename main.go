
package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	microserviceURL := os.Getenv("MICROSERVICE_URL")
	if microserviceURL != "" {
		log.Printf("MICROSERVICE_URL: configured ✓")
	} else {
		log.Printf("MICROSERVICE_URL: not configured ⚠️")
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
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "healthy", "service": "ravi-mcp-server"}`))
	})
	http.HandleFunc("/mcp", mcpHandler(config))

	log.Printf("Starting MCP server on port %s", config.Port)
	log.Printf("MCP JSON-RPC 2.0 Protocol supported methods:")
	log.Printf("  - initialize")
	log.Printf("  - tools/list")
	log.Printf("  - tools/call")

	if err := http.ListenAndServe(":"+config.Port, nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
