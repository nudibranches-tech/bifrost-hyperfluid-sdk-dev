#!/bin/bash

# Test runner script for Bifrost SDK

set -e

case "$1" in
  unit)
    echo "🧪 Running unit tests..."
    go test -v ./sdk/...
    ;;
  integration)
    echo "🔗 Running integration tests..."
    go test -v ./integration_tests
    ;;
  all)
    echo "🚀 Running all tests..."
    echo ""
    echo "1️⃣ Unit tests..."
    go test -v ./sdk/...
    echo ""
    echo "2️⃣ Integration tests..."
    go test -v ./integration_tests
    echo ""
    echo "✅ All tests completed!"
    ;;
  *)
    echo "Usage: $0 {unit|integration|all}"
    echo ""
    echo "  unit        - Run fast unit tests"
    echo "  integration - Run integration tests (requires external services)"
    echo "  all         - Run all tests"
    exit 1
    ;;
esac
