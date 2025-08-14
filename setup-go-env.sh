#!/bin/bash
# Setup script for Go environment to avoid version mismatch issues

echo "🔧 Setting up Go environment..."

# Check if Homebrew Go is installed
if [ -d "/opt/homebrew/Cellar/go" ]; then
    # Find the latest Go version in Homebrew
    GO_VERSION=$(ls -1 /opt/homebrew/Cellar/go/ | sort -V | tail -1)
    export GOROOT="/opt/homebrew/Cellar/go/${GO_VERSION}/libexec"
    export PATH="$GOROOT/bin:$PATH"
    echo "✅ Using Homebrew Go ${GO_VERSION}"
    echo "   GOROOT: $GOROOT"
else
    echo "⚠️  Homebrew Go not found, using system Go"
fi

# Display current Go configuration
echo "📋 Go Environment:"
go version
echo "   GOROOT: $(go env GOROOT)"
echo "   GOPATH: $(go env GOPATH)"

echo "🚀 Environment ready! You can now run:"
echo "   go test -v ./..."
echo "   go run main.go"
