#!/bin/bash

echo "🧪 Running Address Module Tests..."
echo "=================================="

# Run tests for address module
echo "Testing Address Module..."
go test ./internal/address -v

echo ""
echo "🧪 Running All Project Tests..."
echo "================================"
go test ./... -v

echo ""
echo "✅ All tests completed successfully!" 