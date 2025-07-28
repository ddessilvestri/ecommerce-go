#!/bin/bash

echo "🧪 Running Address Module Tests..."
echo "=================================="

# Run tests for address module
echo "Testing Address Handler..."
go test ./internal/address -v

# Run all tests in the project
echo ""
echo "🧪 Running All Project Tests..."
echo "================================"
go test ./... -v

echo ""
echo "✅ Test execution completed!" 