#!/bin/bash

# Initialize project
echo "🚀 Setting up Flood IDF Tool"
git init
git add .
git commit -m "Initial commit"

# Get NASA API key
read -p "🔑 Enter NASA Earthdata API key: " nasa_key
echo "NASA_API_KEY=$nasa_key" > .env

# Build and run
echo "🐳 Building Docker containers..."
docker-compose build
echo "🚀 Starting services..."
docker-compose up -d

echo -e "\n✅ Setup completed!"
echo "🌐 Access your tools:"
echo "   - API Docs: http://localhost:8080/docs"
echo "   - Dashboard: http://localhost:3000"
echo "   - CLI Tool: docker exec -it flood-idf-tool_idf-tool_1 ./idf-tool -cli"
